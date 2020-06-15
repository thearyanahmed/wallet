package schema

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FirstName string `gorm:"type:varchar(60);column:first_name;not null" json:"first_name"`
	LastName string `gorm:"type:varchar(60);column:last_name;not null" json:"first_name"`
	Email string `gorm:"type:varchar(100);column:email;unique;not null" json:"email"`
}

func (User) TableName() string {
	return "users"
}

type Account struct {
	gorm.Model
	UserID int32 `gorm:"type:integer;column:user_id,not null" json:"user_id"`
	RefID string `gorm:"type:varchar(30);unique;column:ref_id,not null" json:"ref_id"`
	Type int8 `gorm:"type:tinyint(1);defualt:1,not null" json:"type"` //TODO learn to use variables inside template strings
	OrgID int32 `gorm:"type:integer;column:org_id,not null" json:"org_id"`
	DefaultWalletCurrency string `gorm:"type:varchar(5);default:'USD',not null" json:"default_wallet_currency"` // TODO use enum

	User User `gorm:"foreignkey:UserID"`
	Wallets []UserWallet `gorm:"foreignkey:WalletID"`
}

func (Account) TableName() string {
	return "accounts"
}

type UserWallet struct {
	gorm.Model

	UserID int32 `gorm:"type:integer;not null" json:"user_id"`
	AccountID int32 `gorm:"type:integer;not null" json:"account_id"`
	CurrencyCode string `gorm:"type:varchar(5);not null" json:"currency_code"`
	CurrencyID int8 `gorm:"type:tinyint(1)";not null" json:"currency_id"`
	AvailableBalance int64 `gorm:"column:available_balance";not null;default:0" json:"available_balance"`
	TotalBalance int64 `gorm:"column:total_balance";not null;default:0" json:"total_balance"`

	User User `gorm:"foreignkey:UserID"`
	Account Account `gorm:"foreignkey:AccountID"`
}

func (UserWallet) TableName() string {
	return "user_wallets"
}
