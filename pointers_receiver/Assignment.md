## Assignment

### Bank Account Management with Pointers and Receivers

#### Objective

The objective of this task is to understand and effectively use pointers and receivers in Go. You will implement a simple bank account management system to practice modifying and accessing data through methods with value and pointer receivers.

#### Problem Statement

You need to create a `BankAccount` struct that represents a bank account with basic functionalities. This includes getting account information, depositing money, and withdrawing money. The goal is to implement methods using both value receivers and pointer receivers to manipulate the account balance. Note that this task does not address challenges related to concurrency or multiple read/writes to the same resource; the focus is on understanding pointers and receivers.

#### Requirements

Define a <code>BankAccount</code> Struct</strong>:

- Fields: `AccountNumber` (string) and `Balance` (float64).

Implement Methods:

- `GetAccountInfo` (value receiver): Returns a string with the account number and balance.
- `Deposit` (pointer receiver): Adds an amount to the balance, ensuring the amount is positive.
- `Withdraw` (pointer receiver): Subtracts an amount from the balance, ensuring the amount is positive and does not exceed the current balance.

Create a <code>main</code> Function</strong>:

- Initialize a `BankAccount` instance with an account number and an initial balance.
- Call the `GetAccountInfo` method to display the account details.
- Perform several deposits and withdrawals, including valid and invalid operations.
- Print the updated account details after each operation.
