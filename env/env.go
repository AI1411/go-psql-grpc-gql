package env

type Env struct {
	Hostname string // GOPG_HOST
	Port     string // GOPG_PORT
	User     string // GOPG_USER
	Password string // GOPG_PASSWORD
	Dbname   string // GOPG_DBNAME
}
