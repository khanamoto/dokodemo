setup:
	go get github.com/oxequa/realize@master
#	go get -u github.com/gorilla/mux
#	go get -u github.com/go-sql-driver/mysql
#	go get -u github.com/jmoiron/sqlx
#	go get -u golang.org/x/crypto/bcrypt

test:
	DATABASE_DSN=$DATABASE_DSN_TEST go test -v