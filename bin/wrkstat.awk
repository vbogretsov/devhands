#!/usr/bin/awk -f

BEGIN {
    if (length(test_name) == 0) {
        test_name = "default_test";
    }

    rps = 0;
    latency_p99 = 0;
}

FNR == 1 {
    if (FILENAME ~ /-[0-9]+\./) {
        n = split(FILENAME, parts, "-");
        requests = parts[n];
    }
}

/Requests\/sec:/ {
    rps = $2;
}

/99\.000%/ {
    split($2, arr, /[a-z]/);
    latency_value = arr[1];
    latency_unit = substr($2, length(arr[1]) + 1, 2);

    if (latency_unit == "s") {
        latency_p99 = latency_value * 1000;
    } else if (latency_unit == "m") {
        latency_p99 = latency_value * 1000 * 60;
    } else {
        latency_p99 = latency_value;
    }
}

END {
    printf "{Test: \"%s\", R: %s, rps: %s, l: %.2f}\n", test_name, requests, rps, latency_p99;
}
