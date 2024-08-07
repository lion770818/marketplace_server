package model

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"
)

var (
	Error_UserIDIsEmpty        = errors.New("user_id is empty")
	Error_TransactionIDIsEmpty = errors.New("transaction_id is empty")
	Error_ConvertFailed        = errors.New("convert failed")
)

type Transaction_PO struct {
	ID                int64           `gorm:"primary_key;auto_increment;comment:'流水號 主鍵'" json:"id"`
	TransferMode      int             `gorm:"type:int(12);comment:'交易模式 0:買 1:賣'" json:"transfer_mode"`
	TransferType      int             `gorm:"type:int(12);comment:'交易種類 0:限價 1:市價'" json:"transaction_type"`
	TransactionID     string          `gorm:"unique;not null; uniqueIndex; comment:'交易訂單'" json:"transaction_id"`
	FromUserID        int64           `gorm:"column:from_user_id; comment:'來源用戶ID'" `
	ToUserID          int64           `gorm:"column:to_user_id; comment:'目的用戶ID'" `
	ProductName       string          `gorm:"size:256;not null; comment:'產品名稱'" json:"product_name"`
	ProductCount      int64           `gorm:"type:bigint(20);default:0; comment:'產品數量'" json:"product_count"`
	ProductNeedAmount decimal.Decimal `gorm:"type:decimal(20,2);default:0; comment:'商品需要的預扣金額'" json:"product_need_amount"`
	Amount            decimal.Decimal `gorm:"type:decimal(20,2);default:0; comment:'實際交易金額'" json:"amount"`
	Currency          string          `gorm:"size:32;not null; comment:'幣種'" json:"currency"`
	CreatedAt         time.Time       `gorm:"autoCreateTime;comment:'創建時間'" json:"created_at"`
	UodateAt          time.Time       `gorm:"autoUpdateTime;comment:'更新時間'" json:"update_at"`
	Status            int8            `gorm:"type:tinyint(1);default:0;comment:'交易狀態 0:未完成 1:已完成 2:取消 3:錯誤'" json:"status"`
}

func (Transaction_PO) TableName() string {
	return "transaction"
}

func (t *Transaction_PO) ToDomain() (*Transaction, error) {

	if t.FromUserID == 0 {
		return nil, Error_UserIDIsEmpty
	}
	if len(t.TransactionID) == 0 {
		return nil, Error_TransactionIDIsEmpty
	}

	user := &Transaction{
		ID:                t.ID,
		TransactionID:     t.TransactionID,
		FromUserID:        t.FromUserID,
		ToUserID:          t.ToUserID,
		ProductName:       t.ProductName,
		ProductCount:      t.ProductCount,
		ProductNeedAmount: t.ProductNeedAmount,
		Amount:            t.Amount,
		Currency:          t.Currency,
		CreatedAt:         t.CreatedAt,
		UodateAt:          t.UodateAt,
		Status:            t.Status,
	}

	return user, nil
}
