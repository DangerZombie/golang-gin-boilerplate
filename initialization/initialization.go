package initialization

import (
	"github.com/DangerZombie/golang-gin-boilerplate/helper/auth"
	"github.com/DangerZombie/golang-gin-boilerplate/helper/database"
	transport "github.com/DangerZombie/golang-gin-boilerplate/http"
	"github.com/DangerZombie/golang-gin-boilerplate/model/entity"
	"github.com/DangerZombie/golang-gin-boilerplate/repository"
	"github.com/DangerZombie/golang-gin-boilerplate/repository/repository_teacher"
	"github.com/DangerZombie/golang-gin-boilerplate/repository/repository_user"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_teacher"
	"github.com/DangerZombie/golang-gin-boilerplate/service/service_user"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gorm.io/gorm"
)

func DbInit() (*gorm.DB, error) {
	// Init DB connection
	driver := viper.GetString("database.driver")
	dbname := viper.GetString("database.dbname")
	host := viper.GetString("database.host")
	user := viper.GetString("database.username")
	password := viper.GetString("database.password")
	port := viper.GetInt("database.port")

	db, err := database.NewDBConnection(driver, dbname, host, user, password, port)
	if err != nil {
		return nil, err
	}

	_ = db.AutoMigrate(&entity.Role{})
	_ = db.AutoMigrate(&entity.JobTitle{})
	_ = db.AutoMigrate(&entity.Teacher{})
	_ = db.AutoMigrate(&entity.User{})

	return db, nil
}

func ServerInit(log *zap.Logger, db *gorm.DB) {
	// base repository
	baseRepository := repository.NewBaseRepository(db)
	teacherRepository := repository_teacher.NewTeacherRepository(baseRepository)
	userRepository := repository_user.NewUserRepository(baseRepository)

	// auth helper
	authHelper := auth.NewAuthHelper(
		baseRepository,
		userRepository,
	)

	// service
	teacherSvc := service_teacher.NewTeacherService(
		log,
		baseRepository,
		teacherRepository,
	)

	userSvc := service_user.NewUserService(
		log,
		authHelper,
		baseRepository,
		userRepository,
	)

	r := gin.Default()

	// group endpoint
	apiGroupTeacher := r.Group("/api/v1/teacher")
	apiGroupUser := r.Group("/api/v1/user")

	// transport
	transportHandler := transport.NewHttp(
		authHelper,
	)

	transportHandler.SwaggerHttpHandler(r)
	transportHandler.TeacherHandler(apiGroupTeacher, teacherSvc)
	transportHandler.UserHandler(apiGroupUser, userSvc)

	r.Run()
}

func NewZapLogger(filename string) (*zap.Logger, error) {
	config := zap.NewProductionConfig()

	// if filename not empty
	if filename != "" {
		config.OutputPaths = append(config.OutputPaths, filename)
		config.ErrorOutputPaths = append(config.ErrorOutputPaths, filename)
	}

	config.Encoding = "json"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.EncodeDuration = zapcore.MillisDurationEncoder
	logger, err := config.Build()
	defer logger.Sync()

	return logger, err
}
