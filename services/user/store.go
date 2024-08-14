package user

import "time"

type Store struct {
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) GetUsers() []User {
	// Create some fake user data and return
	users := []User{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "john.doe@example.com",
			Password:  "password123",
			CreatedAt: time.Now(),
		},
		{
			ID:        2,
			FirstName: "Jane",
			LastName:  "Smith",
			Email:     "jane.smith@example.com",
			Password:  "password123",
			CreatedAt: time.Now(),
		},
		{
			ID:        3,
			FirstName: "Alice",
			LastName:  "Johnson",
			Email:     "alice.johnson@example.com",
			Password:  "password123",
			CreatedAt: time.Now(),
		},
	}

	return users
}

func (s *Store) GetUserByID(id int) User {
	// Create some fake user data and return
	user := User{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		Password:  "password123",
		CreatedAt: time.Now(),
	}

	return user
}
