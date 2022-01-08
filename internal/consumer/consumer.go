/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package consumer

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/superhero-match/consumer-report-user/internal/config"
)

// Consumer interface defines consumer methods.
type Consumer interface {
	FetchMessage(ctx context.Context) (kafka.Message, error)
	Close() error
	CommitMessages(ctx context.Context, m kafka.Message) error
}

// consumer holds Kafka consumer related data.
type consumer struct {
	Consumer *kafka.Reader
}

// NewConsumer configures Kafka consumer that consumes from configured topic.
func NewConsumer(cfg *config.Config) (Consumer, error) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:       cfg.Consumer.Brokers,
		Topic:         cfg.Consumer.Topic,
		GroupID:       cfg.Consumer.GroupID,
		QueueCapacity: int(1),
		MaxWait:       time.Second,
	})

	return &consumer{
		Consumer: r,
	}, nil
}
