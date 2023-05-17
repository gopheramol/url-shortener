package usecaseimpl

import (
	"context"
	"log"
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

func (s *shortnerUseCaseImpl) CreateShortURL(ctx context.Context, url usecase.ShortURL) (int, error) {
	log.Println("Inside shortnerUseCaseImpl CreateShortURL()")
	// Implement create short URL use case logic

	// Add some logic for short url
	// for test adding axbycz

	// url = usecase.ShortURL{
	// 	FullURL: "axbycz",
	// }

	// log.Println(url.FullURL)
	// str, err := s.repo.CreateShortURL(ctx, url)
	// if err != nil {
	// 	log.Println("Error getting short from db", err.Error())
	// }
	// fmt.Println(str)
	return 0, nil
}

func (s *shortnerUseCaseImpl) GetLongURL(ctx context.Context, id int) (usecase.ShortURL, error) {
	log.Println("shortnerUseCaseImpl GetShortURL()")
	// Implement get short URL use case logic
	return usecase.ShortURL{}, nil
}
