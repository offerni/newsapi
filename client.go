package newsapi

import "errors"

type ClientOpts struct {
	ApiKey string
}

func (opts ClientOpts) Validate() error {
	if len(opts.ApiKey) != 0 {
		return nil
	}

	return errors.New("No API Key in ClientOpts; ApiKey is required")
}

func NewClient(opts ClientOpts) (*ClientOpts, error) {
	if err := opts.Validate(); err != nil {
		return nil, err
	}

	return &ClientOpts{opts.ApiKey}, nil
}
