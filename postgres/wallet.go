package postgres

import (
	"time"

	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
)

type Wallet struct {
	ID         int       `postgres:"id"`
	UserID     int       `postgres:"user_id"`
	UserName   string    `postgres:"user_name"`
	WalletName string    `postgres:"wallet_name"`
	WalletType string    `postgres:"wallet_type"`
	Balance    float64   `postgres:"balance"`
	CreatedAt  time.Time `postgres:"created_at"`
}

func (p *Postgres) Wallets(walletType string) ([]wallet.Wallet, error) {
	queryString := "SELECT * FROM user_wallet"
	var args []interface{}
	if walletType != "" {
		queryString += " WHERE wallet_type = $1"
		args = append(args, walletType)
	}

	rows, err := p.Db.Query(queryString, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wallets []wallet.Wallet
	for rows.Next() {
		var w Wallet
		err := rows.Scan(&w.ID,
			&w.UserID, &w.UserName,
			&w.WalletName, &w.WalletType,
			&w.Balance, &w.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		wallets = append(wallets, wallet.Wallet{
			ID:         w.ID,
			UserID:     w.UserID,
			UserName:   w.UserName,
			WalletName: w.WalletName,
			WalletType: w.WalletType,
			Balance:    w.Balance,
			CreatedAt:  w.CreatedAt,
		})
	}
	return wallets, nil
}

func (p *Postgres) Wallet(id uint64) (wallet.Wallet, error) {
	queryString := "SELECT * FROM user_wallet WHERE id = $1"
	row := p.Db.QueryRow(queryString, id)

	var w Wallet
	err := row.Scan(&w.ID,
		&w.UserID, &w.UserName,
		&w.WalletName, &w.WalletType,
		&w.Balance, &w.CreatedAt,
	)
	if err != nil {
		return wallet.Wallet{}, err
	}
	return wallet.Wallet{
		ID:         w.ID,
		UserID:     w.UserID,
		UserName:   w.UserName,
		WalletName: w.WalletName,
		WalletType: w.WalletType,
		Balance:    w.Balance,
		CreatedAt:  w.CreatedAt,
	}, nil
}

func (p *Postgres) CreateWallet(w wallet.WalletPayload) (int, error) {
	queryString := "INSERT INTO user_wallet(user_id, user_name, wallet_name, wallet_type, balance) VALUES($1, $2, $3, $4, $5) RETURNING id"
	row := p.Db.QueryRow(queryString, w.UserID, w.UserName, w.WalletName, w.WalletType, w.Balance)

	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
