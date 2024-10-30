package global

import (
	"github.com/alioth-center/infrastructure/database"
	"github.com/alioth-center/infrastructure/logger"
)

var (
	Config           ProjectConfig
	DatabaseInstance database.DatabaseV2
	Logger           logger.Logger
)
