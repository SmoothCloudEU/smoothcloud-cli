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