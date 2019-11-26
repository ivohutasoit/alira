package util

import (
	"math"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth/gothic"
)

// Store keep our sessions
var Store *sessions.FilesystemStore

func init() {
	Store = sessions.NewFilesystemStore(os.TempDir(), []byte(os.Getenv("SECRET_KEY")))
	Store.MaxLength(math.MaxInt64)
	gothic.Store = Store
}
