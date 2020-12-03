package qywxapi

func CGIBin(opts ...option) *cgiBin {
	self := &cgiBin{}
	for _, opt := range opts {
		opt(self)
	}
	return self
}
