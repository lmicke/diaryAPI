type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Init(user, password, database string) {
	//ToDo: Do the init
}

func (a *App) Run(add string) {
	//ToDo: Do the running
}
