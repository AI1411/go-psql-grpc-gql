package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/AI1411/go-psql_grpc_gql/config"
	"github.com/AI1411/go-psql_grpc_gql/internal/model"
)

const (
	basePrefix = "sessions:"
)

type SessionRepository interface {
	CreateSession(ctx context.Context, session *model.Session, expire int) (string, error)
	GetSessionByID(ctx context.Context, sessionID string) (*model.Session, error)
	DeleteByID(ctx context.Context, sessionID string) error
}

type sessionRepo struct {
	redisClient *redis.Client
	basePrefix  string
	cfg         *config.Config
}

func NewSessionRepository(redisClient *redis.Client, cfg *config.Config) SessionRepository {
	return &sessionRepo{redisClient: redisClient, basePrefix: basePrefix, cfg: cfg}
}

func (s *sessionRepo) CreateSession(ctx context.Context, sess *model.Session, expire int) (string, error) {
	sess.SessionID = uuid.New().String()
	sessionKey := s.createKey(sess.SessionID)

	sessBytes, err := json.Marshal(&sess)
	if err != nil {
		return "", errors.WithMessage(err, "sessionRepo.CreateSession.json.Marshal")
	}
	if err = s.redisClient.Set(ctx, sessionKey, sessBytes, time.Second*time.Duration(expire)).Err(); err != nil {
		return "", errors.Wrap(err, "sessionRepo.CreateSession.redisClient.Set")
	}
	return sess.SessionID, nil
}

func (s *sessionRepo) GetSessionByID(ctx context.Context, sessionID string) (*model.Session, error) {
	sessBytes, err := s.redisClient.Get(ctx, s.createKey(sessionID)).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "sessionRep.GetSessionByID.redisClient.Get")
	}

	sess := &model.Session{}
	if err = json.Unmarshal(sessBytes, &sess); err != nil {
		return nil, errors.Wrap(err, "sessionRepo.GetSessionByID.json.Unmarshal")
	}
	return sess, nil
}

func (s *sessionRepo) DeleteByID(ctx context.Context, sessionID string) error {
	if err := s.redisClient.Del(ctx, s.createKey(sessionID)).Err(); err != nil {
		return errors.Wrap(err, "sessionRepo.DeleteByID")
	}
	return nil
}

func (s *sessionRepo) createKey(sessionID string) string {
	return fmt.Sprintf("%s: %s", s.basePrefix, sessionID)
}
