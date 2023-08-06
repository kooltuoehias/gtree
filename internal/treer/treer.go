package treer

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kooltuoehias/gtree/util"
)

type FolderFiles struct {
	Path     string
	MaxDepth int
}

func (ft FolderFiles) Traverse(level int) {

	filepath.WalkDir(ft.Path, func(path string, dirEntry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == ft.Path {
			return nil
		}
		key, _ := strings.CutPrefix(path, filepath.FromSlash(ft.Path)+string(os.PathSeparator))
		pes := strings.Split(key, string(os.PathSeparator))
		lpes := len(pes)
		if level != 0 && lpes >= level {
			return nil
		}
		for i := 1; i < lpes; i++ {
			fmt.Printf("  ")
		}
		fmt.Printf("\u2514\u2500\u2500 %s\n", pes[lpes-1])
		ft.MaxDepth = util.Max(ft.MaxDepth, len(strings.Split(key, string(os.PathSeparator))))
		return nil
	})

}
