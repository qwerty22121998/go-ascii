package service

type Provider struct {
	ConvertService ConvertService
}

func NewProvider() *Provider {
	return &Provider{
		ConvertService: NewConvertService(),
	}
}
