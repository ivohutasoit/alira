package util

import (
	"os"

	"github.com/gin-contrib/sessions/cookie"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

// Store keep our sessions
var Store *sessions.Store

func init() {
	Store := cookie.NewStore([]byte(os.Getenv("SECRET_KEY")))
	//Store = sessions.NewFilesystemStore(os.TempDir(), []byte(os.Getenv("SECRET_KEY")))
	//Store.MaxLength(math.MaxInt64)
	gothic.Store = Store
}
