package devices

import "gorm.io/gorm"

type SelectFilter func(*selectFilter)

func WithID(id string) SelectFilter {
	return func(f *selectFilter) {
		f.id = &id
	}
}

func WithToken(token string) SelectFilter {
	return func(f *selectFilter) {
		f.token = &token
	}
}

func WithUserID(userID string) SelectFilter {
	return func(f *selectFilter) {
		f.userID = &userID
	}
}

type selectFilter struct {
	id     *string
	userID *string
	token  *string
}

func newFilter(filters ...SelectFilter) *selectFilter {
	f := &selectFilter{}
	f.merge(filters...)
	return f
}

func (f *selectFilter) merge(filters ...SelectFilter) {
	for _, filter := range filters {
		filter(f)
	}
}

func (f *selectFilter) apply(query *gorm.DB) *gorm.DB {
	if f.id != nil {
		query = query.Where("id = ?", *f.id)
	}
	if f.token != nil {
		query = query.Where("auth_token = ?", *f.token)
	}
	if f.userID != nil {
		query = query.Where("user_id = ?", *f.userID)
	}
	return query
}
