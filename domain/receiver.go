package domain

// ReceiverImpl represents ...
type ReceiverImpl struct {
}

// NewReceiverImpl is constructor of Receiver
func NewReceiverImpl() Receiver {
	return &ReceiverImpl{}
}

func (r *ReceiverImpl) StartReceive() (err error) {
	panic("implement me")
}

func (r *ReceiverImpl) StopReceive() (err error) {
	panic("implement me")
}

func (r *ReceiverImpl) IsReceiving() (isReceiving bool) {
	panic("implement me")
}

func (r *ReceiverImpl) RegisterMuxer(muxer Muxer) (err error) {
	panic("implement me")
}
