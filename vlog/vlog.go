package vlog

import (
	"github.com/sjmshsh/hopeKV/utils"
	"github.com/sjmshsh/hopeKV/utils/codec"
)

type Options struct {
}

type VLog struct {
	closer *utils.Closer
}

func (v *VLog) Close() error {
	return nil
}

func NewVLog(opt *Options) *VLog {
	v := &VLog{}
	v.closer = utils.NewCloser(1)
	return v
}

func (v *VLog) StartGC() {
	defer v.closer.Done()
	for {
		select {
		case <-v.closer.Wait():
		}
	}
}

func (v *VLog) Set(entry *codec.Entry) error {
	return nil
}

func (v *VLog) Get(entry *codec.Entry) (*codec.Entry, error) {
	return nil, nil
}
