package root_test

import (
	"errors"
	"testing"

	"github.com/mstine/para-cli/root"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name      string
		rootName  string
		rootPath  string
		expectLen int
		expectErr error
	}{
		{"AddNew", "root2", "path2", 2, nil},
		{"AddExisting", "root1", "path1", 1, root.ErrExists},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			prl := &root.ParaRootList{
				Roots: []root.ParaRoot{
					{
						Name: "root1",
						Path: "path1",
					},
				},
			}
			err := prl.Add(root.ParaRoot{
				Name: tc.rootName,
				Path: tc.rootPath,
			})
			if tc.expectErr != nil {
				if err == nil {
					t.Fatalf("Expected error, got nil instead\n")
				}
				if !errors.Is(err, tc.expectErr) {
					t.Errorf("Expected error %q, got %q instead\n",
						tc.expectErr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("Expected no error, got %q instead\n", err)
			}
			if len(prl.Roots) != tc.expectLen {
				t.Errorf("Expected list length %d, got %d instead\n", tc.expectLen, len(prl.Roots))
			}
			if prl.Roots[1].Name != tc.rootName {
				t.Errorf("Expected PARA Root name %q as index 1, got %q instead\n", tc.rootName, prl.Roots[1].Name)
			}
		})
	}
}
