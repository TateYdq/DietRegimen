sudo kill -9 $(lsof -i:8080 |awk '{print $2}' | tail -n 2)