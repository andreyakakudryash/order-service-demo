#!/bin/bash


echo "=== Order Service WRK Stress Test ==="
echo ""

# Проверка что сервис запущен
echo "Checking if service is running..."
if ! curl -s http://localhost:8080/health > /dev/null; then
    echo "Error: Service is not running on localhost:8080"
    exit 1
fi
echo "Service is running ✓"
echo ""

echo "=== Warmup (10 seconds) ==="
wrk -t2 -c10 -d10s http://localhost:8080/order/b563feb7b2b84b6test

echo ""
echo "=== Test 1: 2 threads, 10 connections, 30 seconds ==="
wrk -t2 -c10 -d30s --latency http://localhost:8080/order/b563feb7b2b84b6test

echo ""
echo "=== Test 2: 4 threads, 50 connections, 30 seconds ==="
wrk -t4 -c50 -d30s --latency http://localhost:8080/order/b563feb7b2b84b6test

echo ""
echo "=== Test 3: 8 threads, 100 connections, 30 seconds ==="
wrk -t8 -c100 -d30s --latency http://localhost:8080/order/b563feb7b2b84b6test

echo ""
echo "=== Test 4: 12 threads, 200 connections, 30 seconds ==="
wrk -t12 -c200 -d30s --latency http://localhost:8080/order/b563feb7b2b84b6test

echo ""
echo "=== Test 5: Mixed endpoints with Lua script ==="

cat > wrk_script.lua << 'EOF'

order_ids = {
    "b563feb7b2b84b6test",
    "U2RuKhsnAs52455rtest",
    "1xF5TS7yiwGkmMwHtest",
    "oNiDeZMp3W7QWJzxtest",
    "SMjNghpbzYPuRRb4test",
    "cHHKGKA1o5M4SWfRtest"
}

counter = 0

request = function()
    counter = counter + 1
    local idx = (counter % #order_ids) + 1
    local path = "/order/" .. order_ids[idx]
    return wrk.format("GET", path)
end
EOF

wrk -t4 -c50 -d30s --latency -s wrk_script.lua http://localhost:8080

rm wrk_script.lua

echo ""
echo "=== WRK stress test completed ==="