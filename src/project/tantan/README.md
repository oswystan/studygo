## install database engine (only for OSX)
`
brew install postgresql
initdb /usr/local/var/postgres -E utf8
pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start
`
## go tool version
go1.5 or later
download binary package from:
http://www.golangtc.com/download

## install 3rd-party package
go get github.com/gorilla/mux
go get gopkg.in/pg.v3

## build
make 

## test
make install
make test

## clean
make clean
