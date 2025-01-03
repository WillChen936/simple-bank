# Simple Bank
A Go-based backend for managing a simple banking system. This project implements fundamental banking features such as account creation, fund transfers, and transaction tracking, using a PostgreSQL database and best practices in software development.


## Features
- **User Authentication**: Secure login with hashed passwords.
- **Account Management**: Create and manage multiple accounts for users.
- **Transactions**: Perform deposits, withdrawals, and transfers between accounts.
- **Database Support**: Built with PostgreSQL for data persistence.
- **Test Coverage**: Comprehensive unit and integration tests for reliability.
- **Secure**: Follows best practices for data security and application architecture.


## Technologies
- **Programming Language**: Go (Golang)
- **Database**: PostgreSQL
- **ORM**: SQLC for type-safe database queries
- **Testing**: Go’s built-in testing tools and libraries

## Directory Structure
```
├── api          # API layer and HTTP handlers
├── db           # Database migrations and SQL queries
├── utils        # Utility functions and helpers
├── main.go      # Application entry point
└── README.md    # Project documentation
```
