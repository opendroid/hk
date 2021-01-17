package server

import (
	"fmt"
	"github.com/opendroid/hk/export"
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"time"
)

// userData saves health data summary for the user.
type userData struct {
	health *export.HealthData
}

func healthData(user string) *export.HealthData {
	if u, ok := users[user]; !ok {
		return u.health
	}
	return nil
}

func saveHD(uID string) {
	// Read XML file
	start := time.Now()
	if health, err := export.Parse(file); err != nil {
		fmt.Printf("Error parsing %q, %v\n", file, err)
	} else {
		// health.Describe()
		// health.UnseenCheck()
		// health.DescribeRecordsTable(os.Stdout, " | ")
		users[uID] = userData{health}
	}
	end := time.Since(start)
	logger.Debug("Analysis done", zap.String("user", uID), zap.Int64("ms", end.Milliseconds()))
}
