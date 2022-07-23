package root

import (
	"errors"
	"fmt"

	"golang.org/x/exp/slices"
)

var (
	ErrExists    = errors.New("PARA Root already exists")
	ErrNotExists = errors.New("PARA Root does not exist")
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
