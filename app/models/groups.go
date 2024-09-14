package models

import (
	"github.com/google/uuid"
)

type Group struct {
	ID      uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name    string
	Members []GroupMember `gorm:"many2many:group_members;"`
}

type GroupMember struct {
	ID    uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name  string
	Email string
}

type GroupResponse struct {
	Name    string                `json:"name"`
	Members []GroupMemberResponse `json:"members"`
}

type GroupMemberResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
