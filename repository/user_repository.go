package repository

import (
	"database/sql"
	"gymbro-api/model"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (ur *UserRepository) CreateUser(user *model.User) error {
	_, err := ur.db.Exec(`
		insert into gymbro.users(id
			                    ,type
								,username
							    ,name
							    ,email
							    ,password)
						values (nextval('gymbro.seq_users')
						       ,$1
							   ,$2
							   ,$3
							   ,$4
							   ,$5)
	`, user.Type, user.Username, user.Name, user.Email, user.Password)

	return err
}

func (ur *UserRepository) GetUserByID(id int64) (*model.User, error) {
	row := ur.db.QueryRow(`
		select u.id
		      ,u.username
			  ,u.name
			  ,u.type
			  ,u.email
	          ,u.created_at
			  ,u.updated_at
		 from gymbro.users u 
		where u.id = ?
	`, id)

	user := &model.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Name, &user.Type, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	return user, err
}
