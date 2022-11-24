package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(wallet.Bitcoin(10))
		assertBalance(t, wallet, wallet.Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(wallet.Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, wallet.Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(wallet.Bitcoin(100))

		assertError(t, err, wallet.ErrInsufficientFunds)
		assertBalance(t, wallet, wallet.Bitcoin(20))

	})
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("didn't get am error but wanted one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
