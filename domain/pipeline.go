package domain

import (
	"github.com/usagiga/voicepipe/entity"
)

// PipelineImpl represents ...
type PipelineImpl struct {
}

// NewPipelineImpl is constructor of Pipeline
func NewPipelineImpl() Pipeline {
	return &PipelineImpl{}
}

func (p *PipelineImpl) Do(voice *entity.Voice) (err error) {
	panic("implement me")
}

func (p *PipelineImpl) RegisterDemuxer(demuxer Demuxer) (err error) {
	panic("implement me")
}

func (p *PipelineImpl) RegisterEffectors(effectors ...Effector) (err error) {
	panic("implement me")
}
