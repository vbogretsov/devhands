#!/usr/bin/awk -f

BEGIN {
    # Set default test name if not provided via -v test_name=
    if (length(test_name) == 0) {
        test_name = "default_test";
    }
    
    # Initialize metrics
    rps = 0;
    latency_p99 = 0;
}

/Requests\/sec:/ {
    rps = $2;
}

/99\.000%/ {
    latency_value = $2;
    latency_unit = $3;
    
    # Convert to milliseconds if in seconds
    latency_p99 = (latency_unit == "s") ? latency_value * 1000 : latency_value;
}

END {
    # Print in JSON-like format with 2 decimal places for latency
    printf "{Test: %s, rps: %s, l: %.2f}\n", test_name, rps, latency_p99;
}
