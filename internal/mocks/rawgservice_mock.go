package mocks

import (
	"context"
	rawgservice "games-shelf-api-go/internal/service"

	"github.com/stretchr/testify/mock"
)

// MockRawgService is a mock implementation of the RawgServiceInterface
type MockRawgService struct {
	mock.Mock
}

// Ensure MockRawgService implements the RawgServiceInterface
var _ rawgservice.RawgServiceInterface = (*MockRawgService)(nil)

func (m *MockRawgService) GetGameDetails(ctx context.Context, rawgId string) (rawgservice.GameResult, error) {
	args := m.Called(ctx, rawgId)
	return args.Get(0).(rawgservice.GameResult), args.Error(1)
}
