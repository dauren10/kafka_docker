go mod init example.com/m

go get github.com/segmentio/kafka-go

in kafka
kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic your-topic
