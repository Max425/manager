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
    company_id             int      not null references "company" on delete cascade,
    name                   text     not null,
    position               text,
    mail                   text unique,
    password               text     not null,
    salt                   text     not null,
    image                  text,
    rating                 integer[],
    active_projects_count  int               default 0,
    overdue_projects_count int               default 0,
    total_projects_count   int               default 0,
    created_at             timestamptz       default timezone('Europe/Moscow'::text, now()),
    updated_at             timestamptz       default timezone('Europe/Moscow'::text, now())
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