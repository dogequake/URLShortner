package service

import (
	"URLShortner/repository"
	"errors"
	"math/rand"
	"net/url"
	"time"
)

const shortIDLength = 6

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func init() {
	rand.Seed(time.Now().UnixNano())
}

type URLService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{repo: repo}
}

func (s *URLService) ValidateURL(rawurl string) error {
	u, err := url.ParseRequestURI(rawurl)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return errors.New("invalid url")
	}
	return nil
}

func (s *URLService) GenerateShortID() string {
	b := make([]rune, shortIDLength)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (s *URLService) SaveURL(longURL string) (string, error) {
	for i := 0; i < 5; i++ {
		shortID := s.GenerateShortID()
		if _, exists := s.repo.Find(shortID); !exists {
			err := s.repo.Save(shortID, longURL)
			return shortID, err
		}
	}
	return "", errors.New("could not generate unique short id")
}

func (s *URLService) GetLongURL(shortID string) (string, bool) {
	return s.repo.Find(shortID)
}
