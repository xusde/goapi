package middleware

import (
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/xusde/goapi/api"
	"github.com/xusde/goapi/internal/tools"
)

var UnauthorizedError = errors.New("invalid username or token")
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")
		var err error
		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}
		// check again db
		var db *tools.DatabaseInterface
		db, err = tools.NewDataBase()
		if err != nil {
			api.InternalErrorHandler(w)
			return 
		}

		var loginDetails *tools.loginDetails
		loginDetails = (*db).GetLoginDetails(username)
		if (loginDetails == nil || (token != (*loginDetails).AuthToken)) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w. UnauthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}