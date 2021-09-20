package domain

import "github.com/usagiga/voicepipe/entity"

type Receiver interface {
	StartReceive() (err error)
	StopReceive() (err error)
	IsReceiving() (isReceiving bool)
	RegisterMuxer(muxer Muxer) (err error)
}

type Sender interface {
	Send(voice *entity.Voice) (err error)
	IsSendingTo(target *entity.Room) (isSendingTo bool)
}

type Muxer interface {
	Handle(voice *entity.Voice) (err error)
	RegisterPipeline(pipeline Pipeline) (err error)
	RegisterReceivers(receivers ...Receiver) (err error)
}

type Demuxer interface {
	Send(voice *entity.Voice) (err error)
	RegisterSenders(senders ...Sender) (err error)
}

type Pipeline interface {
	Do(voice *entity.Voice) (err error)
	RegisterDemuxer(demuxer Demuxer) (err error)
	RegisterEffectors(effectors ...Effector) (err error)
}

type Effector interface {
	Do(target *entity.Voice) (result *entity.Voice, err error)
	Configure(conf interface{}) (err error)
	ConfigureJSON(confJson string) (err error)
}
