1.go mod init example.com/m

2.go get github.com/segmentio/kafka-go

3.in kafka
###
kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic your-topic
