package pubghealth

import (
	"fmt"
	"github.com/bsm/sarama-cluster"
	"github.com/heptiolabs/healthcheck"
	"github.com/pkg/errors"
	"time"
)

// KafkaConnectionCheck verifies the client's connectivity to the Kafka cluster.
func KafkaConnectionCheck(client *cluster.Client, timeout time.Duration, frequency time.Duration) healthcheck.Check {
	checkFunc := func() error {
		brokers := (*client).Brokers()
		if len(brokers) < 1 {
			return errors.New("Kafka client is not connected to any brokers.")
		}
		for _, b := range brokers {
			isConnected, err := b.Connected()
			if err != nil {
				// Error while trying to connect, wrap it.
				return errors.Wrap(err, fmt.Sprintf("Kafka client is unable to connect to the broker at '%s'", b.Addr()))
			}
			if !isConnected {
				// Not connected, no connection errors, report it.
				return errors.Errorf("Kafka client is not connected to the broker at '%s'", b.Addr())
			}
		}
		// S'all good!
		return nil
	}
	return healthcheck.Timeout(healthcheck.Async(checkFunc, frequency), timeout)
}

// KafkaTopicsExist verifies that specific topics exist in the Kafka cluster.
func KafkaTopicsExist(client *cluster.Client, timeout time.Duration, frequency time.Duration, expected []string) healthcheck.Check {
	checkFunc := func() error {
		// Force a refresh of topics:
		err := (*client).RefreshMetadata(expected...)
		if err != nil {
			return err
		}
		actual, err := (*client).Topics()
		if err != nil {
			return err
		}
		// Check that all expected topics are found in the list of actual topics.
		for _, eT := range expected {
			found := false
			for _, aT := range actual {
				if eT == aT {
					found = true
					break
				}
			}
			if !found {
				return errors.Errorf("Topic '%s' was not found in the Kafka cluster.", eT)
			}
		}
		// S'all good!
		return nil
	}
	return healthcheck.Timeout(healthcheck.Async(checkFunc, frequency), timeout)
}
