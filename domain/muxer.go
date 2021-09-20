package domain

import (
	"github.com/usagiga/voicepipe/entity"
)

// MuxerImpl represents ...
type MuxerImpl struct {
}

// NewMuxerImpl is constructor of Muxer
func NewMuxerImpl() Muxer {
	return &MuxerImpl{}
}

func (m *MuxerImpl) Handle(voice *entity.Voice) (err error) {
	panic("implement me")
}

func (m *MuxerImpl) RegisterPipeline(pipeline Pipeline) (err error) {
	panic("implement me")
}

func (m *MuxerImpl) RegisterReceivers(receivers ...Receiver) (err error) {
	panic("implement me")
}
