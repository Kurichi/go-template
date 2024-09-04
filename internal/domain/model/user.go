package model

import "github.com/oklog/ulid/v2"

type UserID ulid.ULID
type User struct {
	ID   UserID
	Name string
}

func NewUserID() (UserID, error) {
	id, err := ulid.New(ulid.Now(), nil)
	if err != nil {
		return UserID{}, err
	}
	return UserID(id), nil
}

func (id UserID) String() string {
	return ulid.ULID(id).String()
}

func NewUser(name string) (*User, error) {
	id, err := NewUserID()
	if err != nil {
		return nil, err
	}

	user := &User{
		ID:   id,
		Name: name,
	}
	return user, nil
}
