package domain

import (
	"context"

	"be-project-monitoring/internal/domain/model"
	"be-project-monitoring/internal/domain/repository"

	"github.com/google/uuid"
)

type (
	Repository interface {
		userRepo
		projectRepo
		participantRepo
	}

	userRepo interface {
		GetUser(ctx context.Context, filter *repository.UserFilter) (*model.User, error)
		GetUsers(ctx context.Context, filter *repository.UserFilter) ([]model.User, error)
		GetCountByFilter(ctx context.Context, filter *repository.UserFilter) (int, error)

		InsertUser(ctx context.Context, user *model.User) error
		UpdateUser(ctx context.Context, user *model.User) error
		DeleteUser(ctx context.Context, id uuid.UUID) error
	}

	projectRepo interface {
		GetProject(ctx context.Context, filter *repository.ProjectFilter) (*model.Project, error)
		GetProjects(ctx context.Context, filter *repository.ProjectFilter) ([]model.Project, error)
		GetProjectCountByFilter(ctx context.Context, filter *repository.ProjectFilter) (int, error)

		InsertProject(ctx context.Context, project *model.Project) error
		UpdateProject(ctx context.Context, project *model.Project) error
		DeleteProject(ctx context.Context, id int) error
	}

	participantRepo interface {
		AddParticipant(ctx context.Context, participant *model.Participant) ([]model.Participant, error)
		GetParticipants(ctx context.Context, projectID int) ([]model.Participant, error)
	}
)
