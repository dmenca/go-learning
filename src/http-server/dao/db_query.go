package dao

import (
	"database/sql"
	log "github.com/sirupsen/logrus"
	"http-server/dao/entity"
)

type DBQuery struct {
	Db *sql.DB
}

func (m *DBQuery) GetDBById(id string) (*entity.DB, error) {
	var db entity.DB
	err := m.Db.QueryRow("select db_id,name from dbs where db_id=?", id).Scan(&db.DBID, &db.Name)
	if err != nil {
		log.Warn("get db by id error", err)
	}
	return &db, err

}
