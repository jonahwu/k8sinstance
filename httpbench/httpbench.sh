echo "GET http://localhost:8000/sleep" | vegeta -cpus 2 attack -header 'SLEEPTIME: 3' -duration=1s -timeout=1s | vegeta report
