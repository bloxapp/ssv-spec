package attestationsynccommittee

import (
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// StartMaximumPossibleDuties starts a cluster runner with 500 attestation and 500 sync committee duties
func StartMaximumPossibleDuties() tests.SpecTest {

	ks := testingutils.Testing4SharesSet()

	multiSpecTest := &tests.MultiMsgProcessingSpecTest{
		Name: "start maximum possible duties",
		Tests: []*tests.MsgProcessingSpecTest{
			{
				Name:           "500 attestations 500 sync committees",
				Runner:         testingutils.CommitteeRunner(ks),
				Duty:           testingutils.TestingCommitteeDuty(testingutils.TestingDutySlot, ValidatorIndexList(500), ValidatorIndexList(500)),
				Messages:       []*types.SignedSSVMessage{},
				OutputMessages: []*types.PartialSignatureMessages{},
			},
		},
	}

	return multiSpecTest
}
