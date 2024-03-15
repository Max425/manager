create table company
(
    id          serial primary key,
    name        text not null,
    positions   text[],
    image       text,
    description text,
    created_at  timestamptz default timezone('Europe/Moscow'::text, now()),
    updated_at  timestamptz default timezone('Europe/Moscow'::text, now())
);

create table employee
(
    id                     serial primary key,
    company_id             int  not null references "company" on delete cascade,
    name                   text not null,
    position               text,
    mail                   text unique,
    password               text not null,
    salt                   text not null,
    image                  text,
    rating                 decimal,
    active_projects_count  int         default 0,
    overdue_projects_count int         default 0,
    total_projects_count   int         default 0,
    created_at             timestamptz default timezone('Europe/Moscow'::text, now()),
    updated_at             timestamptz default timezone('Europe/Moscow'::text, now())
);

create index on employee (company_id, position);

create table project
(
    id            serial primary key,
    company_id    int      not null references "company" on delete cascade,
    name          text     not null,
    stages        text[],
    image         text,
    description   text,
    current_stage int,
    deadline      timestamptz,
    status        smallint not null default 1 check (status between 0 and 1 ),
    complexity    smallint check (complexity between 0 and 10),
    created_at    timestamptz       default timezone('Europe/Moscow'::text, now()),
    updated_at    timestamptz       default timezone('Europe/Moscow'::text, now())
);

create index on employee (company_id);

create table employee_project
(
    id          serial primary key,
    project_id  int not null references "project" on delete cascade,
    employee_id int not null references "employee" on delete cascade
);

create index on employee_project (project_id);
create index on employee_project (employee_id);

-- Создание триггера для увеличения количества проектов у сотрудника
CREATE OR REPLACE FUNCTION increase_project_count_trigger()
    RETURNS TRIGGER AS
$$
BEGIN
    -- Увеличение количества проектов и активных проектов у сотрудника
    UPDATE employee
    SET active_projects_count = active_projects_count + 1,
        total_projects_count  = total_projects_count + 1
    WHERE id = NEW.employee_id;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER increase_project_count
    AFTER INSERT
    ON employee_project
    FOR EACH ROW
EXECUTE FUNCTION increase_project_count_trigger();

-- Создание триггера для уменьшения количества активных проектов у сотрудников
CREATE OR REPLACE FUNCTION decrease_active_project_count_trigger()
    RETURNS TRIGGER AS
$$
BEGIN
    -- Уменьшение количества активных проектов у сотрудника
    UPDATE employee
    SET active_projects_count = active_projects_count - 1
    WHERE id in (select employee_id from employee_project where project_id = OLD.id);

    -- Увеличение количества просроченных проектов у сотрудника
    IF OLD.deadline < NOW() THEN
        UPDATE employee
        SET overdue_projects_count = overdue_projects_count + 1
        WHERE id IN (SELECT employee_id FROM employee_project WHERE project_id = OLD.id);
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER decrease_active_project_count
    AFTER UPDATE
    ON project
    FOR EACH ROW
    WHEN (OLD.status = 1 AND NEW.status = 0)
EXECUTE FUNCTION decrease_active_project_count_trigger();


-- Вставка тестовых данных для таблицы company
INSERT INTO company (name, positions, image, description)
VALUES ('Company A', ARRAY ['Manager', 'Developer'], 'company_a_image.jpg', 'Description for Company A'),
       ('Company B', ARRAY ['CEO', 'Designer'], 'company_b_image.jpg', 'Description for Company B'),
       ('Company C', ARRAY ['Manager', 'Developer', 'QA'], 'company_c_image.jpg', 'Description for Company C');

-- Вставка тестовых данных для таблицы employee
INSERT INTO employee (company_id, name, position, mail, password, salt, image, rating)
VALUES (1, 'John Doe', 'Manager', 'john.doe@example.com', 'password_hash', 'salt123', 'john_doe_image.jpg', 4.5),
       (1, 'Jane Smith', 'Developer', 'jane.smith@example.com', 'password_hash', 'salt456', 'jane_smith_image.jpg',
        4.2),
       (2, 'Alice Johnson', 'CEO', 'alice.johnson@example.com', 'password_hash', 'salt789', 'alice_johnson_image.jpg',
        4.8),
       (3, 'Bob Brown', 'Manager', 'bob.brown@example.com', 'password_hash', 'salt101', 'bob_brown_image.jpg', 4.6);

-- Вставка тестовых данных для таблицы project
INSERT INTO project (company_id, name, stages, image, description, current_stage, deadline, status, complexity)
VALUES (1, 'Project A', ARRAY ['Design', 'Development', 'Testing'], 'project_a_image.jpg', 'Description for Project A',
        2, '2024-04-30', 1, 7),
       (2, 'Project B', ARRAY ['Planning', 'Execution'], 'project_b_image.jpg', 'Description for Project B', 1,
        '2024-05-15', 1, 5),
       (3, 'Project C', ARRAY ['Design', 'Development', 'Testing', 'Deployment'], 'project_c_image.jpg',
        'Description for Project C', 3, '2024-05-20', 0, 8);

-- Вставка тестовых данных для таблицы employee_project
INSERT INTO employee_project (project_id, employee_id)
VALUES (1, 1),
       (1, 2),
       (2, 3),
       (3, 4);
