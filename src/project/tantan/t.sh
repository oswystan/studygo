#!/bin/bash
curl -XPOST -d '{"name":"nick"}' "http://localhost:8000/users"
curl -XPOST -d '{"name":"wy"}' "http://localhost:8000/users"
curl -XGET "http://localhost:8000/users"


curl -XGET "http://localhost:8000/users/123/relationships"
curl -XPUT "http://localhost:8000/users/123/relationships/345"
curl -XPOST "http://localhost:8000/users/123/relationships/345"
