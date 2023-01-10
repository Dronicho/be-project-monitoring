package repository

import (
	"be-project-monitoring/internal/db"

	"github.com/google/uuid"
)

type UserFilter struct {
	IDs       []uuid.UUID `json:"id"`
	Usernames []string    `json:"username"`
	Emails    []string    `json:"email"`
	*db.Paginator
}

func NewUserFilter() *UserFilter {
	return &UserFilter{Paginator: db.DefaultPaginator}
}

func (f *UserFilter) ByIDs(ids ...uuid.UUID) *UserFilter {
	f.IDs = ids
	return f
}

func (f *UserFilter) ByUsernames(usernames ...string) *UserFilter {
	f.Usernames = usernames
	return f
}

func (f *UserFilter) ByEmails(emails ...string) *UserFilter {
	f.Emails = emails
	return f
}

func (f *UserFilter) WithPaginator(limit, offset uint64) *UserFilter {
	f.Paginator = db.NewPaginator(limit, offset)
	return f
}
