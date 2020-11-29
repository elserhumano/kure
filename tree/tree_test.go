package tree_test

import (
	"strings"
	"testing"

	"github.com/GGP1/kure/tree"
)

func TestTreePrint(t *testing.T) {
	paths := []string{
		"kure/atoll/password",
		"kure/atoll/passphrase",
		"sync/atomic",
		"unsafe/pointer",
	}

	tree.Print(paths)
	// Output:
	// ├── kure
	// │   └── atoll
	// │       ├── password
	// │       └── passphrase
	// ├── sync
	// │   └── atomic
	// └── unsafe
	//     └── pointer
}

func TestTreeStructure(t *testing.T) {
	paths := []string{
		"The Hobbit",
		"The Lord of the Rings/The fellowship of the ring",
		"The Lord of the Rings/The two towers",
		"The Lord of the Rings/The return of the king",
	}

	root := tree.Root(paths)
	folders := make(map[string]struct{}, len(paths))

	for _, p := range paths {
		if _, ok := folders[p]; !ok {
			s := strings.Split(p, "/")[0]
			folders[s] = struct{}{}
		}
	}

	expected := len(folders)
	if len(root.Children) != expected {
		t.Errorf("Expected %d branches, got %d", expected, len(root.Children))
	}

	for i, r := range root.Children {
		name := strings.Split(paths[i], "/")[0]

		if r.Name != name {
			t.Errorf("Expected branch name to be %q, got %q", name, r.Name)
		}

		if i == len(root.Children)-1 {
			if len(r.Children) == 0 {
				t.Errorf("Expected %q branch to have a child named %q", r.Name, r.Children[0])
			}
		}
	}
}
