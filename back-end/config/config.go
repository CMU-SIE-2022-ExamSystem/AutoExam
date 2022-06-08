package config

type ServerConfig struct {
	Name        string        `mapstructure:"name"`
	Port        int           `mapstructure:"port"`
	Mysqlinfo   MysqlConfig   `mapstructure:"mysql"`
	LogsAddress string        `mapstructure:"logsAddress"`
	Logger      string        `mapstructure:"logger"`
	Autolabinfo AutolabConfig `mapstructure:"autolab"`
}

type AutolabConfig struct {
	Ip            string `mapstructure:"ip"`
	Client_id     string `mapstructure:"client_id"`
	Client_secret string `mapstructure:"client_secret"`
	Redirect_uri  string `mapstructure:"redirect_uri"`
	Scope         string `mapstructure:"scope"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}
