package repository

import (
	"database/sql"
	"errors"
	"gymbro-api/auth"
	"gymbro-api/model"
	"time"
)

type SessionRepository struct {
	db *sql.DB
}

var ErrSessionNotFound = errors.New("session not found")
var ErrSessionExpired = errors.New("session expired")
var ErrSessionInactive = errors.New("session manually inactivated")
var ErrInvalidPassword = errors.New("invalid password")

func NewSessionRepository(db *sql.DB) *SessionRepository {
	return &SessionRepository{db}
}

func (sr *SessionRepository) ValidateSession(token string) (*model.Session, error) {
	row := sr.db.QueryRow(`
		select s.id
		      ,s.user_id
			  ,s.token
			  ,s.created_at
			  ,s.updated_at
			  ,s.status
			  ,s.expires_at
		 from gymbro.sessions s 
		where s.token = $1::text
	`, token)

	session := &model.Session{}

	err := row.Scan(&session.ID, &session.UserID, &session.Token, &session.CreatedAt, &session.UpdatedAt)

	if err == sql.ErrNoRows {
		return nil, ErrSessionNotFound
	}

	if session.Status == "I" {
		return nil, ErrSessionInactive
	}

	now := time.Now()

	if session.ExpiresAt.Before(now) {
		return nil, ErrSessionExpired
	}

	if err != nil {
		return nil, err
	}

	return session, err
}

func (sr *SessionRepository) Login(login *model.Login) (*model.Session, error) {
	row := sr.db.QueryRow(`
		select u.id
		      ,case when crypt($2::text, u.password) = u.password then true else false end as password_match
		  from gymbro.users u
		 where u.username = $1::text
		   and u.status = 'A'
	`, login.Username, login.Password)

	var userID uint64
	var passwordMatch bool

	err := row.Scan(&userID, &passwordMatch)

	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	if !passwordMatch {
		return nil, ErrInvalidPassword
	}

	row = sr.db.QueryRow(`
		select s.id
		      ,s.user_id
			  ,s.token
			  ,s.created_at
			  ,s.updated_at
			  ,s.status
			  ,s.expires_at
		  from gymbro.sessions s
		 where s.user_id = $1::bigint
		   and s.status = 'A'
		   and s.expires_at > current_timestamp
	`, userID)

	session := &model.Session{}

	err = row.Scan(&session.ID, &session.UserID, &session.Token, &session.CreatedAt, &session.UpdatedAt, &session.Status, &session.ExpiresAt)

	if err == sql.ErrNoRows {
		token, err := auth.GenerateToken(userID)

		if err != nil {
			return nil, err
		}

		session.UserID = userID
		session.Token = token
		session.CreatedAt = time.Now()
		session.UpdatedAt = time.Now()
		session.Status = "A"
		session.ExpiresAt = time.Now().Add(time.Hour * 24)

		_, err = sr.db.Exec(`
			insert into gymbro.sessions(id
									   ,user_id
									   ,token
									   ,created_at
									   ,updated_at
									   ,status
									   ,expires_at)
								values (nextval('gymbro.seq_sessions')
									   ,$1::bigint
									   ,$2::text
									   ,$3::timestamp
									   ,$4::timestamp
									   ,$5::bpchar
									   ,$6::timestamp)
		`, session.UserID, session.Token, session.CreatedAt, session.UpdatedAt, session.Status, session.ExpiresAt)

		if err != nil {
			return nil, err
		}

		return session, nil

	}

	if err != nil {
		return nil, err
	}

	return session, nil
}
