package models

import (
	"database/sql"
)

type User struct {
	ID uint
	Email string
	PasswordHash string
}

type UserService struct {
	DB *sql.DB
}

//Three options to pass a user. Either use User struct (useful when there is a 1:1 relation to DB), or use a NewUser and copy what is needed (for bigger projects) or use the input parameters as string (the usual way - should be used when not a lot of stuff is changing)

//PasswordHash will contain the plain password
// My preferred way: func (us *UserService) Create(user *User) error {
func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)
	//Hash the password
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, fmt.ErrorF("create user: %w", err)
	}
	passwordHash := string(hashedBytes)
	
	//Use the struct to return the user
	user := User{
		Email: email,
		PasswordHash: passwordHash,
	}
	
	row := us.DB.QueryRow("
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2) RETURNING id", email, passwordHash)
	var id int
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return &user, nil //Return the pointer reference
}

func (us *UserService) Update(user *User) error {
	//TODO implement
}
