package effectors

import (
	"github.com/usagiga/voicepipe/domain"
	"github.com/usagiga/voicepipe/entity"
)

// VolumeEffector represents ...
type VolumeEffector struct {
}

// NewVolumeEffectorImpl is constructor of VolumeEffector
func NewVolumeEffectorImpl() domain.Effector {
	return &VolumeEffector{}
}

func (e *VolumeEffector) Do(target *entity.Voice) (result *entity.Voice, err error) {
	panic("implement me")
}

func (e *VolumeEffector) Configure(conf interface{}) (err error) {
	panic("implement me")
}

func (e *VolumeEffector) ConfigureJSON(confJson string) (err error) {
	panic("implement me")
}