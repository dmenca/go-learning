package service

import (
	"context"
	"http-server/api/db"
	"http-server/dao"
)

type DbService struct {
	DbQueryDao dao.DBQuery
}

func (s *DbService) DBGet(ctx context.Context, in *db.DBGetRequest) (*db.DBGetResponse, error) {
	if in.GetDbId() == "" {
		return nil, ErrDbIdNotFound
	}
	response := db.DBGetResponse{}
	dbInfo := db.DBInfo{}
	dbInfo.DbId = in.GetDbId()
	db, err := s.DbQueryDao.GetDBById(in.GetDbId())
	if err != nil {
		return nil, err
	}
	dbInfo.Name = db.Name
	response.DbInfo = &dbInfo
	return &response, nil
}
