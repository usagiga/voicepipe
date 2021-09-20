package domain

import (
	"github.com/usagiga/voicepipe/entity"
)

// DemuxerImpl represents ...
type DemuxerImpl struct {
}

// NewDemuxerImpl is constructor of Demuxer
func NewDemuxerImpl() Demuxer {
	return &DemuxerImpl{}
}

func (d *DemuxerImpl) Send(voice *entity.Voice) (err error) {
	panic("implement me")
}

func (d *DemuxerImpl) RegisterSenders(senders ...Sender) (err error) {
	panic("implement me")
}
