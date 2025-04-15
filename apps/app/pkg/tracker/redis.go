package tracker

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
)

type RedisStorage struct {
	rdb *redis.Client
}

func NewRedisStorage(rdb *redis.Client) Storage {
	return &RedisStorage{rdb: rdb}
}

func (s *RedisStorage) SetLocation(ctx context.Context, l LocationDTO) error {
	userID := l.UserID
	record := Location{
		Lat: l.Lat,
		Lng: l.Lng,
	}

	payload, err := proto.Marshal(&record)
	if err != nil {
		return fmt.Errorf("failed to serialize location: %w", err)
	}

	err = s.rdb.Set(ctx, userID, payload, 0).Err()
	if err != nil {
		return fmt.Errorf("redis SET failed: %w", err)
	}

	return nil
}

func (s *RedisStorage) GetLocation(ctx context.Context, userID string) (LocationDTO, error) {
	var res LocationDTO
	payload, err := s.rdb.Get(ctx, userID).Result()
	if err != nil {
		if err == redis.Nil {
			return res, ErrNotFound
		}
		return res, fmt.Errorf("redis GET failed: %w", err)
	}

	var val Location
	if err := proto.Unmarshal([]byte(payload), &val); err != nil {
		return res, fmt.Errorf("failed to deserialize location: %w", err)
	}

	res.UserID = userID
	res.Lat = val.Lat
	res.Lng = val.Lng

	return res, nil
}
