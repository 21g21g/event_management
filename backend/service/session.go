package service

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("your-session-secret"))

func GetSession(r *http.Request) (*sessions.Session, error) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return nil, err
	}
	return session, nil
}

func SetSession(w http.ResponseWriter, r *http.Request, userID string) error {
	session, _ := store.Get(r, "session-name")
	session.Values["user_id"] = userID
	return sessions.Save(r, w)
}
