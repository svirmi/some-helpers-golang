package finance

import (
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// As we are calling an external API, we need to mock it out to be able to test all the necessary scenarios.
// To help doing that we'll use the net/http/httptest package.
// This package provides utilities for HTTP testing, namely to simulate a server.
// The test for the happy flow is:

func TestConvertCurrency(t *testing.T) {

	expected := `
		{
			"code": 200,
			"info": {
				"_t": "2022-02-12 21:49:18 UTC",
				"credit_count": 1,
				"server_time": "2022-02-12 21:49:18 UTC"
			},
			"msg": "Successfully",
			"response": [
				{
					"c": "1.13268",
					"ch": "-0.00013",
					"cp": "-0.01%",
					"h": "1.13281",
					"id": "1",
					"l": "1.13246",
					"o": "1.13281",
					"s": "EUR/USD",
					"t": "1640638800",
					"tm": "2022-02-12 21:00:00",
					"up": "2022-02-12 21:49:10"
				}
			],
			"status": true
		}
	`

	// The httptest.NewServer will start a local server which will return a predefined response.
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// The output of this function is a httptest.Server struct containing an URL we can use to call the NewFinanceFunctions constructor.
		// So when we call the ConvertCurrency function, the URL for the test server will be used instead of a real one.
		fmt.Fprint(w, expected)
	}))

	defer svr.Close()

	ff := NewFinanceFunctions(svr.URL, "DummyApiKey")

	result, err := ff.ConvertCurrency("EUR", "USD", 10)

	assert.Nil(t, err)
	// As Golang does not have a standard library to round numbers with a certain number of decimal places, we use math.Round(result*10000)/10000) to achieve the same result.
	assert.Equal(t, 11.3268, math.Round(result*10000)/10000)
}

func TestConvertCurrencyWithApiError(t *testing.T) {

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
	}))

	defer svr.Close()

	ff := NewFinanceFunctions(svr.URL, "DummyApiKey")

	result, err := ff.ConvertCurrency("EUR", "USD", 10)

	assert.NotNil(t, err)
	assert.Equal(t, 0.0, result)
}

func TestConvertCurrencyWithInvalidJsonBody(t *testing.T) {

	expected := "invalid json"

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))

	defer svr.Close()

	ff := NewFinanceFunctions(svr.URL, "DummyApiKey")

	result, err := ff.ConvertCurrency("EUR", "USD", 10)

	assert.NotNil(t, err)
	assert.Equal(t, 0.0, result)
}

func TestNewFinanceFunctions(t *testing.T) {
	ff := NewFinanceFunctions("SomeApiUrl", "SomeApiKey")

	assert.Equal(t, "SomeApiUrl", ff.ApiUrl)
	assert.Equal(t, "SomeApiKey", ff.ApiKey)
}

func TestNewFinanceFunctionsWithEmptyApiUrl(t *testing.T) {
	ff := NewFinanceFunctions("", "SomeApiKey")

	assert.NotEmpty(t, ff.ApiUrl)
	assert.Contains(t, ff.ApiUrl, "fcsapi.com")
}
