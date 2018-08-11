package daos

import (
	"github.com/aufaitio/data-access/models"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"golang.org/x/net/context"
)

// JobDAO persists job data in database
type JobDAO struct{}

// NewJobDAO creates a new JobDAO
func NewJobDAO() *JobDAO {
	return &JobDAO{}
}

// Get reads the job with the specified ID from the database.
func (dao *JobDAO) Get(db *mongo.Database, id int64) (*models.Job, error) {
	return dao.get(
		db,
		bson.NewDocument(
			bson.EC.Int64("id", id),
		),
	)
}

// GetByName reads the job with the specified name from the database.
func (dao *JobDAO) GetByName(db *mongo.Database, name string) (*models.Job, error) {
	return dao.get(
		db,
		bson.NewDocument(
			bson.EC.String("name", name),
		),
	)
}

func (dao *JobDAO) get(db *mongo.Database, doc *bson.Document) (*models.Job, error) {
	var job *models.Job
	col := db.Collection("job")
	result := bson.NewDocument()

	err := col.FindOne(
		context.Background(),
		doc,
	).Decode(result)

	if err != nil {
		return job, err
	}

	job, err = models.NewJobFromDoc(result)

	return job, err
}

// Create saves a new job record in the database.
// The Job.ID field will be populated with an automatically generated ID upon successful saving.
func (dao *JobDAO) Create(db *mongo.Database, job *models.Job) error {
	col := db.Collection("job")

	jobBson := models.NewDocFromJob(job)
	_, err := col.InsertOne(
		context.Background(),
		jobBson,
	)

	return err
}

// Update saves the changes to an job in the database.
func (dao *JobDAO) Update(db *mongo.Database, id int64, job *models.Job) error {
	if _, err := dao.Get(db, id); err != nil {
		return err
	}

	jobBson := models.NewDocFromJob(job)
	col := db.Collection("job")
	_, err := col.UpdateOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.Int64("_id", job.ID),
		),
		jobBson,
	)

	return err
}

// Delete deletes an job with the specified ID from the database.
func (dao *JobDAO) Delete(db *mongo.Database, id int64) error {
	_, err := dao.Get(db, id)
	if err != nil {
		return err
	}

	col := db.Collection("job")
	_, err = col.DeleteOne(
		context.Background(),
		bson.NewDocument(
			bson.EC.Int64("id", id),
		),
	)

	return err
}

// Count returns the number of the job records in the database.
func (dao *JobDAO) Count(db *mongo.Database) (int64, error) {
	col := db.Collection("job")

	return col.Count(
		context.Background(),
		bson.NewDocument(),
	)
}

// Query retrieves the job records with the specified offset and limit from the database.
func (dao *JobDAO) Query(db *mongo.Database, offset, limit int) ([]*models.Job, error) {
	jobList := []*models.Job{}
	col := db.Collection("job")
	ctx := context.Background()

	cursor, err := col.Find(
		ctx,
		bson.NewDocument(),
	)
	defer cursor.Close(ctx)
	elm := bson.NewDocument()

	for cursor.Next(ctx) {
		elm.Reset()

		if err := cursor.Decode(elm); err != nil {
			return jobList, err
		}
		job, err := models.NewJobFromDoc(elm)

		if err != nil {
			return jobList, err
		}

		jobList = append(jobList, job)
	}

	return jobList, err
}
