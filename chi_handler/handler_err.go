package chi_handler

import (
	"net/http"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, 400, "Something went wrong")
}
