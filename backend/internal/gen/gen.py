import psycopg2
from psycopg2 import sql
from random import choice, randint, sample, shuffle
from datetime import datetime, timedelta
import yaml
from faker import Faker
import requests
import base64
import urllib.parse

# Инициализация Faker с русской локализацией
fake = Faker('ru_RU')

# Чтение конфигурации из config.yml
with open('/Users/msikanov/WebstormProjects/manager/backend/configs/config.yml', 'r') as file:
    config = yaml.safe_load(file)

db_config = config['db']

# Установите соединение с базой данных PostgreSQL
conn = psycopg2.connect(
    dbname=db_config['dbname'],
    user=db_config['username'],
    password=db_config['password'],
    host=db_config['host'],
    port=db_config['port'],
    sslmode=db_config['sslmode']
)
cursor = conn.cursor()

# Вспомогательные данные
roles = ["Младший backend разработчик", "Старший backend разработчик", "Младший frontend разработчик",
         "Старший frontend разработчик", "Младший QA инженер", "Старший QA инженер"]
stages = ["Инициация", "Разработка", "Тестирование", "Завершение"]

# Функция для генерации аватара и конвертации в Base64
def generate_avatar_base64(name):
    encoded_name = urllib.parse.quote(name)
    avatar_url = f"https://api.dicebear.com/8.x/initials/svg?seed={encoded_name}"
    try:
        response = requests.get(avatar_url)
        response.raise_for_status()
        avatar_base64 = base64.b64encode(response.content).decode('utf-8')
        return avatar_base64
    except requests.RequestException as e:
        print(f"Ошибка при загрузке аватара для {name}: {e}")
        return "PHN2ZyB3aWR0aD0iMTAwIiBoZWlnaHQ9IjEwMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMTAwIiBoZWlnaHQ9IjEwMCIgZmlsbD0iI2NjY2NjYyIvPjx0ZXh0IHg9IjUwIiB5PSI1MCIgZm9udC1zaXplPSI0MCIgZmlsbD0iI2ZmZmZmZiIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZHk9IjAuM2VtIj5OPkE8L3RleHQ+PC9zdmc+"

# Функция для создания рейтинга с учётом просроченных и завершённых проектов
def generate_rating(num_projects):
    rating = [0]  # Начальное значение рейтинга
    overdue_count = 0
    for _ in range(num_projects):  # Генерируем проекты
        change = choice([1, -1])  # Увеличение (+1) или уменьшение (-1)
        rating.append(rating[-1] + change)
        if change == -1:
            overdue_count += 1
    completed_on_time = num_projects - overdue_count
    return rating, overdue_count, completed_on_time

# Создание компании
cursor.execute("""
    INSERT INTO company (name, positions, image, description)
    VALUES (%s, %s::text[], %s, %s)
    RETURNING id
""", ("ООО ТехноИнновации", roles, generate_avatar_base64("ТехноИнновации"), fake.text(max_nb_chars=200)))
company_id = cursor.fetchone()[0]

# Функция для создания сотрудника (без active_projects_count на этом этапе)
def create_employee(company_id):
    name = fake.first_name()
    surname = fake.last_name()
    full_name = f"{name} {surname}"
    role = choice(roles)
    email = f"{name.lower()}.{surname.lower()}@{fake.domain_name()}"
    num_completed_projects = randint(1, 10)
    rating, overdue_count, completed_on_time = generate_rating(num_completed_projects)
    avatar_base64 = generate_avatar_base64(full_name)
    cursor.execute("""
        INSERT INTO employee (company_id, name, position, mail, password, salt, image, rating, 
                             active_projects_count, overdue_projects_count, total_projects_count, 
                             created_at, updated_at)
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
        RETURNING id
    """, (company_id, full_name, role, email, fake.password(), fake.md5(),
          avatar_base64, rating,
          0,  # active_projects_count (будет обновлено позже)
          overdue_count,
          num_completed_projects,  # total_projects_count (пока только завершённые)
          fake.date_time_this_year(), fake.date_time_this_year()))
    return cursor.fetchone()[0], num_completed_projects, overdue_count

# Создание 20 сотрудников
employees = [create_employee(company_id) for _ in range(20)]
employee_ids = [emp[0] for emp in employees]

# Функция для создания проекта с заданным статусом
def create_project(company_id, status):
    project_name = fake.catch_phrase()
    deadline = fake.date_time_between(start_date='+30d', end_date='+90d')
    avatar_base64 = generate_avatar_base64(project_name)
    cursor.execute("""
        INSERT INTO project (company_id, name, stages, image, description, current_stage, deadline, status, complexity, created_at, updated_at)
        VALUES (%s, %s, %s::text[], %s, %s, %s, %s, %s, %s, %s, %s)
        RETURNING id
    """, (company_id, project_name, stages, avatar_base64,
          fake.text(max_nb_chars=200), randint(0, 3), deadline, status,
          randint(1, 10), fake.date_time_this_year(), fake.date_time_this_year()))
    return cursor.fetchone()[0]

# Создание проектов: 10 завершённых и 15 активных (чтобы хватило на всех сотрудников)
completed_project_ids = [create_project(company_id, 0) for _ in range(10)]
active_project_ids = [create_project(company_id, 1) for _ in range(15)]

# Назначение сотрудников на проекты
employee_project_assignments = {emp_id: [] for emp_id in employee_ids}
active_project_count = {emp_id: 0 for emp_id in employee_ids}

# Назначаем сотрудников на завершённые проекты
for emp_id, num_completed, overdue_count in employees:
    assigned_completed = sample(completed_project_ids, min(num_completed, len(completed_project_ids)))
    for project_id in assigned_completed:
        cursor.execute("""
            INSERT INTO employee_project (project_id, employee_id)
            VALUES (%s, %s)
        """, (project_id, emp_id))
        employee_project_assignments[emp_id].append(project_id)

# Генерируем active_projects_count для каждого сотрудника
for emp_id in employee_ids:
    active_project_count[emp_id] = randint(0, 5)  # Случайное количество активных проектов

# Назначаем сотрудников на активные проекты в соответствии с active_projects_count
available_active_projects = active_project_ids.copy()
shuffle(available_active_projects)

for emp_id in employee_ids:
    num_active = active_project_count[emp_id]
    if num_active > 0:
        # Если не хватает активных проектов, создаём дополнительные
        while len(available_active_projects) < num_active:
            new_project_id = create_project(company_id, 1)
            available_active_projects.append(new_project_id)

        # Назначаем сотрудника в num_active активных проектов
        assigned_active = available_active_projects[:num_active]
        available_active_projects = available_active_projects[num_active:]

        for project_id in assigned_active:
            cursor.execute("""
                INSERT INTO employee_project (project_id, employee_id)
                VALUES (%s, %s)
            """, (project_id, emp_id))
            employee_project_assignments[emp_id].append(project_id)

# Обновляем active_projects_count и total_projects_count для сотрудников
for emp_id, num_completed, overdue_count in employees:
    active_count = active_project_count[emp_id]
    total_count = num_completed + active_count  # Завершённые + активные
    cursor.execute("""
        UPDATE employee 
        SET active_projects_count = %s, total_projects_count = %s 
        WHERE id = %s
    """, (active_count, total_count, emp_id))

# Проверка и назначение проектов для сотрудников, у которых нет проектов
for emp_id in employee_ids:
    if not employee_project_assignments[emp_id]:
        project_id = choice(completed_project_ids + active_project_ids)
        cursor.execute("""
            INSERT INTO employee_project (project_id, employee_id)
            VALUES (%s, %s)
        """, (project_id, emp_id))
        employee_project_assignments[emp_id].append(project_id)

# Сохранение изменений и закрытие соединения
conn.commit()
cursor.close()
conn.close()