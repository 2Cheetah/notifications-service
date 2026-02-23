package service

type Service struct {
	NovuClient NovuClient
}

func NewService(nc NovuClient) *Service {
	return &Service{
		NovuClient: nc,
	}
}
