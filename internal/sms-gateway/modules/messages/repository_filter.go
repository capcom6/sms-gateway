package messages

type MessagesSelectFilter struct {
	DeviceID string
}

type MessagesSelectOptions struct {
	WithRecipients bool
	WithDevice     bool
	WithStates     bool
}
