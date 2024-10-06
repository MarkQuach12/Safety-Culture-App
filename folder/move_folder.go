package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	folders := f.folders

	var sourceFolder *Folder
	var destinationFolder *Folder

	for i := range folders {
		if folders[i].Name == name {
			sourceFolder = &folders[i]
		}
		if folders[i].Name == dst {
			destinationFolder = &folders[i]
		}
	}

	if sourceFolder == nil {
		return nil, errors.New("Source folder does not exist")
	}

	if destinationFolder == nil {
		return nil, errors.New("Destination folder does not exist")
	}

	if sourceFolder.OrgId != destinationFolder.OrgId {
		return nil, errors.New("Cannot move a folder to a different organization")
	}

	if sourceFolder.Paths == destinationFolder.Paths {
		return nil, errors.New("Cannot move a folder to itself")
	}

	children, err := f.GetAllChildFolders(sourceFolder.OrgId, sourceFolder.Name)
	if err != nil {
		return nil, err
	}

	// Destination Path
	parentPath := destinationFolder.Paths

	children = append(children, *sourceFolder)

	for i := range children {
		index := strings.Index(children[i].Paths, name)
		newPath := parentPath + "." + children[i].Paths[index:]
		children[i].Paths = newPath
	}

	for i := range folders {
		for _, child := range children {
			if folders[i].Name == child.Name {
				folders[i] = child
			}
		}
	}

	return folders, nil
}