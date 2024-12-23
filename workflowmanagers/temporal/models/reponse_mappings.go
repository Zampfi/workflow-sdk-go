package models

import (
	"github.com/Zampfi/workflow/sdk/workflowmanagers/temporal/constants"
	"github.com/jinzhu/copier"
	historypb "go.temporal.io/api/history/v1"
	v17 "go.temporal.io/api/workflow/v1"
	"go.temporal.io/sdk/client"
)

func (ws *WorkflowExecutionDetails) FromDescribeWorkflowResponse(workflowExecutionInfo *v17.WorkflowExecutionInfo) {
	ws.WorkflowID = workflowExecutionInfo.Execution.WorkflowId
	ws.RunID = workflowExecutionInfo.Execution.RunId
	ws.WorkflowType = workflowExecutionInfo.Type.Name
	ws.TaskQueue = workflowExecutionInfo.TaskQueue
	ws.HistoryLength = workflowExecutionInfo.HistoryLength
	ws.SearchAttributes = workflowExecutionInfo.SearchAttributes
	ws.StartTime = int64(workflowExecutionInfo.StartTime.GetSeconds())
	ws.CloseTime = int64(workflowExecutionInfo.CloseTime.GetSeconds())
	ws.ExecutionTime = int64(workflowExecutionInfo.ExecutionTime.GetSeconds())
	ws.Status = getWorkflowStatus(workflowExecutionInfo)
}

func getWorkflowStatus(workflowExecutionInfo *v17.WorkflowExecutionInfo) string {
	var status constants.WorkflowExecutionStatus
	return status.FromTemporalWorkflowStatus(workflowExecutionInfo.Status)
}

func (e *Event) FromHistoryEvent(historyEvent *historypb.HistoryEvent) {
	copier.Copy(e, historyEvent)
}

func (q *QueryWorkflowResponse) FromQueryWorkflowWithOptionsResponse(response *client.QueryWorkflowWithOptionsResponse) {
	copier.Copy(q, response)
}
