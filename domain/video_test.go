package domain_test

import (
	"encoder/domain"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestValidateIfVideoEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestIdIsNotAUuid(t *testing.T) {
	video := domain.NewVideo()

	video.ID = "asdq"
	video.ResourceID = "dwq"
	video.FilePath = "/path"
	video.CreatedAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}
