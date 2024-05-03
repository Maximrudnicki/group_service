package repository

import (
	"context"
	"group_service/cmd/model"
)

type GroupRepository interface {
	AddStudent(ctx context.Context, userId uint32, groupId string) error
	CreateGroup(ctx context.Context, group model.Group) error
	DeleteGroup(ctx context.Context, groupId string) error
	FindById(ctx context.Context, groupId string) (model.Group, error)
	FindByStudentId(ctx context.Context, userId uint32) ([]model.Group, error)
	FindByTeacherId(ctx context.Context, userId uint32) ([]model.Group, error)
	IsStudent(ctx context.Context, userId uint32, groupId string) bool
	IsTeacher(ctx context.Context, userId uint32, groupId string) bool
	RemoveStudent(ctx context.Context, userId uint32, groupId string) error
}
