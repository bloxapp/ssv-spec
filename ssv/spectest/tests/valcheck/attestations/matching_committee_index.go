package attestations

import (
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// DutyCommitteeIndexNotMatchingAttestations tests that a duty committee index == attestation committee index
func DutyCommitteeIndexNotMatchingAttestations() *tests.SpecTest {
	ks := testingutils.Testing4SharesSet()
	dr := testingutils.AttesterRunner(ks)

	consensusData := &types.ConsensusData{
		Duty: &types.Duty{
			Type:                    types.BNRoleAttester,
			PubKey:                  testingutils.TestingValidatorPubKey,
			Slot:                    12,
			ValidatorIndex:          1,
			CommitteeIndex:          5,
			CommitteesAtSlot:        36,
			CommitteeLength:         128,
			ValidatorCommitteeIndex: 11,
		},
		AttestationData: testingutils.TestingAttestationData,
	}
	startingValue, _ := consensusData.Encode()

	// the starting value is not the same as the actual proposal!
	if err := dr.Decide(testingutils.TestAttesterConsensusData); err != nil {
		panic(err.Error())
	}

	msgs := []*types.SSVMessage{
		testingutils.SSVMsgAttester(testingutils.SignQBFTMsg(ks.Shares[1], 1, &qbft.Message{
			MsgType:    qbft.ProposalMsgType,
			Height:     qbft.FirstHeight,
			Round:      qbft.FirstRound,
			Identifier: []byte{1, 2, 3, 4},
			Data:       testingutils.ProposalDataBytes(startingValue, nil, nil),
		}), nil),
	}

	return &tests.SpecTest{
		Name:                    "duty committee index matches attestation committee index",
		Runner:                  dr,
		Messages:                msgs,
		PostDutyRunnerStateRoot: "c4eb0bb42cc382e468b2362e9d9cc622f388eef6a266901535bb1dfcc51e8868",
		ExpectedError:           "failed to process valcheck msg: could not process msg: proposal invalid: proposal not justified: proposal value invalid: attestation data CommitteeIndex != duty CommitteeIndex",
	}
}
