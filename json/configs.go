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
	Name 			 	 string `json:"name"`
	TemplateName	 	 string `json:"templateName"`
	StartPriority	 	 int	`json:"startPriority"`
	Static				 bool	`json:"static"`
	Maintenance			 bool	`json:"maintenance"`
	Permission			 any	`json:"permission"`
	MinMemory 	 	 	 int 	`json:"minMemory"`
	MaxMemory 	 	 	 int 	`json:"maxMemory"`
	MinOnlineServices	 int	`json:"minOnlineServices"`
	MaxOnlineServices	 int	`json:"maxOnlineServices"`
	MaxPlayers 	 	 	 int 	`json:"maxPlayers"`
	NewServiceProcent	 int 	`json:"newServiceProcent"`
	ServiceVersion	 	 string `json:"serviceVersion"`
	Java		 	 	 string `json:"java"`
}

type ServiceVersion struct {
	PROXY  map[string]VersionConfig  	`json:"PROXY"`
	SERVER map[string]VersionConfig `json:"SERVER"`
}

type VersionConfig struct {
	Versions map[string]interface{} `json:"versions"`
}