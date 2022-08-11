package strategy

import "testing"

func TestLogger(t *testing.T) {
	logManager := NewLogManager(&FileLogger{})
	logManager.Info()
	logManager.Erorr()

	logManager.Logger = &DbLogger{}
	logManager.Info()
	logManager.Erorr()

}
