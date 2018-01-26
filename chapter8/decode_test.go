package main
import (
	"testing"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.ID != 1{
		t.Error("Wrong id, was expecting 1 but got ", post.ID)
	}
}