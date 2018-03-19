package model

type Config struct {
	MysqlDatabase struct {
		Host     string `json:"host"`
		Password string `json:"password"`
		Username string `json:"username"`
		DBName   string `json:"dbname"`
	} `json:"mysqldatabase"`
	Host string `json:"host"`
	Port string `json:"port"`
	LogFile string `json:"logfile"`
	LogLevel LogLevel `json:"loglevel"`
}
