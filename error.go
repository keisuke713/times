package times

type NotAuthedError struct {
	s string
}

func (e *NotAuthedError) Error() string {
	return e.s
}

type ChannelNotFoundError struct {
	s string
}

func (e *ChannelNotFoundError) Error() string {
	return e.s
}
