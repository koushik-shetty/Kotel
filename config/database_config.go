package config

type dbConfig struct {
	dirver   string
	hostname string
	dbName   string
	schema   string
	username string
	password string
}

func (dbConf *dbConfig) Driver() string {
	return dbConf.driver
}

func (dbconf *dbConfig) HostName() string {
	return dbconf.hostname
}

func (dbconf *dbConfig) DBName() string {
	return dbconf.dbName
}

func (dbconf *dbConfig) Schema() string {
	return dbconf.schema
}

func (dbconf *dbConfig) UserName() string {
	return dbconf.username
}

func (dbconf *dbConfig) Password() string {
	return dbconf.password
}

func DefaultDBConfig() *dbConfig {
	return &dbConfig{
		driver:   "postgres",
		hostname: "localhost",
		dbname:   "koteldb",
		schema:   "wall",
		username: "pilgrim",
		password: "western_wall",
	}
}

func NewDBConfig(host, dbname, schema, username, password string) *dbConfig {
	return &dbConfig{
		driver:   "postgres",
		hostname: host,
		dbname:   dbaname,
		schema:   schema,
		username: username,
		password: password,
	}
}
