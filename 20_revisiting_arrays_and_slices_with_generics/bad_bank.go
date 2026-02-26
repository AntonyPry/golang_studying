package main

type Transaction struct {
	From string
	To   string
	Sum  float64
}

type Account struct {
	Name    string
	Balance float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func applyTransaction(account Account, t Transaction) Account {
	if t.From == account.Name {
		account.Balance -= t.Sum
	}
	if t.To == account.Name {
		account.Balance += t.Sum
	}
	return account
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(transactions, applyTransaction, account)
}

func Find[A any](collection []A, f func(A) bool) (value A, found bool) {
	for _, value := range collection {
		if f(value) == true {
			return value, true
		}
	}
	return
}
