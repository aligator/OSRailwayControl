package bindata

import (
	"github.com/markbates/pkger"
)

func Asset(path string) ([]byte, error) {
	f, err := pkger.Open(path)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	info, err := f.Stat()
	if err != nil {
		return nil, err
	}

	data := make([]byte, info.Size())
	_, err = f.Read(data)

	return data, err
}
