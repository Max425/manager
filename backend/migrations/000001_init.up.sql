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
    id         serial primary key,
    company_id int  not null references "company" on delete cascade,
    name       text not null,
    position   text,
    mail       text unique,
    password   text not null,
    salt       text not null,
    image      text,
    rating     decimal,
    created_at timestamptz default timezone('Europe/Moscow'::text, now()),
    updated_at timestamptz default timezone('Europe/Moscow'::text, now())
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
