package user

import (
	"database/sql"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUsers() ([]*User, error) {
	rows, err := s.db.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	users := make([]*User, 0)
	for rows.Next() {
		user, err := scanIntoUser(rows)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *Store) GetUserByID(id int) (*User, error) {
	row := s.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	if row == nil {
		return nil, sql.ErrNoRows
	}

	user, err := scanIntoUser(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// RowScanner is an interface that both sql.Row and sql.Rows can satisfy.
type RowScanner interface {
	Scan(dest ...interface{}) error
}

// scanIntoUser maps the result from RowScanner to a User object.
func scanIntoUser(rs RowScanner) (*User, error) {
	user := new(User)
	err := rs.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
