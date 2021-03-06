package main

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func (t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(30))
	})

	t.Run("Withdraw", func (t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with insuffcient amounts", func (t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(30))

		assertError(t, err, InsufficientFundsError)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}

	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}