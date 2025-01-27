package storage

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	d "github.com/SergeyBogomolovv/read-advisor/services/feed/internal/domain"
	"github.com/jmoiron/sqlx"
)

type storage struct {
	db *sqlx.DB
}

func NewStorage(db *sqlx.DB) *storage {
	return &storage{
		db: db,
	}
}

func (s *storage) SavePreference(ctx context.Context, userID int64, pref d.Preference, value string, priority int) error {
	var builder sq.InsertBuilder
	switch pref {
	case d.PreferenceAuthor:
		builder = sq.Insert("author_preferences").Columns("user_id", "name_author", "priority")
	case d.PreferenceCategory:
		builder = sq.Insert("category_preferences").Columns("user_id", "name_category", "priority")
	case d.PreferencePublisher:
		builder = sq.Insert("publisher_preferences").Columns("user_id", "name_publisher", "priority")
	}
	if _, err := builder.Values(userID, value, priority).RunWith(s.db).ExecContext(ctx); err != nil {
		return err
	}
	return nil
}
