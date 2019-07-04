package app

func (a *App) GetMyLittleService() func() {
	if a.myLittleService == nil {
		a.myLittleService = myLittleService
	}
	return a.myLittleService
}

// little service can be embedded in this file
func myLittleService() {
	print("print from little service")
}
