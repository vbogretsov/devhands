#!/usr/bin/awk -f

BEGIN {
    if (length(test_name) == 0) {
        test_name = "default_test";
    }

    rps = 0;
    latency_p99 = 0;
}

/Requests\/sec:/ {
    rps = $2;
}

/99\.000%/ {
    split($2, arr, /[a-z]/);
    latency_value = arr[1];
    latency_unit = substr($2, length(arr[1]) + 1, 1);

    latency_p99 = (latency_unit == "s") ? latency_value * 1000 : latency_value;
}

END {
    printf "{name: \"%s\", rps: %s, l: %.2f}\n", test_name, rps, latency_p99;
}
