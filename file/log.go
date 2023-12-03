package file

import "os"

type LogFile struct {
	f *os.File
}

func (lf *LogFile) Close() error {
	if err := lf.f.Close(); err != nil {
		return err
	}
	return nil
}

func (lf *LogFile) Write(bytes []byte) error {
	return nil
}

type Options struct {
	name string
}

func OpenLogFile(opt *Options) *LogFile {
	lf := &LogFile{}
	lf.f, _ = os.Create(opt.name)
	return lf
}
