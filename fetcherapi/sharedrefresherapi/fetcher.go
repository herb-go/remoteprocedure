package sharedrefresherapi

import (
	"bytes"

	"github.com/herb-go/fetcher"
)

type Fetcher struct {
	Server *fetcher.Server
}

func (f *Fetcher) RefreshShared(data []byte) ([]byte, error) {
	var result = []byte{}
	p, err := f.Server.CreatePreset()
	if err != nil {
		return nil, err
	}
	_, err = p.With(fetcher.Method("POST")).FetchWithBodyAndParse(bytes.NewBuffer(data), fetcher.Should200(fetcher.AsBytes(&result)))
	if err != nil {
		return nil, err
	}
	return result, nil
}
