package mapper

import (
	"encoding/json"
	"net/http"

	"github.com/roronoadiogo/todo-go-rest/config"
	"go.uber.org/zap/zapcore"
)

var logger = config.InitializeLogger()

func ParseToJson(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)

	if err != nil {
		logger.Error("Failed in the parse of data: ",
			zapcore.Field{Key: "error", Interface: data})
	}
}
