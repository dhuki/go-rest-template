package postgres

import (
	"bytes"
	"strconv"
	"time"

	"github.com/dhuki/go-rest-template/internal/adapter/repository"
	"github.com/dhuki/go-rest-template/internal/infra/configloader"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type postgresClient struct{}

func NewPostgreSQLClient() *postgresClient {
	return &postgresClient{}
}

func (p *postgresClient) NewPgRepository(conf *configloader.DatabaseConfig) (db repository.IRepository, err error) {
	dbMaster, err := p.OpenPostgresConnection(&conf.Master, &conf.DbConnectionInfo)
	if err != nil {
		return nil, err
	}
	dbSlave, err := p.OpenPostgresConnection(&conf.Slave, &conf.DbConnectionInfo)
	if err != nil {
		return nil, err
	}
	return repository.NewRepository(dbMaster, dbSlave), nil
}

func (p *postgresClient) OpenPostgresConnection(dbInfo *configloader.DBInfo, dbConnInfo *configloader.DbConnectionInfo) (*sqlx.DB, error) {
	var bufferStr bytes.Buffer
	bufferStr.WriteString(" host=" + dbInfo.Host)
	bufferStr.WriteString(" port=" + strconv.Itoa(dbInfo.Port))
	bufferStr.WriteString(" user=" + dbInfo.User)
	bufferStr.WriteString(" dbname=" + dbInfo.DBName)
	bufferStr.WriteString(" password=" + dbInfo.Password)
	bufferStr.WriteString(" sslmode=disable fallback_application_name=go-rest-example")
	connectionSource := bufferStr.String()

	db, err := sqlx.Connect("postgres", connectionSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(dbConnInfo.SetMaxIdleCons)
	db.SetMaxOpenConns(dbConnInfo.SetMaxOpenCons)
	db.SetConnMaxIdleTime(time.Duration(dbConnInfo.SetConMaxIdleTime) * time.Minute)
	db.SetConnMaxLifetime(time.Duration(dbConnInfo.SetConMaxLifetime) * time.Minute)
	return db, nil
}
