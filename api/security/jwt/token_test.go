package securityJWT

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/techschool/simplebank/uti"
)

var ranString = uti.RandomString(40)

func GetNewMarker(t *testing.T) Maker {
	d, err := time.ParseDuration("15m")
	require.NoError(t, err)
	require.NotEmpty(t, d)
	marker, err := NewJWTMaker(ranString, OptionJwt{
		TimeDuration: d,
	})
	require.NoError(t, err)
	require.NotEmpty(t, marker)
	return marker
}

func TestNewMarker(t *testing.T) {
	GetNewMarker(t)
}

func CreateToken(t *testing.T) string {
	marker := GetNewMarker(t)
	access_token, error := marker.BuildToken(uti.RandomOwner())
	require.NoError(t, error)
	require.NotEmpty(t, access_token)
	return access_token
}

func TestVerifyToken(t *testing.T) {
	marker := GetNewMarker(t)
	access_token := CreateToken(t)
	print(access_token)
	payload, err := marker.VerifyToken(access_token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)
}
