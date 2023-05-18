package usecaseimpl

import (
	"context"
	"log"
	"url-shortener/internal/app/shortner/repository"
	"url-shortener/internal/app/shortner/usecase"
	"url-shortener/pkg/utils"
)

type shortnerUseCaseImpl struct {
	repo repository.ShortnerRepository
}

func NewShortnerUseCase(repo repository.ShortnerRepository) usecase.ShortnerUseCase {
	return &shortnerUseCaseImpl{
		repo: repo,
	}
}

func (s *shortnerUseCaseImpl) EncodeURL(ctx context.Context, url usecase.ShortURL) (usecase.ShortURL, error) {
	log.Println("shortnerUseCaseImpl EncodeURL()")
	encodedStr := utils.EncodeURL(url.URL)
	url.ShortURL = encodedStr
	resp, err := s.repo.Save(ctx, url)
	if err != nil {
		return usecase.ShortURL{}, err
	}
	resp.ShortURL = "http://short.est/" + url.ShortURL
	return resp, nil
}

func (s *shortnerUseCaseImpl) DecodeURL(ctx context.Context, url string) (usecase.ShortURL, error) {
	log.Println("shortnerUseCaseImpl DecodeURL()")
	shortURL, err := s.repo.Get(ctx, url)
	if err != nil {
		return usecase.ShortURL{}, err
	}
	decodeURL := utils.DecodeURL(shortURL.ShortURL)
	shortURL.ShortURL = decodeURL
	return shortURL, nil
}
