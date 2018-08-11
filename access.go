package access

import (
	"github.com/aufaitio/data-access/models"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// JobDAO specifies the interface of the job DAO needed by JobService.
type JobDAO interface {
	// Get returns the job with the specified job ID.
	Get(db *mongo.Database, id int64) (*models.Job, error)
	// GetByName returns the job with the specified job Name.
	GetByName(db *mongo.Database, name string) (*models.Job, error)
	// Count returns the number of repositories.
	Count(db *mongo.Database) (int64, error)
	// Query returns the list of repositories with the given offset and limit.
	Query(db *mongo.Database, offset, limit int) ([]*models.Job, error)
	// Create saves a new job in the storage.
	Create(db *mongo.Database, job *models.Job) error
	// Update updates the job with given ID in the storage.
	Update(db *mongo.Database, id int64, job *models.Job) error
	// Delete removes the job with given ID from the storage.
	Delete(db *mongo.Database, id int64) error
}

// RepositoryDAO specifies the interface of the repository DAO needed by RepositoryService.
type RepositoryDAO interface {
	// Get returns the repository with the specified repository ID.
	Get(db *mongo.Database, id int64) (*models.Repository, error)
	// Count returns the number of repositories.
	Count(db *mongo.Database) (int64, error)
	// Query returns the list of repositories with the given offset and limit.
	Query(db *mongo.Database, offset, limit int) ([]*models.Repository, error)
	// Query returns the list of repositories with the given offset and limit.
	QueryByDependency(db *mongo.Database, dependencyName string) ([]*models.Repository, error)
	// Create saves a new repository in the storage.
	Create(db *mongo.Database, repository *models.Repository) error
	// Update updates the repository with given ID in the storage.
	Update(db *mongo.Database, id int64, repository *models.Repository) error
	// Delete removes the repository with given ID from the storage.
	Delete(db *mongo.Database, id int64) error
}
