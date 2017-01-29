package expense

import (
	"database/sql"
	"github.com/palantir/stacktrace"
)

type Store struct{
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db:db}
}

func (s *Store) GetAll() (Expenses, error) {
	rows, err := s.db.Query("select * from expenses;")
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	defer rows.Close()
	return getExpensesFromRows(rows)
}

func (s *Store) Get(userid string) (Expenses, error) {
	rows, err := s.db.Query("select * from expenses where userid=$1;", userid)
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	defer rows.Close()
	return getExpensesFromRows(rows)
}

func getExpensesFromRows(rows *sql.Rows) (Expenses, error) {
	expenses := make(Expenses,0)
	for rows.Next() {
		expense := &Expense{}
		err := rows.Scan(&expense.UserId, &expense.Amount, &expense.Description, &expense.Timestamp)
		if err != nil {
			return nil, stacktrace.Propagate(err, "")
		}
		expenses = append(expenses, expense)
	}
	err := rows.Err()
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	return expenses, nil
}