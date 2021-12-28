package main
// should return 841526
import (
	"testing"
	"github.com/iloveicedgreentea/advent_of_code/common"
	"github.com/stretchr/testify/assert"
)

var (
	gam = [12]int{483, 487, 485, 467, 513, 509, 492, 532, 508, 497, 490, 506}
	eps = [12]int{517, 513, 515, 533, 487, 491, 508, 468, 492, 503, 510, 494}
	lines = common.ReadFile("input.txt")
	length = len(lines)
)

func TestCalcReturn(t *testing.T)  {
	gamma, epsilon := calcInputs(lines)

	assert.Equal(t, gamma, gam)
	assert.Equal(t, epsilon, eps)
}

func TestProcessNum(t *testing.T) {
	gamma := binToDec(processNum(gam, length))
	epsilon := binToDec(processNum(eps, length))

	assert.Equal(t, gamma, int64(217))
	assert.Equal(t, epsilon, int64(3878))
}
