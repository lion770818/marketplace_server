package servers

import (
	"fmt"
	"marketplace_server/config"
	"time"

	"marketplace_server/internal/bill"
	model_bill "marketplace_server/internal/bill/model"
	"marketplace_server/internal/user"
	model_user "marketplace_server/internal/user/model"

	"marketplace_server/internal/common/logs"
	"marketplace_server/internal/common/mysql"

	"marketplace_server/internal/common/redis"

	"github.com/jinzhu/gorm"

	//  mysql driver
	//_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

// 持久化管理物件
type RepositoriesManager struct {
	UserRepo user.UserRepo
	AuthRepo user.AuthInterface
	BillRepo bill.BillRepo
	db       *gorm.DB
}

// 建立持久化管理物件
func NewRepositories(cfg *config.SugaredConfig) *RepositoriesManager {

	// 持久化类型的 repo
	mysqlCfg := mysql.Config{
		LogMode:  cfg.Mysql.LogMode,
		Driver:   cfg.Mysql.Driver,
		Host:     cfg.Mysql.Host,
		Port:     cfg.Mysql.Port,
		Database: cfg.Mysql.Database,
		User:     cfg.Mysql.User,
		Password: cfg.Mysql.Password,
	}
	db := mysql.NewDB(mysqlCfg)

	userRepo := user.NewMysqlUserRepo(db)
	billRepo := bill.NewMysqlBillRepo(db)

	// 建立redis連線
	redisCfg := &redis.RedisParameter{
		Network:      "tcp",
		Address:      fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           0,
		DialTimeout:  time.Second * time.Duration(10),
		ReadTimeout:  time.Second * time.Duration(10),
		WriteTimeout: time.Second * time.Duration(10),
		PoolSize:     10,
	}
	redisClient, err := redis.NewRedis(redisCfg)
	if err != nil {
		logs.Errorf("newRedis error=%v", err)
	}

	// auth 策略
	var authRepo user.AuthInterface
	if cfg.Auth.Active == "redis" {
		logs.Debugf("使用redis當驗證緩存")
		authRepo = user.NewRedisAuthRepo(redisClient.GetClient(), cfg.AuthExpireTime)
	} else {
		logs.Debugf("使用jwt當驗證緩存")
		authRepo = user.NewJwtAuth(cfg.Auth.PrivateKey, cfg.AuthExpireTime)
	}

	return &RepositoriesManager{
		UserRepo: userRepo,
		AuthRepo: authRepo,
		BillRepo: billRepo,
		db:       db,
	}
}

// closes the  database connection
func (s *RepositoriesManager) Close() error {
	return s.db.Close()
}

func (s *RepositoriesManager) GetDB() *gorm.DB {
	return s.db
}

// This migrate all tables
func (s *RepositoriesManager) Automigrate() error {
	return s.db.AutoMigrate(&model_user.UserPO{}, &model_bill.BillPO{}).Error
}
