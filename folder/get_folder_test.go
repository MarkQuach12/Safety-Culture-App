package folder_test

import (
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
	org1 := uuid.Must(uuid.NewV4())
	org2 := uuid.Must(uuid.NewV4())

	originalFolders := 	[]folder.Folder{
		{
			Name: "alpha",
			Paths: "alpha",
			OrgId: org1,
		  },
		  {
			Name: "bravo",
			Paths: "alpha.bravo",
			OrgId: org1,
		  },
		  {
			Name: "charlie",
			Paths : "alpha.bravo.charlie",
			OrgId: org1,
		  },
		  {
			Name: "delta",
			Paths: "alpha.delta",
			OrgId: org1,
		  },
		  {
			Name: "echo",
			Paths: "echo",
			OrgId: org1,
		  },
		  {
			Name: "foxtrot",
			Paths: "foxtrot",
			OrgId: org2,
		  },
	}

	tests := [...]struct {
		testCase 	string
		name     	string
		orgID    	uuid.UUID
		folders  	[]folder.Folder
		want     	[]folder.Folder
		expectError error
	}{
		{
			testCase:	"Name is the first node",
			name:		"alpha",
			orgID:		org1,
			folders: 	originalFolders,
			want:		[]folder.Folder{
				{
					Name: "bravo",
					Paths: "alpha.bravo",
					OrgId: org1,
				},
				{
					Name: "charlie",
					Paths : "alpha.bravo.charlie",
					OrgId: org1,
				},
				{
					Name: "delta",
					Paths: "alpha.delta",
					OrgId: org1,
				},
			},
			expectError: nil,
		},
		{
			testCase:	"Name is the second node",
			name:		"bravo",
			orgID:		org1,
			folders: 	originalFolders,
			want:		[]folder.Folder{
				{
					Name: "charlie",
					Paths: "alpha.bravo.charlie",
					OrgId: org1,
				},
			},
			expectError: nil,
		},
		{
			testCase:	"No child (when node exist with parent path)",
			name:		"charlie",
			orgID:		org1,
			folders: 	originalFolders,
			want:		[]folder.Folder{},
			expectError: nil,
		},
		{
			testCase:	"No child (when no parent path)",
			name:		"echo",
			orgID:		org1,
			folders: 	originalFolders,
			want:		[]folder.Folder{},
			expectError: nil,
		},
		{
			testCase:	"Folder does not exist",
			name:		"invalid_folder",
			orgID:		org1,
			folders: 	originalFolders,
			want:		nil,
			expectError: errors.New("Folder does not exist"),
		},
		{
			testCase:	"Folder does not exist in the specified organization",
			name:		"foxtrot",
			orgID:		org1,
			folders: 	originalFolders,
			want:		nil,
			expectError: errors.New("Folder does not exist in the specified organization"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)

			got, err := f.GetAllChildFolders(tt.orgID, tt.name)

			// If an error is expected
			if tt.expectError != nil {
				assert.EqualError(t, err, tt.expectError.Error(), "Expected error did not occur")
			} else {
				assert.NoError(t, err, "Expected no error but returned")
			}

			// Asserting expected outputs (without errors)
			assert.Equal(t, tt.want, got, "Test case %s failed", tt.testCase)
		})
	}
}
