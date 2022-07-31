package model

import "github.com/google/uuid"

type User struct {
	Base
	UserAPI
}

type UserAPI struct {
	Name           *string `json:"name,omitempty" gorm:"type:varchar(256);not null;" example:"Baron"`
	Occupation     *string `json:"occupation,omitempty" gorm:"type:varchar(256);not null;" example:"Programmer"`
	Email          *string `json:"email,omitempty" gorm:"type:varchar(256);not null;index:idx_email_deleted_at,unique;where:deleted_at is null" example:"email@email.com"`
	Password       *string `json:"password,omitempty" gorm:"type:varchar(256);not null;"`
	AvatarFileName *string `json:"avatar_file_name,omitempty" gorm:"type:varchar(256);" example:"avatar.jpg"`
	Role           *string `json:"role,omitempty" gorm:"type:varchar(256);not null;" example:"user"`
	Token          *string `json:"token,omitempty" gorm:"type:varchar(256);" example:"tokentoken"`
}

type UserData struct {
	ID             *uuid.UUID `json:"id,omitempty" gorm:"primaryKey;unique;type:varchar(36);not null" format:"uuid"`
	Name           *string    `json:"name,omitempty" gorm:"type:varchar(256);not null;" example:"Baron"`
	Occupation     *string    `json:"occupation,omitempty" gorm:"type:varchar(256);not null;" example:"Programmer"`
	Email          *string    `json:"email,omitempty" gorm:"type:varchar(256);not null;index:idx_email_deleted_at,unique;where:deleted_at is null" example:"email@email.com"`
	AvatarFileName *string    `json:"avatar_file_name,omitempty" gorm:"type:varchar(256);" example:"avatar.jpg"`
	Token          *string    `json:"token,omitempty" gorm:"type:varchar(256);" example:"tokentoken"`
}
