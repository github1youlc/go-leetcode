package p0169_Majority_Element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	assert.EqualValues(t, 2, majorityElement([]int{1, 4, 2, 2, 2}))
}
