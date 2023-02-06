package integ

import (
	"github.com/indeedeng/iwf-golang-sdk/gen/iwfidl"
	"github.com/indeedeng/iwf-golang-sdk/iwf"
	"time"
)

type stateApiTimeoutWorkflowState1 struct {
	iwf.DefaultStateId
}

func (b stateApiTimeoutWorkflowState1) Start(ctx iwf.WorkflowContext, input iwf.Object, persistence iwf.Persistence, communication iwf.Communication) (*iwf.CommandRequest, error) {
	time.Sleep(time.Minute)
	return nil, nil
}

func (b stateApiTimeoutWorkflowState1) Decide(ctx iwf.WorkflowContext, input iwf.Object, commandResults iwf.CommandResults, persistence iwf.Persistence, communication iwf.Communication) (*iwf.StateDecision, error) {
	return iwf.ForceFailWorkflow("a failing message"), nil
}

func (b stateApiTimeoutWorkflowState1) GetStateOptions() *iwfidl.WorkflowStateOptions {
	return &iwfidl.WorkflowStateOptions{
		StartApiRetryPolicy: &iwfidl.RetryPolicy{
			MaximumAttempts: iwfidl.PtrInt32(1),
		},
		StartApiTimeoutSeconds: iwfidl.PtrInt32(1),
	}
}
