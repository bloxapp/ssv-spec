package valcheckduty

import (
	"encoding/json"
	"github.com/bloxapp/ssv-spec/ssv/spectest/tests/valcheck"
	"github.com/bloxapp/ssv-spec/types"
	"github.com/bloxapp/ssv-spec/types/testingutils"
)

// FarFutureDutySlot tests duty.Slot higher than expected
func FarFutureDutySlot() *valcheck.MultiSpecTest {
	consensusDataBytsF := func(cd *types.ConsensusData) []byte {
		cdCopy := &types.ConsensusData{}
		b, _ := json.Marshal(cd)
		if err := json.Unmarshal(b, cdCopy); err != nil {
			panic(err.Error())
		}
		cdCopy.Duty.Slot = 1000

		ret, _ := cdCopy.Encode()
		return ret
	}

	return &valcheck.MultiSpecTest{
		Name: "far future duty slot",
		Tests: []*valcheck.SpecTest{
			{
				Name:          "sync committee aggregator",
				Network:       types.NowTestNetwork,
				BeaconRole:    types.BNRoleSyncCommitteeContribution,
				Input:         consensusDataBytsF(testingutils.TestSyncCommitteeContributionConsensusData),
				ExpectedError: "duty invalid: duty epoch is into far future",
			},
			{
				Name:          "sync committee",
				Network:       types.NowTestNetwork,
				BeaconRole:    types.BNRoleSyncCommittee,
				Input:         consensusDataBytsF(testingutils.TestSyncCommitteeConsensusData),
				ExpectedError: "duty invalid: duty epoch is into far future",
			},
			{
				Name:          "aggregator",
				Network:       types.NowTestNetwork,
				BeaconRole:    types.BNRoleAggregator,
				Input:         consensusDataBytsF(testingutils.TestAggregatorConsensusData),
				ExpectedError: "duty invalid: duty epoch is into far future",
			},
			{
				Name:          "proposer",
				Network:       types.NowTestNetwork,
				BeaconRole:    types.BNRoleProposer,
				Input:         consensusDataBytsF(testingutils.TestProposerConsensusData),
				ExpectedError: "duty invalid: duty epoch is into far future",
			},
			{
				Name:          "attester",
				Network:       types.NowTestNetwork,
				BeaconRole:    types.BNRoleAttester,
				Input:         consensusDataBytsF(testingutils.TestAttesterConsensusData),
				ExpectedError: "duty invalid: duty epoch is into far future",
			},
		},
	}
}
