package repository

import (
	"context"
	"group_service/cmd/model"
)

type StatisticsRepository interface {
	AddWordToStatistics(ctx context.Context, statId string, word uint32) error
	CreateStatistics(ctx context.Context, stat model.Statistics) error
	GetStatistics(ctx context.Context, statId string) (model.Statistics, error)
	GetId(ctx context.Context, groupId string, teacherId uint32, studentId uint32) (string, error)
	DeleteStatistics(ctx context.Context, statId string) error
}
