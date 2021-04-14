package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	_, errInsertVideo := repo.Insert(video)
	require.Nil(t, errInsertVideo)

	job, err := domain.NewJob("outpuit_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	_, errInsertJob := repoJob.Insert(job)
	require.Nil(t, errInsertJob)

	j, errFindJob := repoJob.Find(job.ID)
	require.NotEmpty(t, j)
	require.Nil(t, errFindJob)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)
}

func TestJobRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	_, errInsertVideo := repo.Insert(video)
	require.Nil(t, errInsertVideo)

	job, err := domain.NewJob("outpuit_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.JobRepositoryDb{Db: db}
	_, errInsertJob := repoJob.Insert(job)
	require.Nil(t, errInsertJob)

	job.Status = "Complete"

	_, errUpdateJob := repoJob.Update(job)
	require.Nil(t, errUpdateJob)

	j, err := repoJob.Find(job.ID)
	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)
}
