#!/bin/bash

db="socialdb"

# clear database
psql -d $db -c "delete from users; delete from relationships; alter sequence users_id_seq restart with 1"


curl -XPOST -d '{"name":"mark"}' "http://localhost:8000/users"
curl -XPOST -d '{"name":"sherlock"}' "http://localhost:8000/users"
curl -XPOST -d '{"name":"john"}' "http://localhost:8000/users"
curl -XPOST -d '{"name":"tom"}' "http://localhost:8000/users"

echo "=============================================="
psql -d $db -c "select * from users;"
curl -XGET "http://localhost:8000/users"
echo "=============================================="


curl -XPUT -d '{"state":"liked"}' "http://localhost:8000/users/1/relationships/2"
curl -XPUT -d '{"state":"liked"}' "http://localhost:8000/users/2/relationships/1"
curl -XPUT -d '{"state":"disliked"}' "http://localhost:8000/users/2/relationships/3"
psql -d $db -c "select * from relationships;"
echo "=============================================="
curl -XGET "http://localhost:8000/users/1/relationships"
echo "=============================================="
curl -XGET "http://localhost:8000/users/2/relationships"
echo "=============================================="

# for abnormal test
curl -XPOST "http://localhost:8000/users/123/relationships/345"
