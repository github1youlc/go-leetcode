package solve

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test(t *testing.T) {
	assert.EqualValues(t, -1, compareVersion("0.1", "1.1"))
	assert.EqualValues(t, 1, compareVersion("1.0.1", "1"))
	assert.EqualValues(t, -1, compareVersion("7.5.2.4", "7.5.3"))
	assert.EqualValues(t, 0, compareVersion("1.01", "1.001"))
	assert.EqualValues(t, 0, compareVersion("1.0", "1.0.0"))
	assert.EqualValues(t, -1, compareVersion("1.2", "1.10"))
}