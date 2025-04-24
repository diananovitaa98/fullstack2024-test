package entity

import "time"

type MyClient struct {
	ID           int
	Name         string
	Slug         string
	IsProject    string
	SelfCapture  string
	ClientPrefix string
	ClientLogo   string
	Address      string
	PhoneNumber  string
	City         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
}
