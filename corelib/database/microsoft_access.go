package database

import (
	"fmt"

	_ "github.com/alexbrainman/odbc"
	"github.com/jmoiron/sqlx"
)

type MicrosoftAccessConfig struct {
	FilePath string // example: C:\\Users\\Boss\\Desktop\\play\\playgo\\Database1.accdb
}

func NewMicrosoftAccess(cfg *MicrosoftAccessConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("Driver={Microsoft Access Driver (*.mdb, *.accdb)};Dbq=%s;",
		cfg.FilePath,
	)
	db, err := sqlx.Open("odbc", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
