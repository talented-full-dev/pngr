// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"fmt"
	"time"
)

type PostStatus string

const (
	PostStatusDraft     PostStatus = "draft"
	PostStatusPublished PostStatus = "published"
)

func (e *PostStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PostStatus(s)
	case string:
		*e = PostStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for PostStatus: %T", src)
	}
	return nil
}

type UserStatus string

const (
	UserStatusDisabled   UserStatus = "disabled"
	UserStatusUnverified UserStatus = "unverified"
	UserStatusActive     UserStatus = "active"
)

func (e *UserStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserStatus(s)
	case string:
		*e = UserStatus(s)
	default:
		return fmt.Errorf("unsupported scan type for UserStatus: %T", src)
	}
	return nil
}

type Post struct {
	ID        int64      `json:"id"`
	AuthorID  int64      `json:"author_id"`
	Title     string     `json:"title"`
	Body      string     `json:"body"`
	Status    PostStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type Reset struct {
	UserID    int64     `json:"user_id"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	ID           int64      `json:"id"`
	Email        string     `json:"email"`
	Pass         string     `json:"pass"`
	Salt         string     `json:"salt"`
	Status       UserStatus `json:"status"`
	Verification string     `json:"verification"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}
