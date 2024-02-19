package own

import "go.uber.org/zap"

func InitZap() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic("init zap failed: " + err.Error())
	}
	zap.ReplaceGlobals(logger)
}
