package repository

import (
	"backend/config"
	"context"
	"os"
	"sync"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

// ADD TEST CONTAINERS

var cfg *config.Config
var conn *sqlx.DB

func setupConfig() (*config.Config, *sqlx.DB) {
	var once sync.Once
	once.Do(func() {
		cfg = config.New()
		conn = sqlx.MustConnect(cfg.SqlDriver, cfg.SqlConnString)
	})

	return cfg, conn
}

// test connections and pings to the repo layer itself
func TestMain(m *testing.M) {
	// WARN: Call once, use everywhere since resource is locked.
	cfg, conn = setupConfig()
	conn := sqlx.MustConnect(cfg.SqlDriver, cfg.SqlConnString)

	if err := conn.PingContext(context.TODO()); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func TestDBTXCreation(t *testing.T) {
	dbtx := New(conn)
	assert.NotNil(t, dbtx)
}

// Test ops bellow
