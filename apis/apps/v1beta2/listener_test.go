package v1beta2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wonderful-ya/emqx-operator116/apis/apps/v1beta2"
	corev1 "k8s.io/api/core/v1"
)

func TestListenerDefaults(t *testing.T) {
	listener := &v1beta2.Listener{
		Ports: v1beta2.Ports{
			MQTTS: 8885,
		},
	}

	listener.Default()

	assert.Equal(t, listener.Type, corev1.ServiceType("ClusterIP"))
	assert.Equal(t, listener.Ports.MQTTS, int32(8885))
	assert.Equal(t, listener.Ports.API, int32(8081))
}
