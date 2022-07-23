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
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/mstine/para-cli/root"
)

func setUp(initList bool) root.ParaRootList {
	if !initList {
		return root.ParaRootList{
			Roots: []root.ParaRoot{},
		}
	} else {
		return root.ParaRootList{
			Roots: []root.ParaRoot{
				{
					Name: "root1",
					Path: "path1",
				},
				{
					Name: "root2",
					Path: "path2",
				},
				{
					Name: "root3",
					Path: "path3",
				},
			},
		}
	}
}

func TestHostActions(t *testing.T) {
	testCases := []struct {
		name           string
		args           []string
		expectedOut    string
		initList       bool
		actionFunction func(io.Writer, *root.ParaRootList, []string) error
	}{
		{
			name:           "AddAction",
			args:           []string{"root1", "path1"},
			expectedOut:    "Added PARA Root: {root1 path1}\n",
			initList:       false,
			actionFunction: addAction,
		},
		{
			name:           "ListAction",
			expectedOut:    "{root1 path1}\n{root2 path2}\n{root3 path3}\n",
			initList:       true,
			actionFunction: listAction,
		},
		{
			name:           "RemoveAction",
			args:           []string{"root1"},
			expectedOut:    "Removed PARA Root: root1\n",
			initList:       true,
			actionFunction: removeAction,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			roots := setUp(tc.initList)

			var out bytes.Buffer

			if err := tc.actionFunction(&out, &roots, tc.args); err != nil {
				t.Fatalf("Expected no error, got %q\n", err)
			}

			if out.String() != tc.expectedOut {
				t.Errorf("Expected output %q, got %q\n", tc.expectedOut, out.String())
			}
		})
	}
}

func TestIntegration(t *testing.T) {
	var out bytes.Buffer

	initial := setUp(true)

	delRoot := "root2"

	final := root.ParaRootList{
		Roots: []root.ParaRoot{
			{
				Name: "root1",
				Path: "path1",
			},
			{
				Name: "root3",
				Path: "path3",
			},
		},
	}

	expectedOut := ""
	for _, v := range initial.Roots {
		expectedOut += fmt.Sprintf("Added PARA Root: %s\n", v)
	}
	for _, v := range initial.Roots {
		expectedOut += fmt.Sprintf("%s\n", v)
	}
	expectedOut += fmt.Sprintf("Removed PARA Root: %s\n", delRoot)
	for _, v := range final.Roots {
		expectedOut += fmt.Sprintf("%s\n", v)
	}

	roots := root.ParaRootList{
		Roots: []root.ParaRoot{},
	}

	for _, v := range initial.Roots {
		if err := addAction(&out, &roots, []string{v.Name, v.Path}); err != nil {
			t.Fatalf("Expected no error, got %q\n", err)
		}
	}

	if err := listAction(&out, &roots, nil); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	if err := removeAction(&out, &roots, []string{delRoot}); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	if err := listAction(&out, &roots, nil); err != nil {
		t.Fatalf("Expected no error, got %q\n", err)
	}

	if out.String() != expectedOut {
		t.Errorf("Expected output %q, got %q\n", expectedOut, out.String())
	}
}
