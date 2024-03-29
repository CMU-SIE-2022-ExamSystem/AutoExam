package config

type ServerConfig struct {
	Name         string        `mapstructure:"name"`
	Ip           string        `mapstructure:"ip"`
	Port         int           `mapstructure:"port"`
	Mysqlinfo    MysqlConfig   `mapstructure:"mysql"`
	LogsAddress  string        `mapstructure:"logsAddress"`
	Logger       string        `mapstructure:"logger"`
	Autolabinfo  AutolabConfig `mapstructure:"autolab"`
	JWTKey       JWTconfig     `mapstructure:"jwt"`
	MongoDB      MongoDBConfig `mapstructure:"mongodb"`
	Redis        RedisConfig   `mapstructure:"redis"`
	Basic_Grader []string
}

type AutolabConfig struct {
	Protocol      string `mapstructure:"protocol"`
	Skip_Secure   bool   `mapstructure:"skip_secure"`
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

type JWTconfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}

type MongoDBConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
