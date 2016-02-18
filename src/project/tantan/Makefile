##########################################################
## 
##########################################################
.PHONY: all build test install clean st

all: build 

build:
	@ echo "build tantan..."
	@ go build
	@ echo "done."

install:
	@ echo "installing ..."
	@ psql -U postgres -f ./pg.sql
	@ echo "done."

test: install
	@ echo "start testing ..."
	@ go test
	@ echo "done."

st: install
	@ echo "start st ..."
	@ ./tantan &
	@ sleep 1
	@ ./st.sh
	@ killall tantan
	@ echo "done."

clean:
	@ echo "cleaning ..."
	@ rm -f tantan
	@ ./st.sh rm_db
	@ echo "done."
##########################################################
