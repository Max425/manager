import psycopg2
from psycopg2 import sql
from random import choice, randint, sample
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
    # Кодируем имя для использования в URL
    encoded_name = urllib.parse.quote(name)
    # Используем DiceBear API для генерации SVG аватара
    avatar_url = f"https://api.dicebear.com/8.x/initials/svg?seed={encoded_name}"
    try:
        response = requests.get(avatar_url)
        response.raise_for_status()
        # Конвертируем SVG в Base64
        avatar_base64 = base64.b64encode(response.content).decode('utf-8')
        return avatar_base64
    except requests.RequestException as e:
        print(f"Ошибка при загрузке аватара для {name}: {e}")
        # Возвращаем запасной Base64 (пустой аватар)
        return "PHN2ZyB3aWR0aD0iMTAwIiBoZWlnaHQ9IjEwMCIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj48cmVjdCB3aWR0aD0iMTAwIiBoZWlnaHQ9IjEwMCIgZmlsbD0iI2NjY2NjYyIvPjx0ZXh0IHg9IjUwIiB5PSI1MCIgZm9udC1zaXplPSI0MCIgZmlsbD0iI2ZmZmZmZiIgdGV4dC1hbmNob3I9Im1pZGRsZSIgZHk9IjAuM2VtIj5OPkE8L3RleHQ+PC9zdmc+"

# Функция для создания рейтинга с учётом просроченных и завершённых проектов
def generate_rating(num_projects):
    rating = [0]  # Начальное значение рейтинга
    overdue_count = 0
    for _ in range(num_projects - 1):  # Генерируем оставшиеся проекты
        # Случайно определяем, просрочен ли проект (50% шанс)
        if choice([True, False]):
            rating.append(rating[-1]-1)
            overdue_count += 1
        else:
            rating.append(rating[-1]+1)
    return rating, overdue_count

# Создание компании
cursor.execute("""
    INSERT INTO company (name, positions, image, description)
    VALUES (%s, %s::text[], %s, %s)
    RETURNING id
""", ("ООО ТехноИнновации", roles, generate_avatar_base64("ТехноИнновации"), fake.text(max_nb_chars=200)))
company_id = cursor.fetchone()[0]

# Функция для создания сотрудника с учётом рейтинга и проектов
def create_employee(company_id):
    name = fake.first_name()
    surname = fake.last_name()
    full_name = f"{name} {surname}"
    role = choice(roles)
    email = f"{name.lower()}.{surname.lower()}@{fake.domain_name()}"
    # Генерируем количество завершённых проектов (от 1 до 10)
    num_completed_projects = randint(1, 10)
    rating, overdue_count = generate_rating(num_completed_projects)
    avatar_base64 = generate_avatar_base64(full_name)
    cursor.execute("""
        INSERT INTO employee (company_id, name, position, mail, password, salt, image, rating, 
                             active_projects_count, overdue_projects_count, total_projects_count, 
                             created_at, updated_at)
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
        RETURNING id
    """, (company_id, full_name, role, email, fake.password(), fake.md5(),
          avatar_base64, rating,
          0,  # active_projects_count (0 по умолчанию, обновляется триггером)
          overdue_count,  # overdue_projects_count
          num_completed_projects,  # total_projects_count
          fake.date_time_this_year(), fake.date_time_this_year()))
    return cursor.fetchone()[0]

# Создание 20 сотрудников
employee_ids = [create_employee(company_id) for _ in range(20)]

# Функция для создания проекта
def create_project(company_id):
    project_name = fake.catch_phrase()
    deadline = fake.date_time_between(start_date='+30d', end_date='+90d')
    avatar_base64 = generate_avatar_base64(project_name)
    # Случайный статус (0 - завершён, 1 - активен)
    status = randint(0, 1)
    cursor.execute("""
        INSERT INTO project (company_id, name, stages, image, description, current_stage, deadline, status, complexity, created_at, updated_at)
        VALUES (%s, %s, %s::text[], %s, %s, %s, %s, %s, %s, %s, %s)
        RETURNING id
    """, (company_id, project_name, stages, avatar_base64,
          fake.text(max_nb_chars=200), randint(0, 3), deadline, status,
          randint(1, 10), fake.date_time_this_year(), fake.date_time_this_year()))
    return cursor.fetchone()[0]

# Создание 10 проектов
project_ids = [create_project(company_id) for _ in range(10)]

# Назначение сотрудников на проекты
employee_project_assignments = {emp_id: [] for emp_id in employee_ids}
active_project_count = {}  # Для отслеживания активных проектов
for project_id in project_ids:
    cursor.execute("SELECT status, deadline FROM project WHERE id = %s", (project_id,))
    status, deadline = cursor.fetchone()
    # Назначаем 3-5 случайных сотрудников на каждый проект
    assigned_employees = sample(employee_ids, randint(3, 5))
    for emp_id in assigned_employees:
        cursor.execute("""
            INSERT INTO employee_project (project_id, employee_id)
            VALUES (%s, %s)
        """, (project_id, emp_id))
        employee_project_assignments[emp_id].append(project_id)
        if status == 1:  # Проект активен
            active_project_count[emp_id] = active_project_count.get(emp_id, 0) + 1

# Обновление active_projects_count для сотрудников
for emp_id, count in active_project_count.items():
    cursor.execute("""
        UPDATE employee SET active_projects_count = %s WHERE id = %s
    """, (count, emp_id))

# Проверка и назначение проектов для сотрудников, у которых нет проектов
for emp_id in employee_ids:
    if not employee_project_assignments[emp_id]:
        project_id = choice(project_ids)
        cursor.execute("""
            INSERT INTO employee_project (project_id, employee_id)
            VALUES (%s, %s)
        """, (project_id, emp_id))

# Сохранение изменений и закрытие соединения
conn.commit()
cursor.close()
conn.close()