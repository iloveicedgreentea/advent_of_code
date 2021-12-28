
package common
// should return 841526
import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestRemoveFromArrayStr(t *testing.T) {
	s := []string{"1", "2", "3"}
	s2 := []string{"1", "2"}
	newS := RemoveFromArrayStr(s, 2)
	assert.Equal(t, newS, s2, "Should remove 3rd element")
}
func TestRemoveFromArrayInt(t *testing.T) {
	s := []int{1,2,3}
	s2 := []int{1,2}
	newS := RemoveFromArrayInt(s, 2)
	assert.Equal(t, newS, s2, "Should remove 3rd element")
}