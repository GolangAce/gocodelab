package handlers

import (
	"gocodelab/gopherfaceauth/common/authenticate"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	authenticate.ExpireUserSession(w, r)
	authenticate.ExpireSecureCookie(w, r)
}
