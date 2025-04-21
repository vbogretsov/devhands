DURATION=10s
THREADS=4
CONNECTIONS=100
REQUEST_RATES=`seq 30000 2500 45000`
COOLDOWN=5

for r in ${REQUEST_RATES[@]}; do
docker compose run --rm wrkx \
    -c ${CONNECTIONS} \
    -t ${THREADS} \
    -d ${DURATION} \
    -R ${r} \
    -L \
    -s /src/post_spots.lua \
    http://app-1:8000/api/delivery/spots -- /data/admin.jwt
sleep ${COOLDOWN}
done
