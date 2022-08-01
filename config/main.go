package config


type LogConfig struct {
	// 日志文件
	Level string `json:"level"`
	Filename string `json:"filename"`
	// 日志的大小
	MaxSize int `json:"maxSize"`
	// 日志保留多少
	MaxBackup int `json:"maxBackup"`

	MaxAge int `json:"maxAge"`
}

var LogConf = &LogConfig{
	Level: "info",
	Filename: "logs/access.log",
	MaxBackup: 5,
	MaxAge: 7,
	MaxSize: 30,
}

var KubeConfigMap = map[string]string{
	"ack": "ack_config",
	"hza": "hza_config",
}
