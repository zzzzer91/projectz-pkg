package alarmy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetHoroscope(t *testing.T) {
	s := NewService()
	resp, err := s.GetHoroscope(context.Background(), "Virgo", "%2B08:00")
	assert.NoError(t, err)
	t.Logf("%#v\n", resp)
}
