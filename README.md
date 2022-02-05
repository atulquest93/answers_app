# Answers_APP Assignment

## Tech Stack
- Go 
- Gin :- Rest Framework for Go
- Gorm :- ORM for Go
- MySQL :- Database 

## To Run Test Case
Test cases will run sequencially as Update / Delete / Read can only be tested once key is present. <br/>
- `go test -p 1 -v`
<br/> `-p 1 ` will limit concurrent tests to 1, to run them sequencially and -v will print verbose. 

## To Run Code
1. Checkout Code locally
2. Replace contents of .env file
3. Start Application using `go run main.go`

