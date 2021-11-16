package logger

func Bg() Logger {
	loggerInstance := getLogger()
	return &logger{
		info:   loggerInstance.Info,
		errorF: loggerInstance.Error,
		fatal:  loggerInstance.Fatal,
		with:   loggerInstance.With,
	}
}
