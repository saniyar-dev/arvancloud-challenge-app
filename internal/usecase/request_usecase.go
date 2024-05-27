package usecase

import (
	"arvancloud-challenge-app/internal/domain"
	"arvancloud-challenge-app/internal/repository"
	"time"
)

type RequestUseCase interface {
	LogRequest(ip, userAgent string) error
}

type requestUseCase struct {
	requestRepo repository.RequestRepository
}

func NewRequestUseCase(requestRepo repository.RequestRepository) RequestUseCase {
	return &requestUseCase{requestRepo: requestRepo}
}

func (uc *requestUseCase) LogRequest(ip, userAgent string) error {
	request := &domain.Request{
		IP:        ip,
		UserAgent: userAgent,
		Timestamp: time.Now(),
	}
	return uc.requestRepo.SaveRequest(request)
}
