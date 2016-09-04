package footballdata

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Client_New(t *testing.T) {
	client := NewClient(http.DefaultClient)
	_, err := client.Competitions().Do()
	require.Nil(t, err)
}

func Test_Client_GoNew(t *testing.T) {
	client := new(Client)
	_, err := client.Competitions().Do()
	require.Nil(t, err)
}
