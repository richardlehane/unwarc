package unwarc

import "testing"

var expect = []string{
	"http://www.records.nsw.gov.au/index.html?q=coal",
	"www.records.nsw.gov.au/",
	"index.html_q=coal",
	"http://www.records.nsw.gov.au",
	"",
	"www.records.nsw.gov.au",
}

func TestSanitise(t *testing.T) {
	for i := 0; i < len(expect); i = i + 3 {
		dir, filename := Sanitise(expect[i])
		if dir != expect[i+1] || filename != expect[i+2] {
			t.Errorf("Expected %s, %s; got %s, %s\n", expect[i+1], expect[i+2], dir, filename)
		}
	}
}

func TestBase(t *testing.T) {
	if Base("blackbooks.warc.gz") != "blackbooks" {
		t.Errorf("Expected blackbooks.warc.gz to return blackbooks, but got %s\n", Base("blackbooks.warc.gz"))
	}
}
