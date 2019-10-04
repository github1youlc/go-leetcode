package p0171_Excel_Sheet_Column_Number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_titleToNumber(t *testing.T) {
	assert.EqualValues(t, 1, titleToNumber("A"))
	assert.EqualValues(t, 26, titleToNumber("Z"))
	assert.EqualValues(t, 28, titleToNumber("AB"))
	assert.EqualValues(t, 701, titleToNumber("ZY"))
}
