package json

type MainConfig struct {
	WorkingDirectory string `json:"working-directory"`
}

type Config struct {
	Language string `json:"language"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Memory   int    `json:"memory"`
}

type SQLiteDatabaseConfig struct {
	Filename string `json:"filename"`
	Prefix 	 string `json:"prefix"`
}

type MariaDBDatabaseConfig struct {
	Host 	 string `json:"host"`
	Port 	 int 	`json:"port"`
	Database string	`json:"database"`
	Username string	`json:"username"`
	Password string	`json:"password"`
	Prefix 	 string `json:"prefix"`
}

type MongoDBDatabaseConfig struct {
	Host 	 string `json:"host"`
	Port 	 int 	`json:"port"`
	Database string	`json:"database"`
	Username string	`json:"username"`
	Password string	`json:"password"`
	Prefix 	 string `json:"prefix"`
}

type GroupConfig struct {
	StartPriority	 	 int	`json:"startPriority"`
	Name 			 	 string `json:"name"`
	TemplateName	 	 string `json:"templateName"`
	Static				 bool	`json:"static"`
	Maintenance			 bool	`json:"maintenance"`
	Permission			 string `json:"permission"`
	MinMemory 	 	 	 int 	`json:"minMemory"`
	MaxMemory 	 	 	 int 	`json:"maxMemory"`
	MinOnlineServices	 int	`json:"minOnlineServices"`
	MaxOnlineServices	 int	`json:"maxOnlineServices"`
	MaxPlayers 	 	 	 int 	`json:"maxPlayers"`
	NewServiceProcent	 int 	`json:"newServiceProcent"`
	ServiceVersion	 	 string `json:"serviceVersion"`
	Java		 	 	 string `json:"java"`
}
