package temporal

import (
	"context"
	"fmt"
	"testing"

	clientmock "github.com/Zampfi/workflow-sdk-go/mocks/workflowmanagers/temporal/client"
	"github.com/Zampfi/workflow-sdk-go/workflowmanagers/temporal/models"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	temporalworkflow "go.temporal.io/sdk/workflow"
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

func SampleWorkflow(ctx temporalworkflow.Context, arg1 string) error {
	fmt.Printf("Running temporal with arg1: %s\n", arg1)
	return nil
}

func SampleActivity(ctx context.Context, arg1 string) (string, error) {
	fmt.Printf("Running activity with arg1: %s\n", arg1)
	return "activity", nil
}

func (s *temporalServiceTestSuite) TestTemporalService_QuerySchedule() {
	ctx := context.Background()
	params := models.QueryScheduleParams{
		ScheduleID: "ScheduleID",
	}

	s.client.EXPECT().QuerySchedule(ctx, params).Return(models.QueryScheduleResponse{}, nil)
	resp, err := s.temporalService.QuerySchedule(ctx, params)

	s.NoError(err)
	s.Equal(resp, models.QueryScheduleResponse{})
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
	}, nil)
	resp, err := s.temporalService.ExecuteAsyncWorkflow(ctx, params)

	s.NoError(err)
	s.Equal(resp.RunID, "runID")
	s.Equal(resp.WorkflowID, "temporalID")
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
	}, nil)
	resp, err := s.temporalService.ExecuteSyncWorkflow(ctx, params)

	s.NoError(err)
	s.Equal(resp.RunID, "runID")
	s.Equal(resp.WorkflowID, "temporalID")
}

func (s *temporalServiceTestSuite) TestTemporalService_ListWorkflows() {
	ctx := context.Background()
	params := models.ListWorkflowsParams{
		PageSize: 10,
		Query:    "(WorkflowID = 'wid1' or (WorkflowType = 'type2' and WorkflowID = 'wid2'))",
	}

	s.client.EXPECT().ListWorkflows(ctx, params).Return(models.ListWorkflowsResponse{}, nil)
	resp, err := s.temporalService.ListWorkflows(ctx, params)

	s.NoError(err)
	s.Equal(resp, models.ListWorkflowsResponse{})
}

func (s *temporalServiceTestSuite) TestTemporalService_GetWorkflowDetails() {
	ctx := context.Background()
	params := models.GetWorkflowDetailsParams{
		WorkflowID: "wid",
		RunID:      "rid",
	}

	s.client.EXPECT().GetWorkflowDetails(ctx, params).Return(models.WorkflowDetailsResponse{}, nil)
	resp, err := s.temporalService.GetWorkflowDetails(ctx, params)

	s.NoError(err)
	s.Equal(resp, models.WorkflowDetailsResponse{})
}

func (s *temporalServiceTestSuite) TestTemporalService_TerminateWorkflow() {
	ctx := context.Background()
	params := models.TerminateWorkflowParams{
		WorkflowID: "wid",
		RunID:      "rid",
		Reason:     "reason",
		Details:    []interface{}{"details"},
	}

	s.client.EXPECT().TerminateWorkflow(ctx, params).Return(nil)
	err := s.temporalService.TerminateWorkflow(ctx, params)
	s.NoError(err)
}

func (s *temporalServiceTestSuite) TestTemporalService_CancelWorkflow() {
	ctx := context.Background()
	params := models.CancelWorkflowParams{
		WorkflowID: "wid",
		RunID:      "rid",
	}

	s.client.EXPECT().CancelWorkflow(ctx, params).Return(nil)
	err := s.temporalService.CancelWorkflow(ctx, params)
	s.NoError(err)
}

func (s *temporalServiceTestSuite) TestTemporalService_SignalWorkflow() {
	ctx := context.Background()
	params := models.SignalWorkflowParams{
		WorkflowID: "wid",
		RunID:      "rid",
		SignalName: "signal",
	}

	s.client.EXPECT().SignalWorkflow(ctx, params).Return(nil)
	err := s.temporalService.SignalWorkflow(ctx, params)
	s.NoError(err)
}

func (s *temporalServiceTestSuite) TestTemporalService_QueryWorkflow() {
	ctx := context.Background()
	params := models.QueryWorkflowParams{
		WorkflowID: "wid",
		RunID:      "rid",
	}

	s.client.EXPECT().QueryWorkflow(ctx, params).Return(models.QueryWorkflowResponse{}, nil)
	resp, err := s.temporalService.QueryWorkflow(ctx, params)

	s.NoError(err)
	s.Equal(resp, models.QueryWorkflowResponse{})
}

func (s *temporalServiceTestSuite) TestTemporalService_Close() {
	ctx := context.Background()
	s.client.EXPECT().Close(ctx)
	s.temporalService.Close(ctx)
}
