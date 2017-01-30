package stores

import (
	"database/sql"
	"github.com/palantir/stacktrace"
	"github.com/dyeduguru/expense-tracker/api"
)

type ExpenseStore struct{
	db *sql.DB
}

func NewExpenseStore(db *sql.DB) *ExpenseStore {
	return &ExpenseStore{db:db}
}

func (s *ExpenseStore) Create(exp *api.Expense) error {
	query := `INSERT INTO expenses(id, userid, amount, description, timestamp)
	values ($1,$2,$3,$4,$5)`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		stacktrace.Propagate(err, "")
	}
	defer stmt.Close()
	_, err = stmt.Exec(exp.Id, exp.UserId, exp.Amount,exp.Description,exp.Timestamp)
	if err != nil {
		return stacktrace.Propagate(err, "")
	}
	return nil
}

func (s *ExpenseStore) ReadAll() (api.Expenses, error) {
	rows, err := s.db.Query("select * from expenses;")
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	defer rows.Close()
	return getExpensesFromRows(rows)
}

func (s *ExpenseStore) Read(userid string) (api.Expenses, error) {
	rows, err := s.db.Query("select * from expenses where $1=$2;", "id", userid)
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	defer rows.Close()
	return getExpensesFromRows(rows)
}

func (s *ExpenseStore) Update(exp *api.Expense) error {
	query := `update expenses set userid=$2,amount=$3,description=$4,timestamp=$5 where id =$1`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		stacktrace.Propagate(err, "")
	}
	defer stmt.Close()
	_, err = stmt.Exec(exp.Id, exp.UserId, exp.Amount,exp.Description,exp.Timestamp)
	if err != nil {
		return stacktrace.Propagate(err, "")
	}
	return nil
}

func (s *ExpenseStore) Delete(id string) error {
	query := `delete from expenses where id =$1`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		stacktrace.Propagate(err, "")
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return stacktrace.Propagate(err, "")
	}
	return nil
}

func getExpensesFromRows(rows *sql.Rows) (api.Expenses, error) {
	expenses := make(api.Expenses,0)
	for rows.Next() {
		expense := &api.Expense{}
		err := rows.Scan(&expense.Id,&expense.UserId, &expense.Amount, &expense.Description, &expense.Timestamp)
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