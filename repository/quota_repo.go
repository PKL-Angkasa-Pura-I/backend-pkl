package repository

/* import (
	"time"

	"github.com/PKL-Angkasa-Pura-I/backend-pkl/model"
	"gorm.io/gorm/clause"
)

func (r *repositoryMysqlLayer) GetSumQuotaOnRange(start time.Time) int {
	var sum int
	end := start.AddDate(0, 1, 0)
	r.DB.Table("submissions").Select("sum(total_trainee)").Where("status = diterima AND").Row().Scan(&sum)

	return sum
} */
