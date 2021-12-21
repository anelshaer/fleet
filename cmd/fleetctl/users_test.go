package main

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/fleetdm/fleet/v4/server/fleet"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserDelete(t *testing.T) {
	_, ds := runServerWithMockedDS(t)

	ds.UserByEmailFunc = func(ctx context.Context, email string) (*fleet.User, error) {
		return &fleet.User{
			ID:    42,
			Name:  "test1",
			Email: "user1@test.com",
		}, nil
	}

	deletedUser := uint(0)

	ds.DeleteUserFunc = func(ctx context.Context, id uint) error {
		deletedUser = id
		return nil
	}

	assert.Equal(t, "", runAppForTest(t, []string{"user", "delete", "--email", "user1@test.com"}))
	assert.Equal(t, uint(42), deletedUser)
}

func writeTmpCsv(t *testing.T, contents string) string {
	tmpFile, err := ioutil.TempFile(t.TempDir(), "*.csv")
	require.NoError(t, err)
	_, err = tmpFile.WriteString(contents)
	require.NoError(t, err)
	return tmpFile.Name()
}

func TestCreateBulkUsers(t *testing.T) {
	_, ds := runServerWithMockedDS(t)
	ds.InviteByEmailFunc = func(ctx context.Context, email string) (*fleet.Invite, error) {
		return nil, nil
	}

	csvFile := writeTmpCsv(t,
		`Name,Email,Password,SSO,API,Roles,Team
	user11,user11@domain.com,P@ssw0rd!2,false,false,,
	user12,user12@domain.com,P@ssw0rd!2,false,false,,`)

	assert.Equal(t, "", runAppForTest(t, []string{"user", "import", "--csv", csvFile}))
}
