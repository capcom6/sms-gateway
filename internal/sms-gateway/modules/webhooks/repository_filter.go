package webhooks

import "gorm.io/gorm"

type SelectFilter func(*selectFilter)

func WithExtID(extID string) SelectFilter {
	return func(f *selectFilter) {
		f.extID = &extID
	}
}

func WithUserID(userID string) SelectFilter {
	return func(f *selectFilter) {
		f.userID = userID
	}
}

type selectFilter struct {
	userID string
	extID  *string
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
	query = query.Where("user_id = ?", f.userID)
	if f.extID != nil {
		query = query.Where("ext_id = ?", *f.extID)
	}
	return query
}
