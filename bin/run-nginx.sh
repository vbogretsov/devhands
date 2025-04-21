DURATION=20s
THREADS=8
CONNECTIONS=200
REQUEST_RATES=`seq 200000 5000 220000`
COOLDOWN=5

for i in ${REQUEST_RATES[@]}; do
docker compose run --rm wrkx \
    -c ${CONNECTIONS} \
    -t ${THREADS} \
    -d ${DURATION} \
    -R ${i} \
    -L \
    http://nginx:80/
sleep ${COOLDOWN}
done
