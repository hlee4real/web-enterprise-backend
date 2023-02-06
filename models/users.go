package models

import "time"

type UsersModel struct {
	Username     string    `bson:"username" json:"username"`
	Password     string    `bson:"password" json:"password"`
	DateOfBirth  string    `bson::"date_of_birth" json:"date_of_birth"`
	Mobile       string    `bson:"mobile" json:"mobile"`
	Role         string    `bson:"role" json:"role"`
	Image        string    `bson:"image" json:"image"`
	DepartmentId string    `bson:"department_id" json:"department_id"`
	CreatedAt    time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time `bson:"updated_at" json:"updated_at"`
}
