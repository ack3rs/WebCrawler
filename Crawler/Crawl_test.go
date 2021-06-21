package Crawler

import "testing"

func TestSanatiseURL(t *testing.T) {

	testData := []struct {
		url    string
		result string
	}{
		{"https://www.acky.com.com/bob/uncle/#respond", "https://www.acky.com.com/bob/uncle/"},
		{"https://www.acky.com.com/bob/uncle/?bob=here#respond", "https://www.acky.com.com/bob/uncle/?bob=here"},
		{"https://www.acky.com/#", "https://www.acky.com/"},
		{"https://www.acky.com/?page=here#", "https://www.acky.com/?page=here"},
	}

	for _, testItem := range testData {
		result, _ := sanatiseURL(testItem.url)

		if result != testItem.result {
			t.Errorf("Error testing '%s' got '%s' expected '%s'", testItem.url, result, testItem.result)
		}
	}

}

func TestCleanURL(t *testing.T) {

	testData := []struct {
		url    string
		result string
	}{
		{"https://www.acky.com.com/bob/uncle/#respond", "https://www.acky.com.com/bob/uncle/"},
		{"https://www.acky.com.com/bob/uncle/?bob=here#respond", "https://www.acky.com.com/bob/uncle/?bob=here"},
		{"https://www.acky.com/#", "https://www.acky.com/"},
	}

	for _, testItem := range testData {
		result, _ := cleanURL(testItem.url, testItem.url)

		if result != testItem.result {
			t.Errorf("Error testing '%s' got '%s' expected '%s'", testItem.url, result, testItem.result)
		}
	}
}

func TestDomainMatch(t *testing.T) {

	testData := []struct {
		D1     string
		D2     string
		result bool
	}{
		{"https://www.acky.com.com/bob/uncle/#respond", "https://www.acky.com.com/bob/uncle/", true},
		{"https://www.acky.com.com/bob/uncle/?bob=here#respond", "https://www.acky.com.com/bob/uncle/?bob=here", true},
		{"https://www.acky.com/#", "https://www.acky.com/", true},
		{"https://www.aky.com/#", "https://www.acky.com/", false},
	}

	for _, testItem := range testData {
		result := doDomainsMatch(testItem.D1, testItem.D2)

		if result != testItem.result {
			t.Errorf("Error testing '%s' with '%s' got '%v' expected '%v'", testItem.D1, testItem.D2, result, testItem.result)
		}
	}

}
