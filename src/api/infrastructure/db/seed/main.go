package seed

import (
	"api/domain"
	"api/interfaces"
)

func Exec(sqlHandler interfaces.SQLHandler) {
	sqlHandler.Create(&domain.Product{Code: "aaa", Price: 100})
	sqlHandler.Create(&domain.Product{Code: "bbb", Price: 200})
}
