## install database engine (only for OSX)
`
brew install postgresql
initdb /usr/local/var/postgres -E utf8
pg_ctl -D /usr/local/var/postgres -l /usr/local/var/postgres/server.log start
`

## build
make 

## test
make install
make test
