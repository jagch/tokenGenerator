package application

import (
	"context"
	"jagch/tokenGenerator/domain"
	"jagch/tokenGenerator/model"
	"runtime"
	"sync"

	"github.com/google/uuid"
)

type TokenService interface {
	GenerateTokens(string, uint64) (*model.Tokens, error)
	CheckToken(string, string) bool
	createToken(int, chan string, chan model.ChToken)
}

type tokenService struct {
	tokenRepo domain.TokenRepository
}

func NewTokenService(tokenRepo domain.TokenRepository) TokenService {
	return &tokenService{
		tokenRepo: tokenRepo,
	}
}

func (s *tokenService) GenerateTokens(whitelabelName string, quantity uint64) (*model.Tokens, error) {
	var tokens model.Tokens

	// Concurrent process
	// Channels
	chSendData := make(chan string)
	chToken := make(chan model.ChToken)

	// Workers based on cpu numbers
	numWorkers := runtime.NumCPU()-1
	
	// Optional: We use this line or we can also remove and use all available cores by default
	runtime.GOMAXPROCS(numWorkers)

	// Goroutines 
	go sendData(whitelabelName, quantity, chSendData)
	go s.createToken(numWorkers, chSendData, chToken)

	// Response of the goroutines
	for tokenCreated := range chToken {
		if tokenCreated.Error != nil {
			return nil, tokenCreated.Error
		}
		tokens.Tokens = append(tokens.Tokens, tokenCreated.Token)
	}

	return &tokens, nil
}

func (s *tokenService) CheckToken(token, whitelabelName string) bool {
	ctx := context.Background()
	key := whitelabelName + ":" + token

	return s.tokenRepo.TokenExists(ctx, key)
}

func sendData(whitelabelName string, quantity uint64, chSendData chan string) { 
	for i := 0; i < int(quantity); i++ {
		chSendData <- whitelabelName
	}
	close(chSendData)
}

func (s *tokenService) createToken(numWorkers int, chSendData chan string, chToken chan model.ChToken) {
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {		
		// Goroutine
		go func() {
			var token model.Token
			var tokens model.Tokens
			for whitelabelName := range chSendData {
				// Create the token: uuid
				uuid := uuid.New()

				// Fill tokens
				token.Token = uuid.String()
				tokens.Tokens = append(tokens.Tokens, token)

				// Fill args
				key := whitelabelName + ":" + token.Token
				err := s.tokenRepo.CreateTokens(key, true)
				chToken <- model.ChToken{
					Token: token,
					Error: err,
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(chToken)
}
