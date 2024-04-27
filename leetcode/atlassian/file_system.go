package atlassian

import "strings"

// https://leetcode.com/problems/design-file-system/
type FileSystem struct {
	dir *directory
}

type directory struct {
	name    string
	value   int
	dir_map map[string]*directory
}

func Constructor3() FileSystem {
	return FileSystem{
		dir: &directory{
			name:    "",
			value:   -1,
			dir_map: make(map[string]*directory),
		},
	}
}

func (this *FileSystem) CreatePath(path string, value int) bool {
	paths, d := strings.Split(path, "/")[1:], this.dir
	for i := 0; i < len(paths); i++ {
		if _, ok := d.dir_map[paths[i]]; !ok {
			if i == len(paths)-1 {
				d.dir_map[paths[i]] = &directory{
					name:  paths[i],
					value: -1,
					dir_map: make(map[string]*directory),
				}
			} else {
				return false
			}
		}
		d = d.dir_map[paths[i]]
	}

	if d.value != -1 {
		return false
	}
	d.value = value

	return true
}

func (this *FileSystem) Get(path string) int {
	paths, d := strings.Split(path, "/")[1:], this.dir
	for i := 0; i < len(paths); i++ {
		if _, ok := d.dir_map[paths[i]]; !ok {
			return -1
		}
		d = d.dir_map[paths[i]]
	}

	return d.value
}

// type FileSystem struct {
// 	dir map[string]int
// }

// func Constructor3() FileSystem {
// 	return FileSystem{
// 		dir: make(map[string]int),
// 	}
// }

// func (this *FileSystem) CreatePath(path string, value int) bool {
// 	if _, ok := this.dir[path]; ok || path == "" || path == "/" {
// 		return false
// 	}

// 	paths := strings.Split(path, "/")
// 	parent := strings.Join(paths[:len(paths)-1], "/")

// 	if _, ok := this.dir[parent]; !ok && len(parent) > 1 {
// 		return false
// 	}

// 	this.dir[path] = value

// 	return true
// }

// func (this *FileSystem) Get(path string) int {
// 	if _, ok := this.dir[path]; !ok {
// 		return -1
// 	}

// 	return this.dir[path]
// }
