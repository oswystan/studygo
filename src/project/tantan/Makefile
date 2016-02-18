##########################################################
## 
##########################################################
.PHONY: all build test install clean st

all: build 

build: tantan

tantan: 
	@ go build

install: tantan
	@ echo "installing ..."
	@ psql -f ./pg.sql
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
