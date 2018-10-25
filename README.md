cd F:\Projects\Sample\go\src\github.com\tinrab\spidey
//create spidey database, create table in db(table schema is available in account\up.sql file)
//update db connection string in account/cmd/account/main.go. DB URL should be postgres://<username>:<password>@localhost/<database name?sslmode=disable
                                                                                                                                    
//run account micro service

go run account\cmd\account\main.go

//run graphql service
go run graphql/main.go

//play with it
//http://localhost:8082/playground

