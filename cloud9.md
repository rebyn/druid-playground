git clone https://github.com/rebyn/druid-playground.git
cd druid-playground
sudo apt install kafkacat docker-compose jq
go run ./generate_events.go
docker-compose up -d