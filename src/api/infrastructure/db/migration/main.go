package migration

import (
	"api/domain"
	"api/interfaces"
)

func Exec(sqlHandler interfaces.SQLHandler) {
	sqlHandler.Migrate(domain.Product{})
}
