package stores

import (
	"database/sql"
	"github.com/dyeduguru/expense-tracker/api"
	"github.com/palantir/stacktrace"
)

type UserStore struct{
	db *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{db:db}
}

func (s *UserStore) Create(user *api.User) error {
	query := `INSERT INTO users(id, admin, username, password, name)
	values ($1,$2,$3,$4,$5)`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		stacktrace.Propagate(err, "")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Admin, user.UserName,user.Password,user.Name)
	if err != nil {
		return stacktrace.Propagate(err, "")
	}
	return nil
}

func (s *UserStore) Read(username string) (*api.User, error) {
	rows, err := s.db.Query("select * from users where username = $1;", username)
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	defer rows.Close()
	users, err := getUsersFromRows(rows)
	if err != nil {
		return nil, stacktrace.Propagate(err, "")
	}
	if len(users) != 1 {
		return nil, stacktrace.NewError("Unexpexted number of matches with username:%v. %v macthed",username, len(users))
	}
	return users[0], nil
}

func (s *UserStore) Update(user *api.User) error {
	query := `update users set admin=$2,username=$3,password=$4,name=$5 where id =$1`
	stmt, err := s.db.Prepare(query)
	if err != nil {
		stacktrace.Propagate(err, "")
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.Id, user.Admin, user.UserName, user.Password, user.Name)
	if err != nil {
		return stacktrace.Propagate(err, "")
	}
	return nil
}

func (s *UserStore) Delete(id string) error {
	query := `delete from users where id =$1`
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