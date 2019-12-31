package view_test

import (
	"testing"

	"github.com/derailed/k9s/internal/client"
	"github.com/derailed/k9s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestRbacNew(t *testing.T) {
	v := view.NewRbac(client.GVR("rbac"))

	assert.Nil(t, v.Init(makeCtx()))
	assert.Equal(t, "Rbac", v.Name())
	assert.Equal(t, 3, len(v.Hints()))
}
