package access

import (
	"github.com/Sirupsen/logrus"
	"github.com/aufaitio/data-access/models"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Scope - scope of request that initiated query
type Scope struct {
	Logger *logrus.Logger
	DB     *mongo.Database
}

// NewScope provides an interface for the data access layer to interact with the app layer
func NewScope(logger *logrus.Logger, db *mongo.Database) *Scope {
	return &Scope{Logger: logger, DB: db}
}

// JobDAO specifies the interface of the job DAO needed by JobService.
type JobDAO interface {
	// Get returns the job with the specified job ID.
	Get(rs Scope, id int64) (*models.Job, error)
	// GetByName returns the job with the specified job Name.
	GetByName(rs Scope, name string) (*models.Job, error)
	// Count returns the number of repositories.
	Count(rs Scope) (int64, error)
	// Query returns the list of repositories with the given offset and limit.
	Query(rs Scope, offset, limit int) ([]*models.Job, error)
	// Create saves a new job in the storage.
	Create(rs Scope, job *models.Job) error
	// Update updates the job with given ID in the storage.
	Update(rs Scope, id int64, job *models.Job) error
	// Delete removes the job with given ID from the storage.
	Delete(rs Scope, id int64) error
}

// RepositoryDAO specifies the interface of the repository DAO needed by RepositoryService.
type RepositoryDAO interface {
	// Get returns the repository with the specified repository ID.
	Get(rs Scope, id int64) (*models.Repository, error)
	// Count returns the number of repositories.
	Count(rs Scope) (int64, error)
	// Query returns the list of repositories with the given offset and limit.
	Query(rs Scope, offset, limit int) ([]*models.Repository, error)
	// Query returns the list of repositories with the given offset and limit.
	QueryByDependency(rs Scope, dependencyName string) ([]*models.Repository, error)
	// Create saves a new repository in the storage.
	Create(rs Scope, repository *models.Repository) error
	// Update updates the repository with given ID in the storage.
	Update(rs Scope, id int64, repository *models.Repository) error
	// Delete removes the repository with given ID from the storage.
	Delete(rs Scope, id int64) error
}
