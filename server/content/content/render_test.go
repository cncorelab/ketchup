package content

import (
	"html/template"
	"reflect"
	"testing"

	"github.com/ketchuphq/ketchup/proto/ketchup/models"
)

type renderTextContentCase struct {
	t models.ContentTextType
	c string

	expected interface{}
}

func (r renderTextContentCase) GetType() models.ContentTextType {
	return r.t
}

func TestRenderTextContent(t *testing.T) {
	cases := []renderTextContentCase{
		{
			t:        models.ContentTextType_text,
			c:        "<div>**Hello**</div>",
			expected: "<div>**Hello**</div>",
		},
		{
			t:        models.ContentTextType_html,
			c:        "<div>**Hello**</div>",
			expected: template.HTML("<div>**Hello**</div>"),
		},
		{
			t:        models.ContentTextType_markdown,
			c:        "**hello**",
			expected: template.HTML("<p><strong>hello</strong></p>\n"),
		},
	}
	for _, k := range cases {
		v, err := renderTextualContent(k.c, k)
		if err != nil {
			t.Fatal(k, err)
		}
		if !reflect.DeepEqual(v, k.expected) {
			t.Errorf("unexpected %#v", v)
		}
	}
}
