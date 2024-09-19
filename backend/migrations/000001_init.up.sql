create table sessions
(
    sid        text primary key,
    role       int         not null,
    company_id int         not null,
    expiration timestamptz not null
);

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
    rating                 decimal,
    active_projects_count  int               default 0,
    overdue_projects_count int               default 0,
    total_projects_count   int               default 0,
--     role                   smallint not null default 0 check (role between 0 and 1 ),
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


update project
set image = 'https://fikiwiki.com/uploads/posts/2022-02/1645039733_10-fikiwiki-com-p-kartinki-logotipov-10.jpg'
where image = 'https://gas-kvas.com/grafic/uploads/posts/2024-01/gas-kvas-com-p-simvoli-dlya-logotipov-na-prozrachnom-fone-34.png';

update employee
set image = 'https://sun9-35.userapi.com/impg/163SluMCn4S9UKiRrQcDDsNXrQ5-pQ6cdq9lPA/Bsn8hG3-Pxc.jpg?size=1024x683&quality=96&sign=29394fc81284454fe65275574c9ad772&c_uniq_tag=2MaoO9ls7_6c-BOT7hqX_Nzkr5o7K5nGt6_jqcC-SV0&type=album'
where image = 'https://get.pxhere.com/photo/outdoor-person-girl-sun-woman-hair-white-photography-cute-summer-female-portrait-model-young-red-fashion-lady-facial-expression-hairstyle-smiling-smile-long-hair-close-up-caucasian-face-dress-happy-happiness-eye-head-skin-beauty-attractive-photo-shoot-pretty-girl-brown-hair-cute-girl-happy-girl-happy-woman-portrait-photography-supermodel-683657.jpg';

delete from employee
where id = 45;

delete from project
where id = 5;