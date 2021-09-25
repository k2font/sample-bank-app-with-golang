package bank

import "testing"

func TestAccount(t *testing.T) {
	// TDD
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los",
			Phone:   "000-0000-0000",
		},
		Number:  1001,
		Balance: 0.0,
	}

	if account.Name == "" {
		t.Error("Cant create an Account object")
	}
}

// 口座にお金を追加できるメソッドのテスト
func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "Shoichiro",
			Address: "Tokyo",
			Phone:   "000-0000-0000",
		},
		Number:  1000,
		Balance: 0,
	}

	account.Deposit(10)
	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositInvalid(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "shoichiro",
			Address: "Tokyo",
			Phone:   "000-0000-0000",
		},
		Number:  1001,
		Balance: 0,
	}
	if err := account.Deposit(-10); err == nil {
		t.Error("only positive number ...")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "shoichiro",
			Address: "Tokyo",
			Phone:   "000-0000-0000",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(10)

	if account.Balance != 0 {
		t.Error("エラー")
	}
}

func TestSatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Error("エラー")
	}
}

func TestSendMoney(t *testing.T) {
	account1 := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	account2 := Account{
		Customer: Customer{
			Name:    "Salah",
			Address: "Tokyo",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account1.Deposit(100)                    // 指定した口座番号に送金
	err := account1.SendMoney(50, &account2) // 指定した口座番号に送金
	if account1.Balance != 50 && account2.Balance != 50 {
		t.Error("エラー", err)
	}
}
