for run in {1..50}; do a=$(( ( RANDOM % 100 )  + 1 )) ; b=$(( ( RANDOM % 100 )  + 1 )) ;hey -n 5 -c 2 -m POST -H 'accept: application/json' -H 'Content-Type: application/json' -d "{ \"limit\": "${a}", \"offset\": "${b}"}" http://0.0.0.0:8083/api/v1/equipment_requests/list ; done

