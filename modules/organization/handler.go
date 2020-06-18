package organization

import (
	request "github.com/thearyanahmed/wallet/internal/req"
	"github.com/thearyanahmed/wallet/internal/res"
	account "github.com/thearyanahmed/wallet/modules/account"
	"github.com/thearyanahmed/wallet/modules/currency"
	"github.com/thearyanahmed/wallet/modules/user"
	"github.com/thearyanahmed/wallet/modules/wallet"
	"net/http"
	"strconv"
)

type handler struct {
	Service
}

const (
	orgCreatedSuccessfully = "Organization created successfully."
	userNotFound = "User not found."
	currencyNotFound = "Currency not found."
)


func NewHandler() *handler {
	return &handler{Service{organizationRepository{}}}
}

func (handler *handler) createNewOrganization(w http.ResponseWriter,r *http.Request) {

	req := request.Request{}

	if valid := req.Validate(r,w,createNewOrganizationRequest); valid == false {
		return
	}

	validated := req.ValidatedFormData(r,[]string{"user_id","name","currency_code","org_id"})

	// check if currency exists
	currencySvc := currency.Service{}

	currency, errs := currencySvc.FindCurrencyByCode(validated["currency_code"])

	if len(errs) > 0 {
		res.SendError(w,currencyNotFound,nil,422)
		return
	}
	// check if user exists

	userID, _ := strconv.Atoi(validated["user_id"])

	userSvc := user.Service{}

	// check if user exists
	_, errs = userSvc.FindUserById(uint(userID))

	if len(errs) > 0 {
		res.SendError(w,userNotFound,nil,422)
		return
	}

	orgSvc := Service{}
	accountSvc := account.Service{}
	walletSvc := wallet.Service{}

	org ,err := orgSvc.CreateOrganization(uint(userID),validated["name"])

	if err != nil {
		res.SendError(w,res.UnprocessableEntity,errs,422)
		return
	}
	// begin transaction

	orgAccount, err := accountSvc.CreateNewAccount(uint(userID),org.ID,validated["currency_code"])

	if err != nil {
		res.SendError(w,res.UnprocessableEntity,errs,422)
		return
	}

	orgWallet, err := walletSvc.CreateNewWallet(uint(userID),org.ID,currency.ID,validated["currency_code"])

	if err != nil {
		res.SendError(w,res.UnprocessableEntity,errs,422)
		return
	}

	// end transaction

	response := createdResponse{
		Organization: orgResponse{
			ID:        org.ID,
			Name:      org.Name,
			CreatedAt: org.CreatedAt,
		},
		Account:   accountResponse{
			RefID:                   orgAccount.RefID,
			DefaultCurrencyForWallet: orgAccount.DefaultWalletCurrency,
		},
		Wallet:    walletResponse{
			AccountReference: orgAccount.RefID,
			CurrencyCode:     orgWallet.CurrencyCode,
			AvailableBalance: orgWallet.AvailableBalance,
			TotalBalance:     orgWallet.TotalBalance,
		},
	}

	res.Send(w,orgCreatedSuccessfully,response,200)

}