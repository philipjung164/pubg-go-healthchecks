# PUBGMG Golang Healthchecks

This library houses healthchecks that can be reused between different applications for Kubernetes liveness and readiness checks.

## DynamoDB

```go
import "git.projectbro.com/isd/pubg-go-healthchecks"

health := healthcheck.NewHandler()

health.AddReadinessCheck("dynamodb-table",
	pubghealth.DynamoTableStatusCheck(db.client, db.TableName(), 150*time.Millisecond, 10*time.Second))
```

## Redis

```go
import "git.projectbro.com/isd/pubg-go-healthchecks"

health := healthcheck.NewHandler()

health.AddReadinessCheck("redis-ping",
	pubghealth.RedisPingCheck(r.client, 150*time.Millisecond, 10*time.Second))
```

## Kafka

```go
import "git.projectbro.com/isd/pubg-go-healthchecks"

health := healthcheck.NewHandler()

health.AddReadinessCheck("kafka-connected", 
	pubghealth.KafkaConnectionCheck(i.client, 250*time.Millisecond))
health.AddReadinessCheck("kafka-topics",
	pubghealth.KafkaTopicsExist(i.client, 250*time.Millisecond, strings.Split(i.cfg.KafkaTopic, ",")))
```
