package mysql

import (
	"time"

	"nasdaqvfs/config"
	"nasdaqvfs/pkg/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

// Return new MySQL db instance
func New(dn config.Database, c *config.MySQLConfig) (*gorm.DB, error) {
	logMode := logger.Silent
	if c.Debug {
		logMode = logger.Error
	}

	uri, err := c.GetURI(dn)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(mysql.Open(uri.MasterURI), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})
	if err != nil {
		return nil, err
	}

	replicas := make([]gorm.Dialector, 0, len(uri.SlaveURIs))
	for _, slaveURI := range uri.SlaveURIs {
		replicas = append(replicas, mysql.Open(slaveURI))
	}

	if db.Use(
		dbresolver.Register(
			dbresolver.Config{
				Replicas:          replicas,
				TraceResolverMode: true,
			},
		).SetConnMaxIdleTime(time.Duration(c.MaxIdleConns) * time.Second).
			SetMaxOpenConns(c.MaxOpenConns).
			SetConnMaxLifetime(time.Duration(c.ConnMaxLifeTime) * time.Second),
	); err != nil {
		return nil, err
	}

	return db, nil
}

// Paginate paginates results as pagination info
func Paginate(
	db *gorm.DB,
	value interface{},
	pq *utils.PaginationQuery,
) (*gorm.DB, error) {
	cols := make([]clause.OrderByColumn, len(pq.GetOrderBy()))
	for idx := range pq.GetOrderBy() {
		cols[idx] = clause.OrderByColumn{
			Column: clause.Column{Name: pq.GetOrderBy()[idx].Column},
			Desc:   pq.GetOrderBy()[idx].Order == utils.OrderByDesc,
		}
	}

	db = db.Model(value)

	if pq.IsPaginate() {
		db = db.Offset(pq.GetOffset()).Limit(pq.GetLimit())
	}

	db = db.Clauses(clause.OrderBy{Columns: cols})

	return db, nil
}

// CountTotal counts total items within table after being filtered (if any)
func CountTotal(
	db *gorm.DB,
	value interface{},
	pq *utils.PaginationQuery,
) (*gorm.DB, error) {
	var totalCount int64 = 0
	if err := db.Model(value).Count(&totalCount).Error; err != nil {
		return db, err
	}

	pq.SetTotalCount(totalCount)

	return db, nil
}
