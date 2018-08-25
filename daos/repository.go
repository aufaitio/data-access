package daos

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
	"github.com/quantumew/data-access/models"
	"golang.org/x/net/context"
)

// RepositoryDAO persists repository data in database
type repositoryDAO struct{}

// NewRepositoryDAO creates a new RepositoryDAO
func NewRepositoryDAO() *repositoryDAO {
	return &repositoryDAO{}
}

func (dao repositoryDAO) get(db *mongo.Database, doc *bson.Document) (*models.Repository, error) {
	var repository *models.Repository
	col := db.Collection("repository")

	err := col.FindOne(
		context.Background(),
		doc,
	).Decode(repository)

	if err != nil {
		return repository, err
	}

	return repository, err
}

// Get reads the repository with the specified name from the database.
func (dao repositoryDAO) Get(db *mongo.Database, name string) (*models.Repository, error) {
	return dao.get(
		db,
		bson.NewDocument(
			bson.EC.String("name", name),
		),
	)
}

// Create saves a new repository record in the database.
func (dao repositoryDAO) Create(db *mongo.Database, repository *models.Repository) error {
	col := db.Collection("repository")
	repoBson := models.NewDocFromRepository(repository)

	_, err := col.InsertOne(
		context.Background(),
		repoBson,
	)

	return err
}

// Update saves the changes to an repository in the database.
func (dao repositoryDAO) Update(db *mongo.Database, name string, repository *models.Repository) error {
	if _, err := dao.Get(db, name); err != nil {
		return err
	}

	repoBson := models.NewDocFromRepository(repository)
	col := db.Collection("repository")

	_, err := col.UpdateOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("name", name),
		),
		repoBson,
	)
	return err
}

// Delete deletes an repository with the specified name from the database.
func (dao repositoryDAO) Delete(db *mongo.Database, name string) error {
	repository, err := dao.Get(db, name)
	if err != nil {
		return err
	}

	col := db.Collection("repository")
	_, err = col.DeleteOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.String("name", repository.Name),
		),
	)

	return err
}

// Count returns the number of the repository records in the database.
func (dao repositoryDAO) Count(db *mongo.Database) (int64, error) {
	return db.Collection("repository").Count(
		context.Background(),
		bson.NewDocument(),
	)
}

// Query retrieves the repository records with the specified offset and limit from the database.
func (dao repositoryDAO) Query(db *mongo.Database, offset, limit int) ([]*models.Repository, error) {
	return dao.query(db, offset, limit, bson.NewDocument())
}

// QueryByDependency queries by dependency.
func (dao repositoryDAO) QueryByDependency(db *mongo.Database, dependencyName string) ([]*models.Repository, error) {
	return dao.query(db, 0, 0, bson.NewDocument(
		bson.EC.SubDocumentFromElements("dependencies",
			bson.EC.ArrayFromElements("$in",
				bson.VC.DocumentFromElements(bson.EC.String("name", dependencyName)),
			),
		),
	))
}

// QueryByName queries for repositories by name.
func (dao repositoryDAO) QueryByName(db *mongo.Database, nameList []string) ([]*models.Repository, error) {
	var queryList *bson.Array

	for _, name := range nameList {
		queryList.Append(bson.VC.String(name))
	}

	return dao.query(db, 0, 0, bson.NewDocument(
		bson.EC.SubDocumentFromElements("name",
			bson.EC.Array("$in", queryList),
		),
	))
}

// Query retrieves the repository records with the specified offset and limit from the database.
func (dao repositoryDAO) query(db *mongo.Database, offset, limit int, filter *bson.Document) ([]*models.Repository, error) {
	var (
		cursor mongo.Cursor
		err    error
	)
	repositoryList := []*models.Repository{}
	col := db.Collection("repository")
	ctx := context.Background()

	if limit > 0 {
		cursor, err = col.Find(
			ctx,
			filter,
			findopt.Limit(int64(limit)),
			findopt.Skip(int64(offset)),
		)
	} else {
		cursor, err = col.Find(
			ctx,
			filter,
			findopt.Skip(int64(offset)),
		)
	}

	if err != nil {
		return repositoryList, err
	}

	defer cursor.Close(ctx)
	elm := bson.NewDocument()

	for cursor.Next(ctx) {
		elm.Reset()

		if err := cursor.Decode(elm); err != nil {
			return repositoryList, err
		}
		job, err := models.NewRepositoryFromDoc(elm)

		if err != nil {
			return repositoryList, err
		}

		repositoryList = append(repositoryList, job)
	}

	return repositoryList, err
}
