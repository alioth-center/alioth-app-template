package global

type ProjectConfig struct {
	HttpEngineConfig HttpEngineConfig `yaml:"http_engine"`
	LogConfig        LogConfig        `yaml:"logger"`
	DatabaseConfig   DatabaseConfig   `yaml:"database"`
}

// HttpEngineConfig is the configuration for the alioth infrastructure http engine
type HttpEngineConfig struct {
	ServeURL  string `yaml:"serve_url"`
	ServeAddr string `yaml:"serve_addr"`
}

// LogConfig is the configuration for the alioth infrastructure logger
type LogConfig struct {
	LogToFile     bool   `yaml:"log_to_file"`
	LogSplit      bool   `yaml:"log_split"`
	LogDirectory  string `yaml:"log_directory"`
	LogLevel      string `yaml:"log_level"`
	LogFileSuffix string `yaml:"log_file_suffix"`
}

// DatabaseConfig is the configuration for the alioth infrastructure database drivers
type DatabaseConfig struct {
	Driver     string `yaml:"driver"`
	Host       string `yaml:"host"`
	Port       int    `yaml:"port"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	Database   string `yaml:"database"`
	SSL        bool   `yaml:"ssl"`
	Location   string `yaml:"location"`
	SyncModels bool   `yaml:"sync_models"`
}
