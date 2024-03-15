drop table if exists company cascade;
drop table if exists employee cascade;
drop table if exists project cascade;
drop table if exists employee_project cascade;

drop function if exists decrease_active_project_count_trigger;
drop function if exists increase_project_count_trigger;