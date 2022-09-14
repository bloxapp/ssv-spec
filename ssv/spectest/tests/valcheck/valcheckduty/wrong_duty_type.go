package valcheckduty

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// WrongDutyType tests duty.Type not attester
func WrongDutyType() *valcheck.MultiSpecTest {
	consensusDataBytsF := func(role types.BeaconRole) []byte {
		data := &types.ConsensusData{
			Duty: &types.Duty{
				Type:                    role,
				PubKey:                  testingutils.TestingValidatorPubKey,
				Slot:                    testingutils.TestingDutySlot,
				ValidatorIndex:          testingutils.TestingValidatorIndex,
				CommitteeIndex:          3,
				CommitteesAtSlot:        36,
				CommitteeLength:         128,
				ValidatorCommitteeIndex: 11,
			},
			AttestationData: &phase0.AttestationData{
				Slot:            testingutils.TestingDutySlot,
				Index:           3,
				BeaconBlockRoot: phase0.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
				Source: &phase0.Checkpoint{
					Epoch: 0,
					Root:  phase0.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
				},
				Target: &phase0.Checkpoint{
					Epoch: 1,
					Root:  phase0.Root{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2},
				},
			},
		}

		input, _ := data.Encode()
		return input
	}

	return &valcheck.MultiSpecTest{
		Name: "wrong duty type",
		Tests: []*valcheck.SpecTest{
			{
				Name:          "sync committee aggregator",
				Network:       types.NowTestNetwork,
				BeaconRole:    types.BNRoleSyncCommitteeContribution,
				Input:         consensusDataBytsF(types.BNRoleProposer),
				ExpectedError: "duty invalid: wrong beacon role type",
			},
			{
				Name:          "sync committee",
				Network:       types.NowTestNetwork,
				BeaconRole:    types.BNRoleSyncCommittee,
				Input:         consensusDataBytsF(types.BNRoleProposer),
				ExpectedError: "duty invalid: wrong beacon role type",
			},
			{
				Name:          "aggregator",
				Network:       types.NowTestNetwork,
				BeaconRole:    types.BNRoleAggregator,
				Input:         consensusDataBytsF(types.BNRoleProposer),
				ExpectedError: "duty invalid: wrong beacon role type",
			},
			{
				Name:          "proposer",
				Network:       types.NowTestNetwork,
				BeaconRole:    types.BNRoleProposer,
				Input:         consensusDataBytsF(types.BNRoleAttester),
				ExpectedError: "duty invalid: wrong beacon role type",
			},
			{
				Name:          "attester",
				Network:       types.NowTestNetwork,
				BeaconRole:    types.BNRoleAttester,
				Input:         consensusDataBytsF(types.BNRoleProposer),
				ExpectedError: "duty invalid: wrong beacon role type",
			},
		},
	}
}
