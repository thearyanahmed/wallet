package oauth

import (
	"fmt"
	"github.com/thearyanahmed/wallet/internal/reminder"
	"github.com/thearyanahmed/wallet/internal/res"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	Server *server.Server
	clientStore *store.ClientStore
)

const (
	accessTokenExpiresIn = time.Hour * 24 * 30
	refreshTokenExpires = time.Hour * 24 * 45
)

func Boot() {

	manager := manage.NewDefaultManager()

	manager.SetAuthorizeCodeTokenCfg(&manage.Config{
		AccessTokenExp:    accessTokenExpiresIn,
		RefreshTokenExp:   refreshTokenExpires,
		IsGenerateRefresh: false,
	})

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore = store.NewClientStore()

	manager.MapClientStorage(clientStore)

	Server := server.NewDefaultServer(manager)
	Server.SetAllowGetAccessRequest(true)
	Server.SetClientInfoHandler(server.ClientFormHandler)

	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	Server.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		reminder.Remind("[+] Setup application logger. Setting internal error handler.")
		log.Println("Internal Error:", err.Error())
		return
	})

	Server.SetResponseErrorHandler(func(re *errors.Response) {
		reminder.Remind("[+] Setup application logger. Setting response error handle.")
		log.Println("Response Error:", re.Error.Error())
	})

	registerOAuthRoutes()
}

func registerOAuthRoutes() {
	http.HandleFunc("/token", serveToken)
	http.HandleFunc("/demo/credentails",demoCredentails)
}

func serveToken(w http.ResponseWriter, r *http.Request) {
	Server.HandleTokenRequest(w, r)
}

func Auth(f http.HandlerFunc, srv *server.Server) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := srv.ValidationBearerToken(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		f.ServeHTTP(w, r)
	})
}

func demoCredentails (w http.ResponseWriter, r *http.Request) {

	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	clientDomain := os.Getenv("CLIENT_DOMAIN")

	err := clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: clientDomain,
	})

	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	data := struct {
		ClientId string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{
		ClientId: clientId,
		ClientSecret: clientSecret,
	}

	res.Send(w,"Successfully done.",data,200)
}


