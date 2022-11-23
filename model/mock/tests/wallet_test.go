package tests

import (
	"TDD/model/mock"
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := mock.Wallet{}
		wallet.Deposit(mock.Bitcoin(10))
		assertBalance(t, wallet, mock.Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := mock.Wallet{mock.Bitcoin(20)}
		err := wallet.Withdraw(mock.Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, mock.Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := mock.Wallet{mock.Bitcoin(20)}
		err := wallet.Withdraw(mock.Bitcoin(100))

		assertError(t, err, mock.ErrInsufficientFunds)
		assertBalance(t, wallet, mock.Bitcoin(20))

	})
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet mock.Wallet, want mock.Bitcoin) {
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
