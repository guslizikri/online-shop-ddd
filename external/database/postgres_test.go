package database

import (
	"online-shop-ddd/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}

}

func TestConnectionPostgre(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		db, err := ConnectPostgres(config.Cfg.DB)
		require.Nil(t, err)
		require.NotNil(t, db)
	})

}
