package services

import (
	"Application/initializers"
	"Application/models"
)

func PingDatabase() (status int, resp models.Response) {
	postgresDB, postgresDBError := initializers.DB().DB()
	if postgresDBError != nil {
		resp = models.DataResponse(
			models.HealthCheckFailMsg,
			models.HealthCheckFailStatus,
			make([]interface{}, 0),
		)
		return
	}
	pingError := postgresDB.Ping()
	if pingError != nil {
		resp = models.DataResponse(
			models.HealthCheckFailMsg,
			models.HealthCheckFailStatus,
			make([]interface{}, 0),
		)
		return
	}
	resp = models.DataResponse(
		models.HealthCheckPassMsg,
		models.HealthCheckPassStatus,
		make([]interface{}, 0),
	)
	return resp.Status, resp
}
