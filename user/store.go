package user

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

func (s *Store) Get(userid string) (api.User, error) {
	rows, err := s.db.Query("select * from users where userid=$1;", userid)
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	defer rows.Close()
	return getUsersFromRows(rows)
}

func getUsersFromRows(rows *sql.Rows) (api.Users, error) {
	users := make(api.Users,0)
	for rows.Next() {
		user := &api.User{}
		err := rows.Scan(&user.Id, &user.Admin, &user.UserName, &user.Password, &user.Name)
		if err != nil {
			return nil, stacktrace.Propagate(err, "")
		}
		users = append(users, user)
	}
	err := rows.Err()
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	return users, nil
}