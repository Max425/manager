package repository

import (
	"context"
	"github.com/Max425/manager/internal/model/core"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type ProjectRepository struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewProjectRepository(db *sqlx.DB, log *slog.Logger) *ProjectRepository {
	return &ProjectRepository{db: db, log: log}
}

func (pr *ProjectRepository) CreateProject(ctx context.Context, project *core.Project) (*core.Project, error) {
	query := "INSERT INTO project (company_id, name, stages, image, description, current_stage, deadline, status, complexity) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	row := pr.db.QueryRowContext(ctx, query, project.CompanyID, project.Name, project.Stages, project.Image, project.Description, project.CurrentStage, project.Deadline, project.Status, project.Complexity)
	var id int
	err := row.Scan(&id)
	if err != nil {
		pr.log.Error("Error creating project", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	project.ID = id
	return project, nil
}

func (pr *ProjectRepository) FindProjectByID(ctx context.Context, id int) (*core.Project, error) {
	var project core.Project
	err := pr.db.GetContext(ctx, &project, "SELECT * FROM project WHERE id=$1", id)
	if err != nil {
		pr.log.Error("Error finding project", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	if project.ID == 0 {
		return nil, core.ErrNotFound
	}

	return &project, nil
}

func (pr *ProjectRepository) FindProjectsByCompanyID(ctx context.Context, companyID int) ([]*core.Project, error) {
	var projects []*core.Project
	err := pr.db.SelectContext(ctx, &projects, "SELECT * FROM project WHERE company_id=$1", companyID)
	if err != nil {
		pr.log.Error("Error finding projects by company ID", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	if len(projects) == 0 {
		return nil, core.ErrNotFound
	}
	return projects, nil
}

func (pr *ProjectRepository) GetProjectEmployees(ctx context.Context, id int) ([]*core.Employee, error) {
	var employees []*core.Employee
	err := pr.db.SelectContext(ctx, &employees, `select * from employee
							where id in (select employee_id from employee_project
             								where project_id = $1);`, id)
	if err != nil {
		pr.log.Error("Error finding projects by company ID", slog.String("error", err.Error()))
		return nil, core.ErrInternal
	}
	if len(employees) == 0 {
		return nil, core.ErrNotFound
	}
	return employees, nil
}

func (pr *ProjectRepository) UpdateProject(ctx context.Context, project *core.Project) (*core.Project, error) {
	_, err := pr.db.ExecContext(ctx, "UPDATE project SET company_id=$1, name=$2, stages=$3, image=$4, description=$5, current_stage=$6, deadline=$7, status=$8, complexity=$9 WHERE id=$10",
		project.CompanyID, project.Name, project.Stages, project.Image, project.Description, project.CurrentStage, project.Deadline, project.Status, project.Complexity, project.ID)
	if err != nil {
		pr.log.Error("Error updating project", slog.String("error", err.Error()))
		return nil, err
	}
	return pr.FindProjectByID(ctx, project.ID)
}

func (pr *ProjectRepository) DeleteProject(ctx context.Context, id int) error {
	_, err := pr.db.ExecContext(ctx, "DELETE FROM project WHERE id=$1", id)
	if err != nil {
		pr.log.Error("Error deleting project", slog.String("error", err.Error()))
		return core.ErrInternal
	}
	return nil
}
