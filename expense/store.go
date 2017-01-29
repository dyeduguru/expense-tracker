package expense

import (
	"database/sql"
	"github.com/palantir/stacktrace"
	"github.com/dyeduguru/expense-tracker/api"
)

type Store struct{
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db:db}
}

func (s *Store) Create(exp *api.Expense) error {
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

func (s *Store) ReadAll() (api.Expenses, error) {
	rows, err := s.db.Query("select * from expenses;")
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	defer rows.Close()
	return getExpensesFromRows(rows)
}

func (s *Store) Read(id string) (*api.Expense, error) {
	rows, err := s.db.Query("select * from expenses where $1=$2;", "id", id)
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	defer rows.Close()
	expenses, err := getExpensesFromRows(rows)
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	if len(expenses) != 1 {
		return nil, stacktrace.NewError("Unexpexted number of matches")
	}
	return expenses[0], nil
}

func (s *Store) Update(exp *api.Expense) error {
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

func (s *Store) Delete(id string) error {
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