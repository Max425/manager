import psycopg2
from psycopg2 import sql
from random import choice, randint, sample
from datetime import datetime, timedelta
import yaml
from faker import Faker

# Инициализация Faker
fake = Faker()

# Чтение конфигурации из config.yml
with open('/Users/msikanov/WebstormProjects/manager/backend/configs/config.yml', 'r') as file:
    config = yaml.safe_load(file)

db_config = config['db']

# Установите соединение с вашей базой данных PostgreSQL
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
         "Старший frontend разработчик", "Младший QA", "Старший QA"]
stages = ["Начало", "В процессе", "Тестирование", "Завершение"]
male_images = [
    "https://mykaleidoscope.ru/x/uploads/posts/2022-09/1663110850_6-mykaleidoscope-ru-p-spokoinii-chelovek-vkontakte-8.jpg",
    "https://get.pxhere.com/photo/man-person-people-photography-travel-male-smile-photograph-80529.jpg",
    "https://mtdata.ru/u3/photoD852/20501229401-0/original.jpg",
    "https://get.pxhere.com/photo/person-girl-woman-hair-photography-portrait-model-youth-fashion-blue-lady-hairstyle-smile-long-hair-face-dress-eye-head-skin-beauty-blond-photo-shoot-brown-hair-portrait-photography-108386.jpg",
    "https://74foto.ru/800/600/http/cdn1.flamp.ru/bc57c2126b20646180c92643db78d9f0.jpg"
]
female_images = [
    "https://avatars.mds.yandex.net/i?id=1be90a756cf43d19ef403e6ea6e45434_l-4615485-images-thumbs&ref=rim&n=13&w=1080&h=1080",
    "https://get.pxhere.com/photo/person-girl-woman-hair-photography-portrait-model-youth-fashion-blue-lady-hairstyle-smile-long-hair-face-dress-eye-head-skin-beauty-blond-photo-shoot-brown-hair-portrait-photography-108386.jpg",
    "https://get.pxhere.com/photo/outdoor-person-girl-sun-woman-hair-white-photography-cute-summer-female-portrait-model-young-red-fashion-lady-facial-expression-hairstyle-smiling-smile-long-hair-close-up-caucasian-face-dress-happy-happiness-eye-head-skin-beauty-attractive-photo-shoot-pretty-girl-brown-hair-cute-girl-happy-girl-happy-woman-portrait-photography-supermodel-683657.jpg",
    "https://uprostim.com/wp-content/uploads/2021/03/image062-51-scaled.jpg"
]
project_images = [
    "https://i.pinimg.com/originals/b9/05/3d/b9053d873e9f69058997913e0fffca2e.png",
    "https://gas-kvas.com/grafic/uploads/posts/2024-01/gas-kvas-com-p-simvoli-dlya-logotipov-na-prozrachnom-fone-34.png",
    "https://russia-dropshipping.ru/800/600/http/thelawofattraction.ru/wp-content/uploads/a/c/0/ac007477e7e0d998fa4d822bc1730255.png",
    "https://img.razrisyika.ru/kart/94/372802-logo-6.jpg",
    "https://free-png.ru/wp-content/uploads/2020/10/Nike-logo-506c4872.png"
]

# Создание компании
cursor.execute("""
    INSERT INTO company (name, positions, image, description)
    VALUES (%s, %s::text[], %s, %s)
    RETURNING id
""", ("IT Компания", roles, fake.image_url(), fake.text()))
company_id = cursor.fetchone()[0]

# Функция для создания сотрудников
def create_employee(company_id):
    name = fake.first_name()
    surname = fake.last_name()
    role = choice(roles)
    email = f"{name.lower()}.{surname.lower()}@{fake.domain_name()}"
    image = choice(male_images if fake.random_int(0, 1) == 0 else female_images)
    cursor.execute("""
        INSERT INTO employee (company_id, name, position, mail, password, salt, image, rating, created_at, updated_at)
        VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s)
        RETURNING id
    """, (company_id, f"{name} {surname}", f'"{role}"', email, fake.password(), fake.md5(), image,
          round(fake.pyfloat(min_value=0, max_value=5, right_digits=2), 2), fake.date_time_this_year(),
          fake.date_time_this_year()))
    return cursor.fetchone()[0]

# Создание сотрудников
employee_ids = [create_employee(company_id) for _ in range(70)]

# Функция для создания проектов
def create_project(company_id):
    project_name = fake.catch_phrase()
    deadline = fake.date_time_between(start_date='+30d', end_date='+90d')
    image = choice(project_images)
    cursor.execute("""
        INSERT INTO project (company_id, name, stages, image, description, current_stage, deadline, status, complexity, created_at, updated_at)
        VALUES (%s, %s, %s::text[], %s, %s, %s, %s, %s, %s, %s, %s)
        RETURNING id
    """, (company_id, project_name, stages, image, fake.text(), randint(0, 3), deadline, randint(0, 1),
          randint(1, 10), fake.date_time_this_year(), fake.date_time_this_year()))
    return cursor.fetchone()[0]

# Создание проектов и назначение сотрудников
for _ in range(30):  # Увеличено количество проектов до 30
    project_id = create_project(company_id)
    assigned_employees = sample(employee_ids, randint(3, 5))
    for emp_id in assigned_employees:
        cursor.execute("""
            INSERT INTO employee_project (project_id, employee_id)
            VALUES (%s, %s)
        """, (project_id, emp_id))

# Сохранение изменений и закрытие соединения
conn.commit()
cursor.close()
conn.close()
