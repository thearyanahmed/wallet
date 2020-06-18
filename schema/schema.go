package schema

import (
	"github.com/thearyanahmed/wallet/database"
	"time"
)
// Future feature
const (
	UserActive = 1
	UserRestricted = 2
)

type User struct {
	ID uint `gorm:"primary_key" json:"id"`

	FirstName string `gorm:"type:varchar(12);column:first_name;not null" json:"first_name"`
	LastName string `gorm:"type:varchar(12);column:last_name;not null" json:"last_name"`
	Email string `gorm:"type:varchar(100);column:email;unique;not null" json:"email"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}

func (user *User) TableName() string {
	return "users"
}

const (
	RegularAccount = 1
	MerchantAccount = 2
	OrganizationAccount = 3
	AdminAccount = 4
)

type Account struct {
	UserID int32 `gorm:"type:integer;column:user_id;not null" json:"user_id"`
	RefID string `gorm:"type:varchar(30);column:ref_id;unique;not null" json:"ref_id"`
	Type int8 `gorm:"type:tinyint(1);defualt:1;not null" json:"type"` //TODO learn to use variables inside template strings
	OrgID int32 `gorm:"type:integer;column:org_id;not null" json:"org_id"`
	DefaultWalletCurrency string `gorm:"type:varchar(5);default:'USD';not null" json:"default_wallet_currency"` // TODO use enum

	ID uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`

	User User `gorm:"foreignkey:UserID"`
	Wallets []Wallet `gorm:"foreignkey:WalletID"`
	Organization Organization `gorm:"foreignkey:OrgID"`
}

type Wallet struct {
	UserID int32 `gorm:"type:integer;not null" json:"user_id"`
	AccountID int32 `gorm:"type:integer;not null" json:"account_id"`
	CurrencyCode string `gorm:"type:varchar(5);not null" json:"currency_code"`
	CurrencyID int8 `gorm:"type:tinyint(1);not null" json:"currency_id"`
	AvailableBalance int64 `gorm:"column:available_balance;default:0;not null" json:"available_balance"`
	TotalBalance int64 `gorm:"column:total_balance;default:0;not null" json:"total_balance"`

	ID uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`

	User User `gorm:"foreignkey:UserID"`
	Account Account `gorm:"foreignkey:AccountID"`
}

type Currency struct {
	ID        uint `gorm:"primary_key" json:"id"`
	Code string `gorm:"type:varchar(5);not null;unique" json:"code"` // USD
	Symbol string `gorm:"type:varchar(5);not null" json:"symbol"` // $

	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}


func (Currency) TableName() string {
	return "currencies"
}

type Organization struct {
	ID        uint `gorm:"primary_key" json:"id"`
	Name string `gorm:"type:varchar(50);not null" json:"name"` // USD
	UserID uint `gorm:"type:integer;not null" json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`

	User User `gorm:"foreignkey:user_id;association_foreignkey:id" json:"-"`
}

func Migrate() {
	db := database.DB()

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Account{})
	db.AutoMigrate(&Wallet{})
	db.AutoMigrate(&Currency{})
	db.AutoMigrate(&Organization{})
}

