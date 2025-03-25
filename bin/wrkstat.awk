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
    latency_value = $2;
    latency_unit = $3;

    printf "latency_unit=%s\n", latency_unit;
    latency_p99 = (latency_unit == "s") ? latency_value * 1000 : latency_value;
}

END {
    printf "{name: \"%s\", rps: %s, l: %.2f}\n", test_name, rps, latency_p99;
}
