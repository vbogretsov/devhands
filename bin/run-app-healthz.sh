DURATION=20s
THREADS=8
CONNECTIONS=200
REQUEST_RATES=`seq 140000 2500 150000`
COOLDOWN=5

for i in ${REQUEST_RATES[@]}; do
docker compose run --rm wrkx \
    -c ${CONNECTIONS} \
    -t ${THREADS} \
    -d ${DURATION} \
    -R ${i} \
    -L \
    http://app-1:8000/healthz
sleep ${COOLDOWN}
done
