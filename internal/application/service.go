package application

import (
	"context"
	"fmt"
	"time"

	"github.com/gTauro8/git-toxagotchi/internal/domain"
	"github.com/gTauro8/git-toxagotchi/internal/infrastructure/storage"
)

type Service struct {
	store    *storage.SQLiteStore
	analyzer *Analyzer
	humor    *HumorEngine
}

func NewService(store *storage.SQLiteStore) *Service {
	return &Service{
		store:    store,
		analyzer: NewAnalyzer(),
		humor:    NewHumorEngine(),
	}
}

func (s *Service) GetOrCreatePet(name string) (*domain.Pet, error) {
	pet, err := s.store.GetPet()
	if err != nil || pet == nil {
		pet = domain.NewPet(name)
		if err := s.store.SavePet(pet); err != nil {
			return nil, fmt.Errorf("save pet: %w", err)
		}
	}
	return pet, nil
}

func (s *Service) GetPet() (*domain.Pet, error) {
	return s.store.GetPet()
}

func (s *Service) SavePet(pet *domain.Pet) error {
	return s.store.SavePet(pet)
}

func (s *Service) ProcessCommit(ctx context.Context, pet *domain.Pet, message, diff string) (string, error) {
	_ = ctx
	analysis := s.analyzer.AnalyzeCommit(message, diff)

	event := &domain.Event{
		Type:      domain.EventCommit,
		Timestamp: time.Now(),
		Metadata: map[string]interface{}{
			"message": message,
			"quality": analysis.MessageQuality,
		},
	}

	if analysis.SecretsDetected {
		pet.ApplyBadEvent(20, 5, 15)
		event.Type = domain.EventSecretDetected
		_ = s.store.SaveEvent(event)
		return s.humor.GetResponse("security"), nil
	}

	if analysis.MessageQuality >= 6 {
		pet.ApplyGoodEvent(10, 5, 15)
	} else {
		pet.ApplyBadEvent(5, 3, 5)
	}

	if analysis.DepsAdded > 0 {
		pet.ApplyBadEvent(3, 2, 3)
	}

	if newStage, ok := domain.ShouldEvolve(pet); ok {
		pet.Stage = newStage
	}

	pet.LastInteractionAt = time.Now()
	if err := s.store.SavePet(pet); err != nil {
		return "", err
	}
	if err := s.store.SaveEvent(event); err != nil {
		return "", err
	}

	return s.humor.GetResponse("commit"), nil
}

func (s *Service) FeedPet(pet *domain.Pet) (string, error) {
	pet.Feed()
	pet.LastInteractionAt = time.Now()
	if err := s.store.SavePet(pet); err != nil {
		return "", err
	}
	return "Pet munches happily. Hunger reduced.", nil
}

func (s *Service) PlayWithPet(pet *domain.Pet) (string, error) {
	pet.Play()
	pet.LastInteractionAt = time.Now()
	if err := s.store.SavePet(pet); err != nil {
		return "", err
	}
	return s.humor.GetMoodComment(pet.Mood), nil
}

func (s *Service) GetAchievements() ([]domain.Achievement, error) {
	return s.store.GetAchievements()
}
