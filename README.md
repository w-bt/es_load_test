# Load Test ES

Script for load testing ES

### Requirements

- Golang 1.21
- ElasticSearch 6.8 or later

# How it works

## Command
- ./out/es_load_test create_index {index name}
- ./out/es_load_test server -> to start server


### USE CASE 1 POPULATE SINGLE DOC

echo "POST http://localhost:8080/es_load/item" | vegeta attack -rate=1000 -max-workers=1000 -duration=300s | tee results_usecase_1000rps_1.bin | vegeta report

vegeta report -type=json results_usecase_1000rps_1.bin > metrics_usecase_1000rps_1.json
cat results_usecase_1000rps_1.bin | vegeta plot > plot_usecase_1000rps_1.html
cat results_usecase_1000rps_1.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/item" | vegeta attack -rate=2000 -max-workers=1000 -duration=300s | tee results_usecase_2000rps_1.bin | vegeta report

vegeta report -type=json results_usecase_2000rps_1.bin > metrics_usecase_2000rps_1.json
cat results_usecase_2000rps_1.bin | vegeta plot > plot_usecase_2000rps_1.html
cat results_usecase_2000rps_1.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/item" | vegeta attack -rate=3000 -max-workers=1000 -duration=300s | tee results_usecase_3000rps_1.bin | vegeta report

vegeta report -type=json results_usecase_3000rps_1.bin > metrics_usecase_3000rps_1.json
cat results_usecase_3000rps_1.bin | vegeta plot > plot_usecase_3000rps_1.html
cat results_usecase_3000rps_1.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

/*
* USE CASE 1.1 POPULATE SINGLE DOC WITH REFRESH
  */

echo "POST http://localhost:8080/es_load/item?refresh=true" | vegeta attack -rate=1000 -max-workers=1000 -duration=300s | tee results_usecase_1000rps_1.1.bin | vegeta report

vegeta report -type=json results_usecase_1000rps_1.1.bin > metrics_usecase_1000rps_1.1.json
cat results_usecase_1000rps_1.1.bin | vegeta plot > plot_usecase_1000rps_1.1.html
cat results_usecase_1000rps_1.1.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/item?refresh=true" | vegeta attack -rate=2000 -max-workers=1000 -duration=300s | tee results_usecase_2000rps_1.1.bin | vegeta report

vegeta report -type=json results_usecase_2000rps_1.1.bin > metrics_usecase_2000rps_1.1.json
cat results_usecase_2000rps_1.1.bin | vegeta plot > plot_usecase_2000rps_1.1.html
cat results_usecase_2000rps_1.1.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/item?refresh=true" | vegeta attack -rate=3000 -max-workers=1000 -duration=300s | tee results_usecase_3000rps_1.1.bin | vegeta report

vegeta report -type=json results_usecase_3000rps_1.1.bin > metrics_usecase_3000rps_1.1.json
cat results_usecase_3000rps_1.1.bin | vegeta plot > plot_usecase_3000rps_1.1.html
cat results_usecase_3000rps_1.1.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

/*
* USE CASE 2 POPULATE BULK DOCS
  */

echo "POST http://localhost:8080/es_load/bulk_item?batch_size=50" | vegeta attack -rate=1000 -max-workers=1000 -duration=300s | tee results_usecase_1000rps_2.bin | vegeta report

vegeta report -type=json results_usecase_1000rps_2.bin > metrics_usecase_1000rps_2.json
cat results_usecase_1000rps_2.bin | vegeta plot > plot_usecase_1000rps_2.html
cat results_usecase_1000rps_2.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/bulk_item?batch_size=50" | vegeta attack -rate=2000 -max-workers=1000 -duration=300s | tee results_usecase_2000rps_2.bin | vegeta report

vegeta report -type=json results_usecase_2000rps_2.bin > metrics_usecase_2000rps_2.json
cat results_usecase_2000rps_2.bin | vegeta plot > plot_usecase_2000rps_2.html
cat results_usecase_2000rps_2.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/bulk_item?batch_size=50" | vegeta attack -rate=3000 -max-workers=1000 -duration=300s | tee results_usecase_3000rps_2.bin | vegeta report

vegeta report -type=json results_usecase_3000rps_2.bin > metrics_usecase_3000rps_2.json
cat results_usecase_3000rps_2.bin | vegeta plot > plot_usecase_3000rps_2.html
cat results_usecase_3000rps_2.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

/*
* USE CASE 2.1 POPULATE BULK DOCS WITH BIG SIZE
  */

echo "POST http://localhost:8080/es_load/bulk_item?batch_size=500" | vegeta attack -rate=1000 -max-workers=1000 -duration=300s | tee results_usecase_1000rps_2.1.bin | vegeta report

vegeta report -type=json results_usecase_1000rps_2.1.bin > metrics_usecase_1000rps_2.1.json
cat results_usecase_1000rps_2.1.bin | vegeta plot > plot_usecase_1000rps_2.1.html
cat results_usecase_1000rps_2.1.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/bulk_item?batch_size=500" | vegeta attack -rate=2000 -max-workers=1000 -duration=300s | tee results_usecase_2000rps_2.1.bin | vegeta report

vegeta report -type=json results_usecase_2000rps_2.1.bin > metrics_usecase_2000rps_2.1.json
cat results_usecase_2000rps_2.1.bin | vegeta plot > plot_usecase_2000rps_2.1.html
cat results_usecase_2000rps_2.1.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/bulk_item?batch_size=500" | vegeta attack -rate=3000 -max-workers=1000 -duration=300s | tee results_usecase_3000rps_2.1.bin | vegeta report

vegeta report -type=json results_usecase_3000rps_2.1.bin > metrics_usecase_3000rps_2.1.json
cat results_usecase_3000rps_2.1.bin | vegeta plot > plot_usecase_3000rps_2.1.html
cat results_usecase_3000rps_2.1.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

/*
* USE CASE 2.2 POPULATE BULK DOCS WITH REFRESH
  */

echo "POST http://localhost:8080/es_load/bulk_item?refresh=true" | vegeta attack -rate=1000 -max-workers=1000 -duration=300s | tee results_usecase_1000rps_2.2.bin | vegeta report

vegeta report -type=json results_usecase_1000rps_2.2.bin > metrics_usecase_1000rps_2.2.json
cat results_usecase_1000rps_2.2.bin | vegeta plot > plot_usecase_1000rps_2.2.html
cat results_usecase_1000rps_2.2.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/bulk_item?refresh=true" | vegeta attack -rate=2000 -max-workers=1000 -duration=300s | tee results_usecase_2000rps_2.2.bin | vegeta report

vegeta report -type=json results_usecase_2000rps_2.2.bin > metrics_usecase_2000rps_2.2.json
cat results_usecase_2000rps_2.2.bin | vegeta plot > plot_usecase_2000rps_2.2.html
cat results_usecase_2000rps_2.2.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/bulk_item?refresh=true" | vegeta attack -rate=3000 -max-workers=1000 -duration=300s | tee results_usecase_3000rps_2.2.bin | vegeta report

vegeta report -type=json results_usecase_3000rps_2.2.bin > metrics_usecase_3000rps_2.2.json
cat results_usecase_3000rps_2.2.bin | vegeta plot > plot_usecase_3000rps_2.2.html
cat results_usecase_3000rps_2.2.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

/*
* USE CASE 2.3 POPULATE BULK DOCS WITH FLUSH
  */

echo "POST http://localhost:8080/es_load/bulk_item?flush=true" | vegeta attack -rate=1000 -max-workers=1000 -duration=300s | tee results_usecase_1000rps_2.3.bin | vegeta report

vegeta report -type=json results_usecase_1000rps_2.3.bin > metrics_usecase_1000rps_2.3.json
cat results_usecase_1000rps_2.3.bin | vegeta plot > plot_usecase_1000rps_2.3.html
cat results_usecase_1000rps_2.3.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/bulk_item?flush=true" | vegeta attack -rate=2000 -max-workers=1000 -duration=300s | tee results_usecase_2000rps_2.3.bin | vegeta report

vegeta report -type=json results_usecase_2000rps_2.3.bin > metrics_usecase_2000rps_2.3.json
cat results_usecase_2000rps_2.3.bin | vegeta plot > plot_usecase_2000rps_2.3.html
cat results_usecase_2000rps_2.3.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "POST http://localhost:8080/es_load/bulk_item?flush=true" | vegeta attack -rate=3000 -max-workers=1000 -duration=300s | tee results_usecase_3000rps_2.3.bin | vegeta report

vegeta report -type=json results_usecase_3000rps_2.3.bin > metrics_usecase_3000rps_2.3.json
cat results_usecase_3000rps_2.3.bin | vegeta plot > plot_usecase_3000rps_2.3.html
cat results_usecase_3000rps_2.3.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

/*
* USE CASE 3 SIMPLE SEARCHING
  */

echo "GET http://localhost:8080/es_load/item" | vegeta attack -rate=1000 -max-workers=1000 -duration=300s | tee results_usecase_1000rps_3.bin | vegeta report

vegeta report -type=json results_usecase_1000rps_3.bin > metrics_usecase_1000rps_3.json
cat results_usecase_1000rps_3.bin | vegeta plot > plot_usecase_1000rps_3.html
cat results_usecase_1000rps_3.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "GET http://localhost:8080/es_load/item" | vegeta attack -rate=2000 -max-workers=1000 -duration=300s | tee results_usecase_2000rps_3.bin | vegeta report

vegeta report -type=json results_usecase_2000rps_3.bin > metrics_usecase_2000rps_3.json
cat results_usecase_2000rps_3.bin | vegeta plot > plot_usecase_2000rps_3.html
cat results_usecase_2000rps_3.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "GET http://localhost:8080/es_load/item" | vegeta attack -rate=3000 -max-workers=1000 -duration=300s | tee results_usecase_3000rps_3.bin | vegeta report

vegeta report -type=json results_usecase_3000rps_3.bin > metrics_usecase_3000rps_3.json
cat results_usecase_3000rps_3.bin | vegeta plot > plot_usecase_3000rps_3.html
cat results_usecase_3000rps_3.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

/*
* USE CASE 4 SEARCHING WITH AGGREGATIONS
  */

Field_name = brand_name | categories_name | tags

echo "GET http://localhost:8080/es_load/item?agg=brand_name" | vegeta attack -rate=1000 -max-workers=1000 -duration=300s | tee results_usecase_1000rps_4.bin | vegeta report

vegeta report -type=json results_usecase_1000rps_4.bin > metrics_usecase_1000rps_4.json
cat results_usecase_1000rps_4.bin | vegeta plot > plot_usecase_1000rps_4.html
cat results_usecase_1000rps_4.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "GET http://localhost:8080/es_load/item?agg=brand_name" | vegeta attack -rate=2000 -max-workers=1000 -duration=300s | tee results_usecase_2000rps_4.bin | vegeta report

vegeta report -type=json results_usecase_2000rps_4.bin > metrics_usecase_2000rps_4.json
cat results_usecase_2000rps_4.bin | vegeta plot > plot_usecase_2000rps_4.html
cat results_usecase_2000rps_4.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"

==================================================

echo "GET http://localhost:8080/es_load/item?agg=brand_name" | vegeta attack -rate=3000 -max-workers=1000 -duration=300s | tee results_usecase_3000rps_4.bin | vegeta report

vegeta report -type=json results_usecase_3000rps_4.bin > metrics_usecase_3000rps_4.json
cat results_usecase_3000rps_4.bin | vegeta plot > plot_usecase_3000rps_4.html
cat results_usecase_3000rps_4.bin | vegeta report -type="hist[0,50ms,100ms,150ms,300ms,600ms,1200ms,2400ms,4800ms,9600ms]"
