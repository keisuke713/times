package times

type NotAuthedError struct {
	s string
}

func (e *NotAuthedError) Error() string {
	return e.s
}

func (e *NotAuthedError) AddPreText(preText string) *NotAuthedError {
	e.s = (preText + e.s)
	return e
}

type ChannelNotFoundError struct {
	s string
}

func (e *ChannelNotFoundError) Error() string {
	return e.s
}

func (e *ChannelNotFoundError) AddPreText(preText string) *ChannelNotFoundError {
	e.s = (preText + e.s)
	return e
}
