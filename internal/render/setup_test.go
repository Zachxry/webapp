package render

import (
	"WebApp/internal/config"
	"WebApp/internal/models"
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	gob.Register(models.ConfirmInfo{})

	// change this when in production
	testApp.InProduction = false

	session = scs.New()               // returns a new session manager with the default options. It is safe for concurrent use.
	session.Lifetime = 24 * time.Hour // session set for 24 hours
	session.Cookie.Persist = true     // cookie is persistent
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session // stores the session in our app-wide configuration.

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}
