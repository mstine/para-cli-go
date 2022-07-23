/*
Copyright © 2022 Matt Stine

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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

func TestRemove(t *testing.T) {
	testCases := []struct {
		name      string
		rootName  string
		expectLen int
		expectErr error
	}{
		{"RemoveExisting", "root1", 1, nil},
		{"RemoveNotFound", "root3", 1, root.ErrNotExists},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			prl := &root.ParaRootList{
				Roots: []root.ParaRoot{
					{
						Name: "root1",
						Path: "path1",
					},
					{
						Name: "root2",
						Path: "path2",
					},
				},
			}
			err := prl.Remove(tc.rootName)
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
			if prl.Roots[0].Name == tc.rootName {
				t.Errorf("PARA Root name %q should not be in the list\n", tc.rootName)
			}
		})
	}
}
