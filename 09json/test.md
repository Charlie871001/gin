curl http://localhost:8080/loginJSON -H 'content-type:application/json' -d {"user":"root","password":"admin"} -X POST

curl http://localhost:8080/loginJSON -H 'content-type:application/json' -d "{\"user\":\"root\",\"password\":\"admin\"}" -X POST

curl http://localhost:8080/loginJSON -H 'content-type:application/json' -d "{\"user\":\"root\",\"password\":\"admin1\"}" -X POST