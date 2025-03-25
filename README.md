# devhands

My repo for devhands hands on practice

TODO: Write readme

## Commands

s log/root-t4-c100-* | xargs -n 1 bin/wrkstat.awk -v test_name=nginx-t4-c100 | sed 's/$/,/g' | sort -t ',' -n -k2.4


bin/lsmt.sh -t 4-4-1 -c100-100-1 -R 10000-100000-10000 -d 60 http://localhost/

