package timeoutduration

import (
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/bloxapp/ssv-spec/qbft"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests"
	"github.com/bloxapp/ssv-spec/qbft/spectest/tests/timeout"
	"github.com/bloxapp/ssv-spec/types"
)

// Round3Duration tests timeout duration for round 3 where the current time is the expected start of the round
func Round3Duration() *tests.MultiSpecTest {
	testingNetwork := types.HoleskyNetwork
	height := qbft.FirstHeight
	var round qbft.Round = 3
	dutyStartTime := testingNetwork.EstimatedTimeAtSlot(phase0.Slot(height))

	return &tests.MultiSpecTest{
		Name: "round 3 duration",
		Tests: []tests.SpecTest{
			&timeout.TimeoutDurationTest{
				Name:             "sync committee",
				Role:             types.BNRoleSyncCommittee,
				Height:           height,
				Round:            round,
				Network:          testingNetwork,
				CurrentTime:      dutyStartTime + 8,
				ExpectedDuration: 2,
			},
			&timeout.TimeoutDurationTest{
				Name:             "sync committee contribution",
				Role:             types.BNRoleSyncCommitteeContribution,
				Height:           height,
				Round:            round,
				Network:          testingNetwork,
				CurrentTime:      dutyStartTime + 12,
				ExpectedDuration: 2,
			},
			&timeout.TimeoutDurationTest{
				Name:             "attester",
				Role:             types.BNRoleAttester,
				Height:           height,
				Round:            round,
				Network:          testingNetwork,
				CurrentTime:      dutyStartTime + 8,
				ExpectedDuration: 2,
			},
			&timeout.TimeoutDurationTest{
				Name:             "aggregator",
				Role:             types.BNRoleAggregator,
				Height:           height,
				Round:            round,
				Network:          testingNetwork,
				CurrentTime:      dutyStartTime + 12,
				ExpectedDuration: 2,
			},
			&timeout.TimeoutDurationTest{
				Name:             "block proposer",
				Role:             types.BNRoleProposer,
				Height:           height,
				Round:            round,
				Network:          testingNetwork,
				CurrentTime:      dutyStartTime + 6,
				ExpectedDuration: 2,
			},
		},
	}

}
