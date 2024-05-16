package file

import (
	"fmt"
	"github.com/sjmshsh/hopeKV/utils"
	"os"
)

type MockFile struct {
	f *os.File
}

func (lf *MockFile) Close() error {
	if err := lf.f.Close(); err != nil {
		return err
	}
	return nil
}

func (lf *MockFile) Write(bytes []byte) (int, error) {
	return lf.f.Write(bytes)
}

func (lf *MockFile) Read(bytes []byte) (int, error) {
	return lf.f.Read(bytes)
}

type Options struct {
	Name string
	Dir  string
}

// OpenMockFile mock 文件
func OpenMockFile(opt *Options) *MockFile {
	var err error
	lf := &MockFile{}
	lf.f, err = os.Open(fmt.Sprintf("%s/%s", opt.Dir, opt.Name))
	utils.Panic(err)
	return lf
}
