echo "GET http://localhost:8112/" | vegeta attack -rate 1/100ms -duration=1m -max-connections=10 | tee results.bin | vegeta report
  vegeta report -type=json results.bin > metrics.json
  cat results.bin | vegeta plot > plot.html
  cat results.bin | vegeta report -type="hist[0,100ms,200ms,300ms]"
