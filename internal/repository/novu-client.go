package repository

type NovuClient struct {
	baseUrl string
}

func NewNovuClient(baseUrl string) *NovuClient {
	return &NovuClient{
		baseUrl: baseUrl,
	}
}
