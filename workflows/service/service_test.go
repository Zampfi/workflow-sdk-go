package service

import (
	"context"
	"fmt"
	"testing"

	clientmock "github.com/Zampfi/citadel/mocks/workflows/client"
	"github.com/Zampfi/citadel/workflows/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type workflowServiceTestSuite struct {
	suite.Suite
	*require.Assertions
	client          *clientmock.MockClient
	workflowService WorkflowService
}

func TestWorkflowServiceTestSuite(t *testing.T) {
	suite.Run(t, new(workflowServiceTestSuite))
}

func (s *workflowServiceTestSuite) SetupTest() {
	t := s.T()
	s.Assertions = require.New(t)

	s.client = clientmock.NewMockClient(t)
	s.workflowService = NewWorkflowService()
	s.workflowService.(*workflowService).setClient(s.client)
}

func (s *workflowServiceTestSuite) TearDownTest() {
	s.client.AssertExpectations(s.T())
}

func (s *workflowServiceTestSuite) TestWorkflowService_Connect() {
	ctx := context.Background()
	s.client.EXPECT().Connect(ctx, models.ConnectClientParams{}).Return(nil)
	err := s.workflowService.Connect(ctx, models.ConnectClientParams{})

	s.NoError(err)
}

func SampleWorkflow(ctx context.Context, arg1 string) error {
	fmt.Printf("Running workflow with arg1: %s\n", arg1)
	return nil
}

func (s *workflowServiceTestSuite) TestWorkflowService_StartAsyncWorkflow() {
	ctx := context.Background()
	params := models.StartWorkflowParams{
		Workflow: "SampleWorkflow",
		Args:     []interface{}{"arg1"},
		Options: models.StartWorkflowOptions{
			TaskQueue: "taskQueue",
		},
	}

	s.client.EXPECT().StartAsyncWorkflow(ctx, params).Return(models.StartWorkflowResponse{
		RunID:      "runID",
		WorkflowID: "workflowID",
		FirstRunID: "firstRunID",
	}, nil)
	resp, err := s.workflowService.StartAsyncWorkflow(ctx, params)

	s.NoError(err)
	s.Equal(resp.RunID, "runID")
	s.Equal(resp.WorkflowID, "workflowID")
	s.Equal(resp.FirstRunID, "firstRunID")
}

func (s *workflowServiceTestSuite) TestWorkflowService_StartSyncWorkflow() {
	ctx := context.Background()
	params := models.StartWorkflowParams{
		Workflow: "SampleWorkflow",
		Args:     []interface{}{"arg1"},
		Options: models.StartWorkflowOptions{
			TaskQueue: "taskQueue",
		},
	}

	s.client.EXPECT().StartSyncWorkflow(ctx, params).Return(models.StartWorkflowResponse{
		RunID:      "runID",
		WorkflowID: "workflowID",
		FirstRunID: "firstRunID",
	}, nil)
	resp, err := s.workflowService.StartSyncWorkflow(ctx, params)

	s.NoError(err)
	s.Equal(resp.RunID, "runID")
	s.Equal(resp.WorkflowID, "workflowID")
	s.Equal(resp.FirstRunID, "firstRunID")
}
