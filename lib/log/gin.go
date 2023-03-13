package log

type GinLogger struct {
	*Logger
}

func (l *Logger) NewGinLogger() *GinLogger {
	return &GinLogger{l}
}

func (gl *GinLogger) Write(p []byte) (n int, err error) {
	gl.Debug("Gin", string(p))
	return len(p), nil
}
