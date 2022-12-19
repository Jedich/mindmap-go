package repository

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"mindmap-go/utils"
)

func parseMySQLError(err error) error {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
		return utils.DuplicateEntryError
	}
	return err
}
