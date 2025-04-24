package oneGroup

import (
	"fmt"
	"go.uber.org/zap"
)

func InitZap() {
	zaps := zap.NewDevelopmentConfig()
	zaps.OutputPaths = []string{"./zap.log"}
	build, err := zaps.Build()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(build)
	fmt.Println("zap 初始化成功")
}
