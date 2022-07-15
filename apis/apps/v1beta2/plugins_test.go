package v1beta2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wonderful-ya/emqx-operator116/apis/apps/v1beta2"
)

func TestPluginsDefault(t *testing.T) {
	plugins := &v1beta2.Plugins{
		Items: []v1beta2.Plugin{
			{
				Name:   "foo",
				Enable: true,
			},
			{
				Name:   "bar",
				Enable: false,
			},
		},
	}

	plugins.Default()
	assert.ElementsMatch(t, plugins.Items,
		[]v1beta2.Plugin{
			{
				Name:   "foo",
				Enable: true,
			},
			{
				Name:   "bar",
				Enable: false,
			},
			{
				Name:   "emqx_management",
				Enable: true,
			},
		},
	)
}

func TestPluginsString(t *testing.T) {
	plugins := &v1beta2.Plugins{
		Items: []v1beta2.Plugin{
			{
				Name:   "foo",
				Enable: true,
			},
			{
				Name:   "bar",
				Enable: false,
			},
		},
	}

	assert.Equal(t,
		plugins.String(),
		"{foo, true}.\n{bar, false}.\n",
	)
}
