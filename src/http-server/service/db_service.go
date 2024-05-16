package service

import (
	"context"
	"http-server/api/db"
)

type DbService struct {
}

func (s *DbService) DBGet(ctx context.Context, in *db.DBGetRequest) (*db.DBGetResponse, error) {
	if in.GetDbId() == "" {
		return nil, ErrDbIdNotFound
	}
	response := db.DBGetResponse{}
	dbInfo := db.DBInfo{}
	dbInfo.DbId = in.GetDbId()
	dbInfo.Name = in.GetDbId() + "-Name"
	response.DbInfo = &dbInfo
	return &response, nil
}
