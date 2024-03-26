package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type SessionData struct {
	Role      int
	CompanyID int
}

type SessionRepository struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewSessionRepository(db *sqlx.DB, log *slog.Logger) *SessionRepository {
	return &SessionRepository{
		db:  db,
		log: log,
	}
}

func (s *SessionRepository) SetSession(ctx context.Context, SID string, data SessionData, lifetime time.Duration) error {
	query := `INSERT INTO sessions (sid, role, company_id, expiration) VALUES ($1, $2, $3, $4) ON CONFLICT (sid) DO UPDATE SET role = $2, company_id = $3, expiration = $4`

	expiration := time.Now().Add(lifetime)

	_, err := s.db.ExecContext(ctx, query, SID, data.Role, data.CompanyID, expiration)
	if err != nil {
		return errors.Wrap(err, "failed to set session")
	}
	return nil
}

func (s *SessionRepository) GetSession(ctx context.Context, SID string) (SessionData, error) {
	var data SessionData

	query := `SELECT role, company_id FROM sessions WHERE sid = $1`

	err := s.db.QueryRowxContext(ctx, query, SID).Scan(&data.Role, &data.CompanyID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return SessionData{}, fmt.Errorf("session not found")
		}
		return SessionData{}, errors.Wrap(err, "failed to get session")
	}

	return data, nil
}

func (s *SessionRepository) DeleteSession(ctx context.Context, SID string) error {
	query := `DELETE FROM sessions WHERE sid = $1`

	_, err := s.db.ExecContext(ctx, query, SID)
	if err != nil {
		return errors.Wrap(err, "failed to delete session")
	}
	return nil
}
