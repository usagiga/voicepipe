package domain

import (
	"github.com/usagiga/voicepipe/entity"
)

// EffectorImpl represents ...
type EffectorImpl struct {
}

// NewEffectorImpl is constructor of Effector
func NewEffectorImpl() Effector {
	return &EffectorImpl{}
}

func (e *EffectorImpl) Do(target *entity.Voice) (result *entity.Voice, err error) {
	panic("implement me")
}

func (e *EffectorImpl) Configure(conf interface{}) (err error) {
	panic("implement me")
}

func (e *EffectorImpl) ConfigureJSON(confJson string) (err error) {
	panic("implement me")
}
