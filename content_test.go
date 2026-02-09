package liber

import (
	"reflect"
	"testing"

	"github.com/javiorfo/go-liber/body"
	"github.com/javiorfo/go-liber/internal/epub"
	"github.com/javiorfo/go-liber/reftype"
)

func TestContentBuilder(t *testing.T) {
	mockBody := body.Raw("test-body")
	rt := reftype.Text("test-ref")
	filename := "test.epub"

	builder := ContentBuilder(mockBody, rt).Filename(filename)
	content := builder.Build()

	if content.Body != mockBody {
		t.Errorf("expected body %v, got %v", mockBody, content.Body)
	}
	if content.ReferenceType != rt {
		t.Errorf("expected reference type %v, got %v", rt, content.ReferenceType)
	}
	if content.Filename.IsNil() {
		t.Errorf("expected filename %s, got %s", filename, content.Filename)
	}

	child := epub.Content{Body: body.Raw("child")}
	builder.AddChildren(child)
	if len(builder.SubContents) != 1 {
		t.Errorf("expected 1 child, got %d", len(builder.SubContents))
	}

	ref := epub.ContentReference{Title: "ref-title"}
	builder.AddContentReferences(ref)
	if len(builder.ContentReferences) != 1 {
		t.Errorf("expected 1 content reference, got %d", len(builder.ContentReferences))
	}
}

func TestContentReferenceBuilder(t *testing.T) {
	title := "Chapter 1"
	id := "ch1-id"

	builder := ContentReferenceBuilder(title).ID(id)
	ref := builder.Build()

	if ref.Title != title {
		t.Errorf("expected title %s, got %s", title, ref.Title)
	}
	if ref.ID.AsValue() != id {
		t.Errorf("expected ID %s, got %s", id, ref.ID)
	}

	subRef := epub.ContentReference{Title: "Sub-section"}
	builder.AddChildren(subRef)

	if len(builder.SubContentReferences) != 1 {
		t.Errorf("expected 1 sub-reference, got %d", len(builder.SubContentReferences))
	}
	if builder.SubContentReferences[0].Title != "Sub-section" {
		t.Errorf("expected sub-reference title 'Sub-section', got %s", builder.SubContentReferences[0].Title)
	}
}

func TestBuilderChaining(t *testing.T) {
	res := ContentReferenceBuilder("Root").
		ID("root-id").
		AddChildren(epub.ContentReference{Title: "Child 1"}).
		AddChildren(epub.ContentReference{Title: "Child 2"}).
		Build()

	if len(res.SubContentReferences) != 2 {
		t.Errorf("expected 2 children from chained calls, got %d", len(res.SubContentReferences))
	}
}

func TestMakeSlice(t *testing.T) {
	t.Run("Returns empty slice when no arguments passed", func(t *testing.T) {
		result := MakeSlice[epub.Content]()

		if result == nil {
			t.Error("Expected an empty slice, but got nil")
		}
		if len(result) != 0 {
			t.Errorf("Expected length 0, got %d", len(result))
		}
	})

	t.Run("Handles epub.Content values", func(t *testing.T) {
		c1 := epub.Content{}
		c2 := epub.Content{}

		result := MakeSlice(c1, c2)

		expected := []epub.Content{c1, c2}
		if !reflect.DeepEqual(result, expected) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Handles epub.ContentReference values", func(t *testing.T) {
		ref := epub.ContentReference{}

		result := MakeSlice(ref)

		if len(result) != 1 {
			t.Errorf("Expected length 1, got %d", len(result))
		}
	})
}
