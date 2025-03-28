package global

import (
	"gin-vue-admin/utils/timer"
	ut "github.com/go-playground/universal-translator"

	"go.uber.org/zap"

	"gin-vue-admin/config"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	//GVA_LOG    *oplogging.Logger
	GVA_LOG   *zap.Logger
	GVA_Timer timer.Timer = timer.NewTimerTask()
	GVA_Trans ut.Translator
)
