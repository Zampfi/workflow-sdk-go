package service

import (
	"context"
	"fmt"
	"testing"

	clientmock "github.com/Zampfi/citadel/mocks/workflows/temporal/client"
	"github.com/Zampfi/citadel/workflows/temporal/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type temporalServiceTestSuite struct {
	suite.Suite
	*require.Assertions
	client          *clientmock.MockClient
	temporalService TemporalService
}

func TestTemporalServiceTestSuite(t *testing.T) {
	suite.Run(t, new(temporalServiceTestSuite))
}

func (s *temporalServiceTestSuite) SetupTest() {
	t := s.T()
	s.Assertions = require.New(t)

	s.client = clientmock.NewMockClient(t)
	s.temporalService = NewTemporalService()
	s.temporalService.(*temporalService).setClient(s.client)
}

func (s *temporalServiceTestSuite) TearDownTest() {
	s.client.AssertExpectations(s.T())
}

func (s *temporalServiceTestSuite) TestTemporalService_Connect() {
	ctx := context.Background()
	s.client.EXPECT().Connect(ctx, models.ConnectClientParams{}).Return(nil)
	err := s.temporalService.Connect(ctx, models.ConnectClientParams{})

	s.NoError(err)
}

func SampleTemporal(ctx context.Context, arg1 string) error {
	fmt.Printf("Running temporal with arg1: %s\n", arg1)
	return nil
}

func (s *temporalServiceTestSuite) TestTemporalService_StartAsyncTemporal() {
	ctx := context.Background()
	params := models.ExecuteWorkflowParams{
		Workflow: "SampleTemporal",
		Args:     []interface{}{"arg1"},
		Options: models.StartWorkflowOptions{
			TaskQueue: "taskQueue",
		},
	}

	s.client.EXPECT().ExecuteAsyncWorkflow(ctx, params).Return(models.WorkflowResponse{
		RunID:      "runID",
		WorkflowID: "temporalID",
		FirstRunID: "firstRunID",
	}, nil)
	resp, err := s.temporalService.ExecuteAsyncWorkflow(ctx, params)

	s.NoError(err)
	s.Equal(resp.RunID, "runID")
	s.Equal(resp.WorkflowID, "temporalID")
	s.Equal(resp.FirstRunID, "firstRunID")
}

func (s *temporalServiceTestSuite) TestTemporalService_StartSyncTemporal() {
	ctx := context.Background()
	params := models.ExecuteWorkflowParams{
		Workflow: "SampleTemporal",
		Args:     []interface{}{"arg1"},
		Options: models.StartWorkflowOptions{
			TaskQueue: "taskQueue",
		},
	}

	s.client.EXPECT().ExecuteSyncWorkflow(ctx, params).Return(models.WorkflowResponse{
		RunID:      "runID",
		WorkflowID: "temporalID",
		FirstRunID: "firstRunID",
	}, nil)
	resp, err := s.temporalService.ExecuteSyncWorkflow(ctx, params)

	s.NoError(err)
	s.Equal(resp.RunID, "runID")
	s.Equal(resp.WorkflowID, "temporalID")
	s.Equal(resp.FirstRunID, "firstRunID")
}
