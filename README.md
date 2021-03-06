# Expense Tracker

Simple App to track expenses. Uses REST api with JWT auth written in go in the backend and a React frontend

# Setup
* For setting up this project, go get using the following command
```
go get github.com/dyeduguru/expense-tracker
```
* Configure postgres sql connection in main.go
* Run the following command to build the executable:
```
go install main.go
```
*Run the executable. This runs the server on https://localhost:3000

* The following endpoints are available:

```
POST /user Creates new user

{
	Id       string
	Admin    bool
	UserName string
	Password string
	Name     string
}

POST /token Creates a token for the user

{
	username	string
	password	string
}

GET /expense Lists expenses

POST /expense Creates expense

{
	Id          string
	Timestamp   time.Time
	Amount      float32
	Description string
}

PATCH /expense Updates expense

{
	Id          string
	Timestamp   time.Time
	Amount      float32
	Description string
}

PUT /expense Deletes expense

{
	Id          string
	Timestamp   time.Time
	Amount      float32
	Description string
}

```
