for run in {1..50}; do a=$(( ( RANDOM % 10000 )  + 1 )) ; b=$(( ( RANDOM % 10000 )  + 1 )) ;hey -n 5 -c 2 -m POST -H 'accept: application/json' -H 'Content-Type: application/json' -d "{ \"equipment_request_id\": "${a}", \"equipmentId\": "${b}"}" http://0.0.0.0:8083/api/v1/update/equipment_id ; done