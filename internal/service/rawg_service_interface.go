package rawgservice

import (
	"context"
)

type RawgServiceInterface interface {
	GetGameDetails(ctx context.Context, rawgId string) (GameResult, error)
}
