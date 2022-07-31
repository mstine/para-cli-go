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
package root

import (
	"errors"
	"fmt"
	"os"

	"golang.org/x/exp/slices"
)

var (
	ErrExists           = errors.New("PARA Root already exists")
	ErrNotExists        = errors.New("PARA Root does not exist")
	ErrPathNotDirectory = errors.New("path specified for PARA Root is not a directory")
)

type ParaRoot struct {
	Name string // must be unique
	Path string // must be unique
}

type ParaRootList struct {
	Roots []ParaRoot
}

func (list *ParaRootList) Add(root ParaRoot) error {
	i := slices.IndexFunc(list.Roots, func(r ParaRoot) bool {
		return root.Name == r.Name || root.Path == r.Path
	})
	if i > -1 {
		return fmt.Errorf("%w: %s", ErrExists, root)
	}
	if fi, err := os.Stat(root.Path); err != nil {
		if !os.IsNotExist(err) {
			// If os.Stat returns any other error than "does not exist," propagate the error
			return fmt.Errorf("%w: %s", err, root)
		} else {
			// If os.Stat returns "does not exist" error, make the directory path
			os.Mkdir(root.Path, 0755)
		}
	} else {
		// We can assume path exists - make sure it's a directory
		if !fi.IsDir() {
			return fmt.Errorf("%w: %s", ErrPathNotDirectory, root)
		}
	}
	list.Roots = append(list.Roots, root)
	return nil
}

func (list *ParaRootList) Remove(name string) error {
	i := slices.IndexFunc(list.Roots, func(r ParaRoot) bool {
		return r.Name == name
	})
	if i > -1 {
		list.Roots = append(list.Roots[:i], list.Roots[i+1:]...)
		return nil
	}
	return fmt.Errorf("%w: %s", ErrNotExists, name)
}
