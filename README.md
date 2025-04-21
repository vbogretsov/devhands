# devhands

My repo for devhands hands on practice

TODO: Write readme

## Commands

ls log/nginx-* | xargs -n 1 bin/wrkstat.awk -v test_name=test | sed 's/$/,/g' | sort -t ',' -n -k2.4


bin/lsmt.sh -t 4-4-1 -c100-100-1 -R 5000-100000-5000 -d 60 -p app-wait-10 -L http://localhost/api/wait/10


dp run --rm wrk -c 400 -t 8 -d 60s -s /src/post_spots.lua http://app-1:8000/api/spots -- /data/admin.jwt /data/spots
