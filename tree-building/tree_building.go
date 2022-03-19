package tree

import (
	"errors"
	"sort"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	sort.Slice(records, func(i, j int) bool {
		return records[i].ID < records[j].ID
	})

	tree := make([]*Node, len(records))
	for i, record := range records {
		switch {
		case record.ID != i:
			return nil, errors.New("non-continuous")
		case record.ID != 0 && record.ID == record.Parent:
			return nil, errors.New("cycle")
		case record.Parent > record.ID:
			return nil, errors.New("higher id parent of lower id")
		}
		tree[record.ID] = &Node{record.ID, nil}
		if record.ID != 0 {
			tree[record.Parent].Children = append(tree[record.Parent].Children, tree[record.ID])
		}
	}
	return tree[0], nil
}
