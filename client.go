package wasabi

import (
	"context"

	"github.com/inloco/go-wasabi/assignments"
	"github.com/inloco/go-wasabi/experiments"
)

// Client is a (mockable) interface for our Wasabi client.
// Each method is 1:1 to the routes exposed by the Wasabi application.
type Client interface {
	GenerateAssignment(ctx context.Context, experimentName string, userID string) (*assignments.Assignment, error)

	CreateExperiment(ctx context.Context, experiment *experiments.Experiment) (*experiments.Experiment, error)
	GetExperiments(ctx context.Context) ([]*experiments.Experiment, error)
	GetExperimentByID(ctx context.Context, experimentID string) (*assignments.Assignment, error)
}