package _043__Simple_Bank_System

type Bank struct {
	balance []int64
}

func Constructor(balance []int64) Bank {
	return Bank{
		balance: balance,
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 < 0 || account1 > len(this.balance) {
		return false
	}
	if account2 < 0 || account2 > len(this.balance) {
		return false
	}

	account1 = account1-1
	account2 = account2-1


	account1Balance := this.balance[account1]
	if account1Balance < money {
		return false
	}

	account1Balance = account1Balance - money
	this.balance[account1] = account1Balance
	this.balance[account2] = this.balance[account2] + money

	return true
}


func (this *Bank) Deposit(account int, money int64) bool {
	if account < 0 || account > len(this.balance) {
		return false
	}

	account = account-1

	this.balance[account] = this.balance[account] + money
	return true
}


func (this *Bank) Withdraw(account int, balance int64) bool {
	if account < 0 || account > len(this.balance) {
		return false
	}

	account = account-1

	money := this.balance[account]

	money = money - balance
	if money >= 0 {
		this.balance[account] = money
		return true
	}
	return false
}

