package user

import "time"

type User struct {
	UserID      string    `dynamodbav:"userId" json:"userId"` // Hash key
	Name        string    `dynamodbav:"name" json:"name"`
	Email       string    `dynamodbav:"email" json:"email"`
	DateCreated time.Time `dynamodbav:"dateCreated,omitempty" json:"dateCreated"`
}
