package tests

import (
	"context"
	"group_service/cmd/model"
	"group_service/cmd/repository"
	"testing"

	"github.com/stretchr/testify/assert"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGroupRepository(t *testing.T) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		t.Fatalf("error connecting to MongoDB: %v", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database("pardodb")
	collection := db.Collection("groups")

	repo := repository.NewGroupRepositoryImpl(collection)

	testGroup := model.Group{
		Title:     "Test Group",
		TeacherId: 123,
		Students:  []uint32{456, 789},
	}

	t.Run("Test Create Group", func(t *testing.T) {
		err := repo.CreateGroup(context.Background(), testGroup)
		assert.NoError(t, err, "Expected no error while saving the word")
	})
	
	t.Run("Test Find By Student Id", func(t *testing.T) {
		groups, err := repo.FindByStudentId(context.Background(), 789)
		assert.NoError(t, err, "Expected no error")
		assert.NotEmpty(t, groups, "Should return groups")
	})

	t.Run("Test IsStudent", func(t *testing.T) {
		res := repo.IsStudent(context.Background(), 789, "662e3a12b4b7cdbbc46ac880")

		assert.True(t, res)
		
		res = repo.IsStudent(context.Background(), 45, "662e3a12b4b7cdbbc46ac880")
		assert.False(t, res)

		res = repo.IsStudent(context.Background(), 45, "662e3a12b480")
		assert.False(t, res)
	})

	t.Run("Test IsTeacher", func(t *testing.T) {
		res := repo.IsTeacher(context.Background(), 123, "662e3a12b4b7cdbbc46ac880")
		assert.True(t, res)
		
		res = repo.IsTeacher(context.Background(), 456, "662e3a12b4b7cdbbc46ac880")
		assert.False(t, res)

		res = repo.IsTeacher(context.Background(), 45, "662e3a12b480")
		assert.False(t, res)
	})

	t.Run("Test RemoveStudent", func(t *testing.T) {
		res := repo.IsStudent(context.Background(), 456, "662e3a12b4b7cdbbc46ac880")

		assert.True(t, res)
		
		res = repo.IsStudent(context.Background(), 789, "662e3a12b4b7cdbbc46ac880")

		assert.True(t, res)

		err = repo.RemoveStudent(context.Background(), 456, "662e3a12b4b7cdbbc46ac880")
		assert.NoError(t, err, "Expected no error")

		res = repo.IsStudent(context.Background(), 456, "662e3a12b4b7cdbbc46ac880")

		assert.False(t, res)
	})

	t.Run("Test AddStudent", func(t *testing.T) {
		err := repo.AddStudent(context.Background(), 456, "662e3a12b4b7cdbbc46ac880")
		assert.NoError(t, err, "Expected no error")
		
		res := repo.IsStudent(context.Background(), 456, "662e3a12b4b7cdbbc46ac880")
		assert.True(t, res)
	})
}
