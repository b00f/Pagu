package voucher

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/pagu-project/Pagu/internal/engine/command"
	"github.com/pagu-project/Pagu/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestStatusNormal(t *testing.T) {
	voucher, db, _, _ := setup(t)

	t.Run("one code status normal", func(t *testing.T) {
		now := time.Now()
		validMonths := uint8(2)

		db.EXPECT().GetVoucherByCode("12345678").Return(
			entity.Voucher{
				ID:          1,
				Code:        "12345678",
				Desc:        "some_desc",
				Recipient:   "some_recipient",
				ValidMonths: validMonths,
				Amount:      uint(100),
				TxHash:      "some_transaction_hash",
				ClaimedBy:   0,
				Model: gorm.Model{
					CreatedAt: now,
				},
			}, nil,
		).AnyTimes()

		expTime := now.AddDate(0, int(validMonths), 0).Format("02/01/2006, 15:04:05")

		cmd := command.Command{
			User: &entity.User{
				ID: 1,
			},
		}

		result := voucher.statusHandler(cmd, entity.AppIdDiscord, "", "12345678")
		assert.True(t, result.Successful)
		assert.Equal(t, result.Message, fmt.Sprintf("Code: 12345678\nAmount: 100 PAC\n"+
			"Expire At: %s\nRecipient: some_recipient\nDescription: some_desc\nClaimed: YES\nTx Link: https://pacviewer.com/transaction/some_transaction_hash"+
			"\n", expTime))
	})

	t.Run("wrong code", func(t *testing.T) {
		db.EXPECT().GetVoucherByCode("invalid_code").Return(
			entity.Voucher{}, errors.New(""),
		).AnyTimes()

		cmd := command.Command{
			User: &entity.User{
				ID: 1,
			},
		}

		result := voucher.statusHandler(cmd, entity.AppIdDiscord, "", "invalid_code")
		assert.False(t, result.Successful)
		assert.Equal(t, result.Message, "An error occurred: voucher code is not valid, no voucher found")
	})

	t.Run("list vouchers status normal", func(t *testing.T) {
		now := time.Now()
		validMonths := uint8(2)

		db.EXPECT().ListVoucher().Return(
			[]*entity.Voucher{
				{
					ID:          1,
					Code:        "code1",
					ValidMonths: validMonths,
					Amount:      uint(100),
					TxHash:      "some_transaction_hash",
					Model: gorm.Model{
						CreatedAt: now,
					},
				},
				{
					ID:          2,
					Code:        "code2",
					ValidMonths: validMonths,
					Amount:      uint(100),
					Model: gorm.Model{
						CreatedAt: now,
					},
				},
				{
					ID:          3,
					Code:        "code3",
					ValidMonths: validMonths,
					Amount:      uint(100),
					Model: gorm.Model{
						CreatedAt: now.AddDate(0, -3, 0),
					},
				},
			}, nil,
		).AnyTimes()

		cmd := command.Command{
			User: &entity.User{
				ID: 1,
			},
		}

		result := voucher.statusHandler(cmd, entity.AppIdDiscord, "")
		assert.True(t, result.Successful)
		assert.Equal(t, result.Message, "Total Codes: 3\nTotal Amount: 300 PAC\n\n\n"+
			"Claimed: 1\nTotal Claimed Amount: 100\nTotal Expired: 1"+
			"\n")
	})
}
