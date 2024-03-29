package model

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

var (
	Error_UserIDIsEmpty = errors.New("user id is empty")
	Error_ConvertFailed = errors.New("convert failed")
)

// po (presentation object) 持久化對象
// 用戶db持久化對象
type UserPO struct {
	UserID    int64           `gorm:"primary_key;auto_increment;comment:'流水號 主鍵'" json:"user_id"`
	Username  string          `gorm:"size:100;not null; comment:'使用者名稱'" json:"user_name"`
	Password  string          `gorm:"size:100;not null; comment:'使用者密碼'" json:"password"`
	Currency  string          `gorm:"size:32;not null; comment:'幣種'" json:"currency"`
	Amount    decimal.Decimal `gorm:"type:decimal(20,2); comment:'金額'" json:"amount"`
	CreatedAt time.Time       `gorm:"autoCreateTime;comment:'創建時間'" json:"created_at"`
	UpdateAt  time.Time       `gorm:"autoUpdateTime;comment:'更新時間'" json:"update_at"`
}

func (UserPO) TableName() string {
	return "user"
}

// 通常數據庫儲存的欄位會比較多, 輸出給前端的數據會比較少, 在此轉換
// ToDomain converts a UserRepo to a domain.User
func (u *UserPO) ToDomain() (*User, error) {

	if u.UserID == 0 {
		return nil, Error_UserIDIsEmpty
	}

	user := &User{
		UserID:    u.UserID,
		Username:  u.Username,
		Password:  u.Password,
		Currency:  u.Currency,
		Amount:    u.Amount,
		CreatedAt: u.CreatedAt,
		UpdateAt:  u.UpdateAt,
	}

	return user, nil
}
