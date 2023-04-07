package mapper

import (
	"encoding/json"
	"net/http"

	"github.com/roronoadiogo/todo-go-rest/config"
)

var logger = config.InitializeLogger()

func ParseToJson(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	resp, err := json.Marshal(w)
	if err != nil {
		logger.Sugar().Debugf("Error in the parse to Json: %s", err)
	}

	w.Write([]byte(resp))
}
