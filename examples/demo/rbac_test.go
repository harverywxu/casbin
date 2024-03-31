package demo

import (
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	args        []interface{}
	effectAllow bool
}

func TestRBAC(t *testing.T) {
	cases := []TestCase{
		//{
		//	args:        []interface{}{"superAdmin", "project", "read"},
		//	effectAllow: true,
		//},
		//{
		//	args:        []interface{}{"superAdmin", "project", "read"},
		//	effectAllow: true,
		//},
		//{
		//	args:        []interface{}{"quyuan", "project", "read"},
		//	effectAllow: true,
		//},
		{
			args:        []interface{}{"quyuan", "project", "write"},
			effectAllow: true,
		},
	}

	e, _ := casbin.NewEnforcer("rbac.conf", "rbac.csv", true)

	t.Logf("RBAC test start") // output for debug

	for _, c := range cases[:] {
		ok, err := e.Enforce(c.args...)
		assert.NoError(t, err)
		assert.Equal(t, c.effectAllow, ok, c)
	}
}
