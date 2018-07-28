package game

func (game *Game) Minor(v ...interface{}) {
	if game.Logs != nil {
		game.Logs.Minor.Println(v...)
	}
}

func (game *Game) Major(v ...interface{}) {
	if game.Logs != nil {
		game.Logs.Major.Println(v...)
	}
}

func (game *Game) Debug(v ...interface{}) {
	if game.Logs != nil {
		game.Logs.Debug.Println(v...)
	}
}

func (game *Game) Info(v ...interface{}) {
	if game.Logs != nil {
		game.Logs.Info.Println(v...)
	}
}

func (game *Game) Warn(v ...interface{}) {
	if game.Logs != nil {
		game.Logs.Warn.Println(v...)
	}
}

func (game *Game) Error(v ...interface{}) {
	if game.Logs != nil {
		game.Logs.Error.Println(v...)
	}
}

func (game *Game) Fatal(v ...interface{}) {
	if game.Logs != nil {
		game.Logs.Fatal.Fatalln(v...)
	}
}
