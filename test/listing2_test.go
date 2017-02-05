// Sample test to show how to write a basic unit table test.
package test

import (
	"net/http"
	"testing"
)

const checkMark2 = "\u2713"
const ballotX2 = "\u2717"

// TestDownload validates the http Get function can download
// content and handles different status conditions properly.
func TestDownload2(t *testing.T) {
	var urls = []struct {
		url        string
		statusCode int
	}{
		{
			"http://www.google.com",
			http.StatusOK,
		},
		{
			"http://rss.cnn.com/rss/cnn_topstbadurl.rss",
			http.StatusNotFound,
		},
	}

	t.Log("Given the need to test downloading different content.")
	{
		for _, u := range urls {
			t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
				u.url, u.statusCode)
			{
				resp, err := http.Get(u.url)
				if err != nil {
					t.Fatal("\t\tShould be able to Get the url.",
						ballotX2, err)
				}
				t.Log("\t\tShould be able to Get the url",
					checkMark2)

				defer resp.Body.Close()

				if resp.StatusCode == u.statusCode {
					t.Logf("\t\tShould be a \"%d\" status. %v",
						u.statusCode, checkMark2)
				} else {
					t.Errorf("\t\tShould have a \"%d\" status %v %v",
						u.statusCode, ballotX2, resp.StatusCode)
				}
			}
		}
	}
}
