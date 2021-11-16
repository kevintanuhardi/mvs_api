package logger

import (
	"testing"

	"go.uber.org/zap"
)

func ExampleBg() {
	Bg().Info("Info", zap.String("payload", "here"))
	Bg().Error("Error", zap.String("payload", "here"))
}
func TestBG(t *testing.T) {
	Bg().Info("Info", zap.String("payload", "here"))
	Bg().Error("Error", zap.String("payload", "here"))
}
