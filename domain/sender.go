package domain

import (
	"github.com/usagiga/voicepipe/entity"
)

// SenderImpl represents ...
type SenderImpl struct {
}

// NewSenderImpl is constructor of Sender
func NewSenderImpl() Sender {
	return &SenderImpl{}
}

func (s *SenderImpl) Send(voice *entity.Voice) (err error) {
	panic("implement me")
}

func (s *SenderImpl) IsSendingTo(target *entity.Room) (isSendingTo bool) {
	panic("implement me")
}
