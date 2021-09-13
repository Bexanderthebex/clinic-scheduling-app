curl -i -X PUT -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:9200/hospitals -d @mappings/hospitals.json

curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @connectors/users.json
curl -i -X POST -H "Accept:application/json" -H  "Content-Type:application/json" http://localhost:8083/connectors/ -d @connectors/source.json