package global

import (
	"github.com/alioth-center/infrastructure/config"
	"github.com/alioth-center/infrastructure/database"
	"github.com/alioth-center/infrastructure/database/mysql"
	"github.com/alioth-center/infrastructure/database/postgres"
	"github.com/alioth-center/infrastructure/database/sqlite"
	"github.com/alioth-center/infrastructure/logger"
	"github.com/alioth-center/infrastructure/utils/values"
	"path/filepath"
	"time"
)

var syncModels = []any{
	// todo: add models here
}

func init() {
	initializeConfig()
	initializeLogger()
	initializeDatabase()
}

func initializeConfig() {
	readErr := config.LoadConfig(&Config, "./config/config.yaml")
	if readErr != nil {
		panic(readErr)
	}
}

func initializeLogger() {
	if Config.LogConfig.LogFileSuffix == "" {
		Config.LogConfig.LogFileSuffix = "_akasha_whisper_logs_stdout.jsonl"
	}
	switch {
	case !Config.LogConfig.LogToFile:
		Logger = logger.NewCustomLoggerWithOpts(
			logger.WithLevelOpts(logger.Level(Config.LogConfig.LogLevel)),
		)
	case !Config.LogConfig.LogSplit:
		Logger = logger.NewCustomLoggerWithOpts(
			logger.WithLevelOpts(logger.Level(Config.LogConfig.LogLevel)),
			logger.WithFileWriterOpts(filepath.Join(Config.LogConfig.LogDirectory, Config.LogConfig.LogFileSuffix)),
		)
	default:
		Logger = logger.NewCustomLoggerWithOpts(
			logger.WithLevelOpts(logger.Level(Config.LogConfig.LogLevel)),
			logger.WithCustomWriterOpts(
				logger.NewTimeBasedRotationFileWriter(Config.LogConfig.LogDirectory, func(time time.Time) (filename string) {
					return values.BuildStrings(time.Format("2006-01-02"), Config.LogConfig.LogFileSuffix)
				}),
			),
		)
	}
}

func initializeDatabase() {
	var db database.DatabaseV2
	switch Config.DatabaseConfig.Driver {
	case "":
	case postgres.DriverName:
		pgCfg := postgres.Config{
			Host:      Config.DatabaseConfig.Host,
			Port:      Config.DatabaseConfig.Port,
			Username:  Config.DatabaseConfig.Username,
			Password:  Config.DatabaseConfig.Password,
			Database:  Config.DatabaseConfig.Database,
			Location:  Config.DatabaseConfig.Location,
			EnableSSL: Config.DatabaseConfig.SSL,
			Debug:     false,
		}

		models := syncModels
		if !Config.DatabaseConfig.SyncModels {
			models = []any{}
		}
		pgDB, initErr := postgres.NewWithLogger(pgCfg, Logger, models...)
		if initErr != nil {
			panic(initErr)
		}

		db = pgDB
	case mysql.DriverName:
		mysqlCfg := mysql.Config{
			Server:    Config.DatabaseConfig.Host,
			Port:      Config.DatabaseConfig.Port,
			Username:  Config.DatabaseConfig.Username,
			Password:  Config.DatabaseConfig.Password,
			Database:  Config.DatabaseConfig.Database,
			Location:  Config.DatabaseConfig.Location,
			ParseTime: Config.DatabaseConfig.Location != "",
		}

		models := syncModels
		if !Config.DatabaseConfig.SyncModels {
			models = []any{}
		}
		mysqlDB, initErr := mysql.NewWithLogger(mysqlCfg, Logger, models...)
		if initErr != nil {
			panic(initErr)
		}

		db = mysqlDB
	case sqlite.DriverName:
		sqliteCfg := sqlite.Config{
			Database: "./data/akasha_whisper.db",
		}

		models := syncModels
		if !Config.DatabaseConfig.SyncModels {
			models = []any{}
		}
		sqliteDB, initErr := sqlite.NewWithLogger(sqliteCfg, Logger, models...)
		if initErr != nil {
			panic(initErr)
		}

		db = sqliteDB
	default:
		panic("unsupported database driver")
	}

	DatabaseInstance = db
}
