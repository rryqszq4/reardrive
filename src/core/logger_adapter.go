package core

type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Warning(v ...interface{})
	Error(v ...interface{})

	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errorf(format string, v ...interface{})

	Close()
	//Reload() error
}

type LoggerAdapter struct {
	l *logger_t
}

func NewLoggerAdapter(file string, level int) Logger {

	l := NewLogger(file, level)

	return &LoggerAdapter{l}
}

func (l *LoggerAdapter) Debug(v ...interface{}) {
	l.l.Debug(v...)
}

func (l *LoggerAdapter) Info(v ...interface{}) {
	l.l.Info(v...)
}

func (l *LoggerAdapter) Warning(v ...interface{}) {
	l.l.Warn(v...)
}

func (l *LoggerAdapter) Error(v ...interface{}) {
	l.l.Error(v...)
}

func (l *LoggerAdapter) Debugf(format string, v ...interface{}) {
	l.l.Debugf(format, v...)
}

func (l *LoggerAdapter) Infof(format string, v ...interface{}) {
	l.l.Infof(format, v...)
}

func (l *LoggerAdapter) Warningf(format string, v ...interface{}) {
	l.l.Warnf(format, v...)
}

func (l *LoggerAdapter) Errorf(format string, v ...interface{}) {
	l.l.Errorf(format, v...)
}

func (l *LoggerAdapter) Close(){
	l.l.Close()
}

/*func (l * LoggerAdapter) Reload() error{
	return l.l.Reload()
}*/