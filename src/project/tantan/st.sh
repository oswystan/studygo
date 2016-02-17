#!/bin/bash

log_file="log.txt"
db="socialdb"

create_user()
{
    echo "create users ...."
    curl -XPOST -d '{"name":"mark"}' "http://localhost:8000/users"
    curl -XPOST -d '{"name":"sherlock"}' "http://localhost:8000/users"
    curl -XPOST -d '{"name":"john"}' "http://localhost:8000/users"
    curl -XPOST -d '{"name":"tom"}' "http://localhost:8000/users"
    echo "done."
}
get_user()
{
    echo "get user list..."
    curl -XGET "http://localhost:8000/users"
    echo "done."
}

show_db()
{
    psql -d $db -c "select * from users;"
    psql -d $db -c "select * from relationships;"
}

create_rs()
{
    echo "create relationships...."
    curl -XPUT -d '{"state":"liked"}' "http://localhost:8000/users/1/relationships/2"
    curl -XPUT -d '{"state":"liked"}' "http://localhost:8000/users/2/relationships/1"
    curl -XPUT -d '{"state":"disliked"}' "http://localhost:8000/users/2/relationships/3"
    echo "done."
}

get_rs()
{
    echo "get relationships..."
    curl -XGET "http://localhost:8000/users/1/relationships"
    curl -XGET "http://localhost:8000/users/2/relationships"
    echo "done."
}


################################
# some basic functions for log
################################
log_start()
{
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}]##########################################################"
    echo "[${strNow}] start program  : $0"
    echo "[${strNow}]##########################################################"
    echo ""
}

log()
{
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo "[${strNow}] INFO: $*"
}

log_end()
{
    strNow=`date +'%Y-%m-%d %H:%M:%S'`
    echo ""
    echo "[${strNow}]##########################################################"
    echo "[${strNow}] finished $0"
    echo "[${strNow}]##########################################################"
}

clear_db()
{
    psql -d $db -c "delete from users; delete from relationships; alter sequence users_id_seq restart with 1" 1>/dev/null
}

do_work()
{
    log_start

    clear_db
    create_user
    get_user
    create_rs
    get_rs
    show_db

    log_end
}

################################
## main
################################
if [ $# -eq 0 ]; then
    do_work 
else
    $*
fi
