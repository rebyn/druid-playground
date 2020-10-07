git clone https://github.com/rebyn/druid-playground.git
cd druid-playground
sudo apt install kafkacat docker-compose jq
go run ./generate_events.go
docker-compose up -d

curl http://169.254.169.254/latest/meta-data/instance-id
aws ec2 describe-instance-attribute --instance-id i-0b682862a57a2b242 --attribute groupSet
aws ec2 authorize-security-group-ingress --group-id sg-0e5bd4ff3fdbe47a8 --protocol tcp --port 8888 --cidr 27.71.207.174/22
curl icanhazip.com