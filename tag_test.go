package goldfish_re

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_utils_parseTag(t *testing.T) {
	tag1 := "object=User"
	tag2 := "object=User,attribute=Plan"
	tag3 := "object=User,attribute=Plan,value=gold"
	tag4 := "object=User,attribute=Plan,value"

	obj, attr, val, err := parseTag(tag1)
	assert.EqualValues(t, "User", obj)
	assert.Empty(t, attr)
	assert.Empty(t, val)
	assert.Nil(t, err)

	obj, attr, val, err = parseTag(tag2)
	assert.EqualValues(t, "User", obj)
	assert.EqualValues(t, "Plan", attr)
	assert.Empty(t, val)
	assert.Nil(t, err)

	obj, attr, val, err = parseTag(tag3)
	assert.EqualValues(t, "User", obj)
	assert.EqualValues(t, "Plan", attr)
	assert.EqualValues(t, "gold", val)
	assert.Nil(t, err)

	obj, attr, val, err = parseTag(tag4)
	assert.Empty(t, val)
	assert.Empty(t, val)
	assert.Empty(t, val)
	assert.NotNil(t, err)
}
