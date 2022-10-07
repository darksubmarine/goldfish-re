package goldfish_re

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_term_token(t *testing.T) {

	t1 := newDiscreteTerm(10)
	assert.EqualValues(t, "10", t1.token())

	// TODO the token for string value should be base64 encoded?
	t2 := newDiscreteTerm("some string")
	assert.EqualValues(t, "some string", t2.token())

	t3 := newStringVarTermWithValue("User", "plan", "some string")
	assert.EqualValues(t, "User.plan", t3.token())

}
