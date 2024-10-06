package folder_test

import (
	// "errors"
	"errors"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_folder_MoveFolder(t *testing.T) {
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
			Paths: "alpha.bravo.charlie",
			OrgId: org1,
		},
		{
			Name: "delta",
			Paths: "alpha.delta",
			OrgId: org1,
		},
		{
			Name: "echo",
			Paths: "alpha.delta.echo",
			OrgId: org1,
		},
		{
			Name: "foxtrot",
			Paths: "foxtrot",
			OrgId: org2,
		},
		{
			Name: "golf",
			Paths: "golf",
			OrgId: org1,
		},
	}

	tests := [...]struct {
		testCase 	string
		child     	string
		parent		string
		folders  	[]folder.Folder
		want     	[]folder.Folder
		expectError error
	}{
		{
			testCase:	"Successful move parent to child node",
			child:		"bravo",
			parent:		"delta",
			folders: 	originalFolders,
			want:		[]folder.Folder{
				{
					Name: "alpha",
					Paths: "alpha",
					OrgId: org1,
				  },
				  {
					Name: "bravo",
					Paths: "alpha.delta.bravo",
					OrgId: org1,
				  },
				  {
					Name: "charlie",
					Paths: "alpha.delta.bravo.charlie",
					OrgId: org1,
				  },
				  {
					Name: "delta",
					Paths: "alpha.delta",
					OrgId: org1,
				  },
				  {
					Name: "echo",
					Paths: "alpha.delta.echo",
					OrgId: org1,
				  },
				  {
					Name: "foxtrot",
					Paths: "foxtrot",
					OrgId: org2,
				  },
				  {
					Name: "golf",
					Paths: "golf",
					OrgId: org1,
				  },
			},
			expectError: nil,
		},
		{
			testCase:	"Another Success",
			child:		"bravo",
			parent:		"golf",
			folders: 	originalFolders,
			want:		[]folder.Folder{
				{
					Name: "alpha",
					Paths: "alpha",
					OrgId: org1,
				  },
				  {
					Name: "bravo",
					Paths: "golf.bravo",
					OrgId: org1,
				  },
				  {
					Name: "charlie",
					Paths: "golf.bravo.charlie",
					OrgId: org1,
				  },
				  {
					Name: "delta",
					Paths: "alpha.delta",
					OrgId: org1,
				  },
				  {
					Name: "echo",
					Paths: "alpha.delta.echo",
					OrgId: org1,
				  },
				  {
					Name: "foxtrot",
					Paths: "foxtrot",
					OrgId: org2,
				  },
				  {
					Name: "golf",
					Paths: "golf",
					OrgId: org1,
				  },
			},
			expectError: nil,
		},
		{
			testCase:	"Source folder does not exist",
			child:		"invalid_folder",
			parent:		"delta",
			folders: 	originalFolders,
			want:		nil,
			expectError: errors.New("Source folder does not exist"),
		},
		{
			testCase:	"Destination folder does not exist",
			child:		"bravo",
			parent:		"invalid_folder",
			folders: 	originalFolders,
			want:		nil,
			expectError: errors.New("Destination folder does not exist"),
		},
		{
			testCase:	"Cannot move a folder to a different organization",
			child:		"bravo",
			parent:		"foxtrot",
			folders: 	originalFolders,
			want:		nil,
			expectError: errors.New("Cannot move a folder to a different organization"),
		},
		{
			testCase:	"Cannot move a folder to itself",
			child:		"bravo",
			parent:		"bravo",
			folders: 	originalFolders,
			want:		nil,
			expectError: errors.New("Cannot move a folder to itself"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.testCase, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)

			got, err := f.MoveFolder(tt.child, tt.parent)

			// If an error is expected
			if tt.expectError != nil {
				assert.EqualError(t, err, tt.expectError.Error(), "Expected error did not occur")
			} else {
				assert.NoError(t, err, "Expected no error but returned")
			}

			// Asserting expected outputs (without errors)
			assert.Equal(t, tt.want, got, "Test Case %s failed", tt.testCase)
		})
	}

}
