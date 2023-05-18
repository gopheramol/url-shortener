package mock

import (
	"context"
	"errors"
	"url-shortener/internal/app/shortner/usecase"
)

type MockShortnerUseCase struct {
	EncodeURLFunc func(ctx context.Context, url usecase.ShortURL) (usecase.ShortURL, error)
	DecodeURLFunc func(ctx context.Context, url string) (usecase.ShortURL, error)
}

func (m *MockShortnerUseCase) EncodeURL(ctx context.Context, url usecase.ShortURL) (usecase.ShortURL, error) {
	if m.EncodeURLFunc != nil {
		return m.EncodeURLFunc(ctx, url)
	}
	return usecase.ShortURL{}, nil
}

func (m *MockShortnerUseCase) DecodeURL(ctx context.Context, url string) (usecase.ShortURL, error) {
	if m.DecodeURLFunc != nil {
		return m.DecodeURLFunc(ctx, url)
	}
	return usecase.ShortURL{}, nil
}

// repo

type MockShortnerRepository struct {
	SaveFunc func(ctx context.Context, url usecase.ShortURL) (usecase.ShortURL, error)
	GetFunc  func(ctx context.Context, url string) (usecase.ShortURL, error)
}

func (m *MockShortnerRepository) Save(ctx context.Context, url usecase.ShortURL) (usecase.ShortURL, error) {
	if m.SaveFunc != nil {
		return m.SaveFunc(ctx, url)
	}
	return usecase.ShortURL{}, errors.New("SaveFunc is not implemented")
}

func (m *MockShortnerRepository) Get(ctx context.Context, url string) (usecase.ShortURL, error) {
	if m.GetFunc != nil {
		return m.GetFunc(ctx, url)
	}
	return usecase.ShortURL{}, errors.New("GetFunc is not implemented")
}
