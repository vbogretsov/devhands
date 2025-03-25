#!/usr/bin/awk -f

BEGIN {
    # Initialize variables
    rps = 0;
    latency_p99 = 0;
    test_name = "'$1'"; # Get test name from first argument or set default
    if (test_name == "''") {
        test_name = "default_test";
    }
}

/Requests\/sec:/ {
    # Extract RPS value
    rps = $2;
}

/99.000%/ {
    # Extract 99th percentile latency
    latency_value = $2;
    latency_unit = $3;
    
    # Convert to milliseconds if in seconds
    if (latency_unit == "s") {
        latency_p99 = latency_value * 1000;
    } else {
        latency_p99 = latency_value;
    }
}

END {
    # Print in the requested format
    printf "{Test: %s, rps: %s, l: %.2f}\n", test_name, rps, latency_p99;
}
