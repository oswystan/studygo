#!/bin/bash

log_file="log.txt"
db="socialdb"

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

create_user()
{
    log "create users ...."
    curl -XPOST -d '{"name":"mark"}' "http://localhost:8000/users"
    curl -XPOST -d '{"name":"sherlock"}' "http://localhost:8000/users"
    curl -XPOST -d '{"name":"john"}' "http://localhost:8000/users"
    curl -XPOST -d '{"name":"tom"}' "http://localhost:8000/users"
    log "done."
}
get_user()
{
    log "get user list..."
    curl -XGET "http://localhost:8000/users"
    log "done."
}

show_db()
{
    log "================================================="
    log "database data:"
    psql -d $db -c "select * from users;"
    psql -d $db -c "select * from relationships;"
    log "================================================="
}

create_rs()
{
    log "create relationships...."
    curl -XPUT -d '{"state":"liked"}' "http://localhost:8000/users/1/relationships/2"
    curl -XPUT -d '{"state":"liked"}' "http://localhost:8000/users/2/relationships/1"
    curl -XPUT -d '{"state":"disliked"}' "http://localhost:8000/users/2/relationships/3"
    log "done."
}

get_rs()
{
    log "get relationships..."
    curl -XGET "http://localhost:8000/users/1/relationships"
    curl -XGET "http://localhost:8000/users/2/relationships"
    log "done."
}

clear_db()
{
    log "clear database data..."
    psql -d $db -c "delete from users; delete from relationships; alter sequence users_id_seq restart with 1" 1>/dev/null
    log "done."
}

rm_db()
{
    psql -c "drop database if exists socialdb"
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
