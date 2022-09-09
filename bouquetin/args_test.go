package bouquetin

import (
	"testing"
)

func TestParseArgs(t *testing.T) {
	empty := []string{}
	emptyButCommand := []string{"add"}
	noUrl := []string{"add", "--user", "a", "--note", "b"}
	correct := []string{"add", "https://github.com", "--user", "a", "--note", "b"}
	// weirdOrder := []string{"add", "https://github.com", "--note", "a", "--user", "b"}
	// double := []string{"add", "https://github.com", "--user", "a", "--title", "b", "--title c"}
	// doublePlus := []string{"add", "https://github.com", "--user", "a", "--title", "b", "--title", "c", "--note", "d"}
	// emptyLast := []string{"add", "https://github.com", "--user", "a", "--title", "b", "--title", "c", "--note"}

	_, _, _, _, _, err1 := ParseArgs(empty)
	_, _, _, _, _, err2 := ParseArgs(emptyButCommand)
	_, _, _, _, _, err3 := ParseArgs(noUrl)
	c1, url1, u1, t1, n1, err4 := ParseArgs(correct)
	// c2, url2, u2, t2, n2, err5 := ParseArgs(weirdOrder)
	// c3, url3, u3, t3, n3, err6 := ParseArgs(double)
	// c4, url4, u4, t4, n4, err7 := ParseArgs(doublePlus)
	// c5, url5, u5, t5, n5, err8 := ParseArgs(emptyLast)

	if err1 != nil {
		t.Errorf("empty: Empty args, should fail")
	}

	if err1.Error() != "no args provided" {
		t.Errorf("empty: Wrong error triggered")
	}

	if err2 != nil {
		t.Errorf("emptyButCommand: Empty command, should fail")
	}

	if err2.Error() != "wrong command" {
		t.Errorf("emptyButCommand: Wrong error triggered")
	}

	if err3 != nil {
		t.Errorf("noUrl: URL not specified, should fail")
	}

	if err3.Error() != "no url specified" {
		t.Errorf("noUrl: Wrong error triggered")
	}

	if c1 != "add" || url1 != "https://github.com" || u1 != "a" || t1 != "" || n1 != "b" {
		t.Errorf("correct: Unexpected output")

	}

	if err4 == nil {
		t.Errorf("correct: Raise an unexpected error")
	}
}
