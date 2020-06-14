package oauth

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/thearyanahmed/wallet/internal/reminder"
	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"log"
	"net/http"
)

var Server *server.Server
var clientStore *store.ClientStore

func Boot() {

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

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
	//clientId := uuid.New().String()[:8]
	//clientSecret := uuid.New().String()[:8]

	clientId := uuid.New().String()[:8]
	clientSecret := uuid.New().String()[:8]

	err := clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: "http://localhost:8000",
	})

	if err != nil {
		fmt.Println(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"CLIENT_ID": clientId, "CLIENT_SECRET": clientSecret})
}


