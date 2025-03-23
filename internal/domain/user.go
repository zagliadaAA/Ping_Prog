package domain

type User struct {
	ID       int64
	Username string
	ChatID   int64
}

func NewUser(username string, chatID int64) *User {
	return &User{
		Username: username,
		ChatID:   chatID,
	}
}
