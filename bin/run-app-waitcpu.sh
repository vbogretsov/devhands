#!/bin/sh

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
    http://app-1:8000/wait/1
sleep ${COOLDOWN}
done
