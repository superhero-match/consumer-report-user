/*
  Copyright (C) 2019 - 2020 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package db

import (
	"database/sql"
	"fmt"

	"github.com/superhero-match/consumer-report-user/internal/config"

	_ "github.com/go-sql-driver/mysql" // MySQL driver.
)

// DB holds the database connection.
type DB struct {
	DB              *sql.DB
	stmtStoreReport *sql.Stmt
}

// NewDB returns database.
func NewDB(cfg *config.Config) (dbs *DB, err error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			cfg.DB.User,
			cfg.DB.Password,
			cfg.DB.Host,
			cfg.DB.Port,
			cfg.DB.Name,
		),
	)
	if err != nil {
		return nil, err
	}

	stmtReport, err := db.Prepare(`call insert_new_report(?,?,?,?)`)
	if err != nil {
		return nil, err
	}

	return &DB{
		DB:              db,
		stmtStoreReport: stmtReport,
	}, nil
}
