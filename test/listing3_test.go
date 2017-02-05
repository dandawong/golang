// Sample test to show how to mock an HTTP GET call internally.
// Differs slightly from the book to show more.
package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const checkMark3 = "\u2713"
const ballotX3 = "\u2717"

// feed is mocking the XML document we except to receive.
var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
<channel>
	<title>Going Go Programming</title>
	<description>Golang : https://github.com/goinggo</description>
	<link>http://www.goinggo.net/</link>
	<item>
		<pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
		<title>Object Oriented Programming Mechanics</title>
		<description>Go is an object oriented language.</description>
		<link>http://www.goinggo.net/2015/03/object-oriented</link>
	</item>
</channel>
</rss>`

// mockServer returns a pointer to a server to handle the get call.
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload validates the http Get function can download content
// and the content can be unmarshaled and clean.
func TestDownload3(t *testing.T) {
	statusCode := http.StatusOK

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tWhen checking \"%s\" for status code \"%d\"",
			server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatal("\t\tShould be able to make the Get call.",
					ballotX3, err)
			}
			t.Log("\t\tShould be able to make the Get call.",
				checkMark3)

			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t\tShould receive a \"%d\" status. %v %v",
					statusCode, ballotX3, resp.StatusCode)
			} else {
				t.Logf("\t\tShould receive a \"%d\" status. %v",
					statusCode, checkMark3)
			}
		}
	}
}
