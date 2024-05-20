package sqlc

import (
	"Immortals/internal/util"
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomNode(t *testing.T) Nodes {
	arg := CreateNodeParams{
		Name:     sql.NullString{String: util.RandomName(), Valid: true},
		ClientID: sql.NullString{String: util.RandomClientID(), Valid: true},
	}

	node, err := testQueries.CreateNode(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, node)

	require.Equal(t, arg.Name, node.Name)
	require.Equal(t, arg.ClientID, node.ClientID)

	require.NotZero(t, node.ID)
	return node
}
func TestCreateNode(t *testing.T) {
	createRandomNode(t)
}

func TestGetNode(t *testing.T) {
	node1 := createRandomNode(t)
	node2, err := testQueries.GetNode(context.Background(), node1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, node2)

	require.Equal(t, node1.ID, node2.ID)
	require.Equal(t, node1.Name, node2.Name)
}

func TestUpdateAccount(t *testing.T) {
	node1 := createRandomNode(t)

	arg := UpdateNodeParams{
		ID:   node1.ID,
		Name: sql.NullString{String: util.RandomName(), Valid: true},
	}

	node2, err := testQueries.UpdateNode(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, node2)

	require.Equal(t, node1.ID, node2.ID)
	require.Equal(t, arg.Name, node2.Name)
}

func TestDeleteNode(t *testing.T) {
	node1 := createRandomNode(t)
	err := testQueries.DeleteNode(context.Background(), node1.ID)
	require.NoError(t, err)

	node2, err := testQueries.GetNode(context.Background(), node1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, node2)
}

func TestListNodes(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomNode(t)
	}

	nodes, err := testQueries.ListNodes(context.Background())
	require.NoError(t, err)
	for _, node := range nodes {
		require.NotEmpty(t, node)
	}
}
