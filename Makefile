setup:
	go get github.com/oxequa/realize@master

test:
	DATABASE_DSN=$DATABASE_DSN_TEST go test -v