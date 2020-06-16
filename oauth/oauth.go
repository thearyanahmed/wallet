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
	accessTokenExpiresIn = time.Hour * 24 * 23
	refreshTokenExpires = time.Hour * 24 * 45
)

func init() {
	reminder.Remind(fmt.Sprintf("[+] \nThe appliaction is set for one and one application only. It does not hold any state for currently logged in user other than one and only one.Validating a route with oauth.Auth will not authorize about the requeted resource."))
}

func Boot() {

	manager := manage.NewDefaultManager()

	manager.SetAuthorizeCodeExp(time.Minute * 10)

	cfg := &manage.Config{
		// access token expiration time
		AccessTokenExp: time.Hour * 2,
		// refresh token expiration time
		RefreshTokenExp: time.Hour * 24 * 3,
		// whether to generate the refreshing token
		IsGenerateRefresh: true,
	}
	manager.SetAuthorizeCodeTokenCfg(cfg)

	manager.SetClientTokenCfg(&manage.Config{
		AccessTokenExp:    time.Hour * 100,
		RefreshTokenExp:   time.Hour * 120,
		IsGenerateRefresh: true,
	})

	manager.SetRefreshTokenCfg(&manage.RefreshingConfig{
		AccessTokenExp:     accessTokenExpiresIn,
		RefreshTokenExp:    refreshTokenExpires,
		IsGenerateRefresh:  true,
		IsResetRefreshTime: true,
		IsRemoveAccess:     true,
		IsRemoveRefreshing: false,
	})

	// token memory store
	manager.MustTokenStorage(store.NewMemoryTokenStore())

	// client memory store
	clientStore = store.NewClientStore()

	manager.MapClientStorage(clientStore)

	Server = server.NewServer(server.NewConfig(),manager)
	Server.SetAllowGetAccessRequest(true)
	Server.SetClientInfoHandler(server.ClientFormHandler)

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

// At the moment, this auth is set for one client and one client only.
func Auth(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := Server.ValidationBearerToken(r)
		if err != nil {
			res.SendError(w,"Unauthorized.",nil,http.StatusUnauthorized)
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


