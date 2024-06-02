package operator

import (
	reflect2 "reflect"
	"testing"

	comparable2 "github.com/ssvlabs/ssv-spec/types/testingutils/comparable"

	"github.com/ssvlabs/ssv-spec/types"
	"github.com/stretchr/testify/require"
)

type OperatorTest struct {
	Name                  string
	Operator              types.Operator
	Message               types.SignedSSVMessage
	ExpectedHasQuorum     bool
	ExpectedFullCommittee bool
	ExpectedError         string
}

func (test *OperatorTest) TestName() string {
	return "operator " + test.Name
}

// Returns the number of unique signers in the message signers list
func (test *OperatorTest) GetUniqueMessageSignersCount() int {
	uniqueSigners := make(map[uint64]bool)

	for _, element := range test.Message.GetOperatorIDs() {
		uniqueSigners[element] = true
	}

	return len(uniqueSigners)
}

func (test *OperatorTest) Run(t *testing.T) {

	// Validate message
	err := test.Message.Validate()
	if len(test.ExpectedError) != 0 {
		require.EqualError(t, err, test.ExpectedError)
	} else {
		require.NoError(t, err)
	}

	// Get unique signers
	numSigners := test.GetUniqueMessageSignersCount()

	// Test expected thresholds results
	require.Equal(t, test.ExpectedHasQuorum, test.Operator.HasQuorum(numSigners))
	require.Equal(t, test.ExpectedFullCommittee, (len(test.Operator.Committee) == numSigners))

	comparable2.CompareWithJson(t, test, test.TestName(), reflect2.TypeOf(test).String())
}
