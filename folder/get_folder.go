package folder

import (
	"errors"
	"strings"

	"github.com/gofrs/uuid"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	filteredFolders := f.GetFoldersByOrgID(orgID)

	// Search for if file exist in orgID
	var folderExistsInOrg bool
	var parentPath string
	for _, f:= range f.folders {
		if f.Name == name && f.OrgId == orgID {
			parentPath = f.Paths + "."
			folderExistsInOrg = true
			break
		}
	}

	// Continues the search if it does not exist (eg it exist but not in the same org)
	if !folderExistsInOrg && parentPath == "" {
		for _, folder := range f.folders {
			if folder.Name == name {
				return nil, errors.New("Folder does not exist in the specified organization")
			}
		}
	}

	// Folder does not exist in the specified organization
	if !folderExistsInOrg && parentPath == "" {
		return nil, errors.New("Folder does not exist")
	}

	children := []Folder{}
	for _, f := range filteredFolders {
		if strings.HasPrefix(f.Paths, parentPath) {
			children = append(children, f)
		}
	}

	return children, nil
}
