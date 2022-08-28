package firestore_client

import (
	"cloud.google.com/go/firestore"
	"context"
	"go.uber.org/zap"
)

var _ = NewFirestoreClient // make compiler happy

// NewFirestoreClient create a new FirestoreClient as interface
func NewFirestoreClient(projectID string, logger *zap.Logger) FirestoreClient {
	client, err := firestore.NewClient(context.Background(), projectID)
	if err != nil {
		logger.Error("NewFirestoreClient error", zap.Error(err))
		return nil
	}
	fs := firestoreClient{
		client: client,
		logger: logger,
	}
	return &fs
}
