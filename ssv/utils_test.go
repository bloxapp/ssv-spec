package ssv_test

import (
	"github.com/bloxapp/ssv-spec/ssv"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

var testConsensusData = &types.ConsensusData{
	Duty:            testingutils.TestingAttesterDuty,
	AttestationData: testingutils.TestingAttestationData,
}
var TestConsensusDataByts, _ = testConsensusData.Encode()

func NewTestingDutyExecutionState() *ssv.State {
	return ssv.NewRunnerState(3, testingutils.TestingAttesterDuty)
}
