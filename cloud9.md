git clone https://github.com/rebyn/druid-playground.git
cd druid-playground
sudo apt install kafkacat docker-compose jq zip gzip tar
go run ./generate_events.go
docker-compose up -d
kafkacat -z snappy -b 127.0.0.1:9093 -t events -P -l ./generated/events.json

curl http://169.254.169.254/latest/meta-data/instance-id
aws ec2 describe-instance-attribute --instance-id i-02448ec8312be7c09 --attribute groupSet
aws ec2 authorize-security-group-ingress --group-id sg-05f4672b4be2a991c --protocol tcp --port 8888 --cidr 27.71.207.174/22
curl icanhazip.com
54.169.203.32

pwd
zip -r events.zip /home/ubuntu/environment/druid-playground/generated

aws s3 cp /home/ubuntu/environment/druid-playground/generated/events.json s3://tbox-rds-mssql-backup/social-commerce-platform/druid/events.json

aws s3 cp s3://tbox-rds-mssql-backup/social-commerce-platform/druid/events.json /home/ubuntu/environment/druid-playground/generated/events.json



CLOUD9 UBUNTU AMI: ami-03ff67a9197933525

Druid: iso: ISO8601 with 'T' separator, like       "2000-01-01T01:02:03.456"
{"CreatorId":1,"UserId":94,"PostId":77,"Timestamp":"2020-10-02T22:05:00+07:00"}