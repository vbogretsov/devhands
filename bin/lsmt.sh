#!/bin/bash

# Function to display usage
usage() {
    echo "Usage: $0 -t <thread_range> -c <connection_range> -R <request_rate_range> -d <duration> -L <url>"
    echo "  Ranges should be in the format: start-end-step"
    echo "Example: $0 -t 1-4-1 -c 10-100-10 -R 1000-5000-1000 -d 120s -L -p root http://localhost/"
    exit 1
}

# Check if no arguments are provided
if [ $# -eq 0 ]; then
    usage
fi

# Parse command-line arguments
while getopts ":t:c:R:d:L:" opt; do
    case $opt in
        t) THREAD_RANGE="$OPTARG" ;;
        c) CONNECTION_RANGE="$OPTARG" ;;
        R) REQUEST_RATE_RANGE="$OPTARG" ;;
        d) DURATION="$OPTARG" ;;
        L) URL="$OPTARG" ;;
        p) PREFIX="$OPTARG" ;;
        *) usage ;;
    esac
done

# Validate required arguments
if [ -z "$THREAD_RANGE" ]; then
    echo "Error: missing required argument -t"
    usage
fi

if [ -z "$CONNECTION_RANGE" ]; then
    echo "Error: missing required argument -c"
    usage
fi

if [ -z "$REQUEST_RATE_RANGE" ]; then
    echo "Error: missing required argument -R"
    usage
fi

if [ -z "$DURATION" ]; then
    echo "Error: missing required argument -d"
    usage
fi

if [ -z "$PREFIX" ]; then
    echo "Error: missing required argument -p"
    usage
fi

# Function to generate a range of numbers with a step
generate_range() {
    local range=$1
    local start end step
    IFS='-' read -r start end step <<< "$range"
    seq "$start" "$step" "$end"
}

# Generate ranges for threads, connections, and request rates
THREADS=($(generate_range "$THREAD_RANGE"))
CONNECTIONS=($(generate_range "$CONNECTION_RANGE"))
REQUEST_RATES=($(generate_range "$REQUEST_RATE_RANGE"))

# Iterate through all combinations of parameters
for t in "${THREADS[@]}"; do
    for c in "${CONNECTIONS[@]}"; do
        for R in "${REQUEST_RATES[@]}"; do
            echo "Running wrk with -t$t -c$c -R$R -d$DURATION -L$URL"
            /local/wrk -t"$t" -c"$c" -d"$DURATION" -R"$R" -L "$URL" > ./log/${PREFIX}-t${t}-c${c}-R${R}.log
            sleep 20
            echo "---------------------------------------------"
        done
    done
done

echo "All tests completed."
