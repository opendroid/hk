package firestore_client

import (
	"cloud.google.com/go/firestore"
	"go.uber.org/zap"
)

// FirestoreClient defines methods to update Firestore data
type FirestoreClient interface {
	SaveAuthenticatedUser(sid string)
	Close()
}

// firestoreClient
type firestoreClient struct {
	client *firestore.Client
	logger *zap.Logger
}

// SaveAuthenticatedUser logs an authenticated user in the
func (fc *firestoreClient) SaveAuthenticatedUser(_ string) {
	if fc == nil {
		return
	}
}

// Close a closer interface for
func (fc *firestoreClient) Close() {
	if fc != nil {
		_ = fc.client.Close()
	}
}
