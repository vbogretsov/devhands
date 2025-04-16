DURATION=10s
THREADS=4
CONNECTIONS=100
REQUEST_RATES=
COOLDOWN=5
ITERS=`seq 1 1 4`

for i in ${ITERS[@]}; do
docker compose run --rm wrk \
    -c ${CONNECTIONS} \
    -t ${THREADS} \
    -d ${DURATION} \
    -s /src/get_locations.lua \
    http://app-1:8000/api/tracker/locations -- /data/admin.jwt 1000000
sleep ${COOLDOWN}
done
