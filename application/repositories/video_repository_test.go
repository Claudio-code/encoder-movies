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

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	_, errInsertVideo := repo.Insert(video)
	require.Nil(t, errInsertVideo)

	v, err := repo.Find(video.ID)
	require.Nil(t, err)
	require.NotEmpty(t, v.ID)
	require.Equal(t, v.FilePath, video.FilePath)
}

func TestVideoRepositoryDbUpdate(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.NewV4().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	_, errInsertVideo := repo.Insert(video)
	require.Nil(t, errInsertVideo)

	video.FilePath = "pathFuny"

	_, errUpdateVideo := repo.Update(video)
	require.Nil(t, errUpdateVideo)

	v, err := repo.Find(video.ID)
	require.Nil(t, err)
	require.NotEmpty(t, v.ID)
	require.Equal(t, v.FilePath, video.FilePath)
}
