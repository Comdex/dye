package dye

import (
	"testing"
)

func expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v, got %v", b, a)
	}
}

func refute(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v, got %v", b, a)
	}
}

func Test_Dye_Binary(t *testing.T) {
	Binary("/usr/local/bin/pygmentize")
}

func Test_Dye_Which(t *testing.T) {
	which := Which()
	expect(t, which, "/usr/local/bin/pygmentize")
}

func Test_Dye_Highlight(t *testing.T) {
	output, _ := Highlight("func main() {}", "go", "html", "utf-8")
	expect(t, output, "<div class=\"highlight\"><pre><span class=\"kd\">func</span> <span class=\"nx\">main</span><span class=\"p\">()</span> <span class=\"p\">{}</span>\n</pre></div>\n")
}
