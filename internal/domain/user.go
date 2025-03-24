package domain

type User struct {
	ID       int
	Username string
	ChatID   int
}

func NewUser(username string, chatID int) *User {
	return &User{
		Username: username,
		ChatID:   chatID,
	}
}
