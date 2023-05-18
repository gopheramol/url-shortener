package usecaseimpl

import (
	"context"
	"url-shortener/internal/app/shortner/repository"
	"url-shortener/internal/app/shortner/usecase"
)

type shortnerUseCaseImpl struct {
	repo repository.ShortnerRepository
}

func NewShortnerUseCase(repo repository.ShortnerRepository) usecase.ShortnerUseCase {
	return &shortnerUseCaseImpl{
		repo: repo,
	}
}

func (s *shortnerUseCaseImpl) EncodeURL(ctx context.Context, url usecase.ShortURL) (int, error) {

	s.repo.Save(ctx, usecase.ShortURL{})
	return 0, nil
}

func (s *shortnerUseCaseImpl) DecodeURL(ctx context.Context, id int) (usecase.ShortURL, error) {
	// Implement get short URL use case logic
	return usecase.ShortURL{}, nil
}
