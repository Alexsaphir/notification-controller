package notifier

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLark_Post(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		var payload LarkPayload
		err = json.Unmarshal(b, &payload)
		require.NoError(t, err)

		require.Equal(t, "💫 gitrepository/webapp.gitops-system", payload.Card.Header.Title.Content)
		require.Equal(t, "turquoise", payload.Card.Header.Template)
	}))
	defer ts.Close()

	lark, err := NewLark(ts.URL)
	require.NoError(t, err)

	err = lark.Post(context.TODO(), testEvent())
	require.NoError(t, err)
}
