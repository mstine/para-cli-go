package cmd

import (
	"bytes"
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
			expectedOut:    "Added PARA Root:  {root1 path1}\n",
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
			expectedOut:    "Removed PARA Root:  root1\n",
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
