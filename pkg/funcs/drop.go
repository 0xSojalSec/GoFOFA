package funcs

import (
	"github.com/lubyruffy/gofofa/pkg/pipeast"
	"github.com/lubyruffy/gofofa/pkg/piperunner"
)

func dropHook(fi *pipeast.FuncInfo) string {
	// 调用zq
	return `ZqQuery(GetRunner(), map[string]interface{}{
    "query": "drop ` + fi.Params[0].RawString() + `",
})`
}

func init() {
	piperunner.RegisterWorkflow("drop", dropHook, "", nil) // 删除字段
}
