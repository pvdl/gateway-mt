// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

package badgerauth_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"storj.io/common/rpc/rpcstatus"
	"storj.io/common/testcontext"
	"storj.io/gateway-mt/pkg/auth/badgerauth"
	"storj.io/gateway-mt/pkg/auth/badgerauth/badgerauthtest"
	"storj.io/gateway-mt/pkg/auth/badgerauth/pb"
)

func TestNodeAdmin_GetRecord(t *testing.T) {
	badgerauthtest.RunSingleNode(t, badgerauth.Config{
		ID: badgerauth.NodeID{'a', 'd', 'm', 'g', 'e', 't'},
	}, func(ctx *testcontext.Context, t *testing.T, node *badgerauth.Node) {
		admin := badgerauth.NewAdmin(node.DB)
		records, keys, _ := badgerauthtest.CreateFullRecords(ctx, t, node, 2)

		_, err := admin.GetRecord(ctx, &pb.GetRecordRequest{Key: []byte{}})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.NotFound)

		_, err = admin.GetRecord(ctx, &pb.GetRecordRequest{Key: []byte{'a'}})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.NotFound)

		_, err = admin.GetRecord(ctx, &pb.GetRecordRequest{Key: make([]byte, 33)})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.InvalidArgument)

		resp, err := admin.GetRecord(ctx, &pb.GetRecordRequest{Key: keys[0].Bytes()})
		require.NoError(t, err)
		require.Equal(t, records[keys[0]].EncryptedAccessGrant, resp.Record.EncryptedAccessGrant)

		resp, err = admin.GetRecord(ctx, &pb.GetRecordRequest{Key: keys[1].Bytes()})
		require.NoError(t, err)
		require.Equal(t, records[keys[1]].EncryptedAccessGrant, resp.Record.EncryptedAccessGrant)
	})
}

func TestNodeAdmin_InvalidateRecord(t *testing.T) {
	badgerauthtest.RunSingleNode(t, badgerauth.Config{
		ID: badgerauth.NodeID{'a', 'd', 'm', 'i', 'n', 'v'},
	}, func(ctx *testcontext.Context, t *testing.T, node *badgerauth.Node) {
		admin := badgerauth.NewAdmin(node.DB)
		records, keys, _ := badgerauthtest.CreateFullRecords(ctx, t, node, 2)

		_, err := admin.InvalidateRecord(ctx, &pb.InvalidateRecordRequest{Reason: "test"})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.NotFound)

		_, err = admin.InvalidateRecord(ctx, &pb.InvalidateRecordRequest{
			Key:    make([]byte, 33),
			Reason: "something",
		})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.InvalidArgument)

		_, err = admin.InvalidateRecord(ctx, &pb.InvalidateRecordRequest{
			Key:    []byte{'a'},
			Reason: "something",
		})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.NotFound)

		_, err = admin.InvalidateRecord(ctx, &pb.InvalidateRecordRequest{
			Key:    keys[0].Bytes(),
			Reason: "",
		})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.InvalidArgument)

		_, err = admin.InvalidateRecord(ctx, &pb.InvalidateRecordRequest{
			Key:    keys[0].Bytes(),
			Reason: "your key is disabled",
		})
		require.NoError(t, err)

		resp, err := admin.GetRecord(ctx, &pb.GetRecordRequest{Key: keys[0].Bytes()})
		require.NoError(t, err)
		require.Equal(t, records[keys[0]].EncryptedAccessGrant, resp.Record.EncryptedAccessGrant)
		require.Equal(t, "your key is disabled", resp.Record.InvalidationReason)
		require.NotZero(t, resp.Record.InvalidatedAtUnix)

		resp, err = admin.GetRecord(ctx, &pb.GetRecordRequest{Key: keys[1].Bytes()})
		require.NoError(t, err)
		require.Equal(t, records[keys[1]].EncryptedAccessGrant, resp.Record.EncryptedAccessGrant)
		require.Equal(t, "", resp.Record.InvalidationReason)
		require.Zero(t, resp.Record.InvalidatedAtUnix)
	})
}

func TestNodeAdmin_UnpublishRecord(t *testing.T) {
	badgerauthtest.RunSingleNode(t, badgerauth.Config{
		ID: badgerauth.NodeID{'a', 'd', 'm', 'p', 'u', 'b'},
	}, func(ctx *testcontext.Context, t *testing.T, node *badgerauth.Node) {
		admin := badgerauth.NewAdmin(node.DB)
		records, keys, _ := badgerauthtest.CreateFullRecords(ctx, t, node, 2)

		_, err := admin.UnpublishRecord(ctx, &pb.UnpublishRecordRequest{})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.NotFound)

		_, err = admin.UnpublishRecord(ctx, &pb.UnpublishRecordRequest{Key: make([]byte, 33)})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.InvalidArgument)

		_, err = admin.UnpublishRecord(ctx, &pb.UnpublishRecordRequest{Key: keys[0].Bytes()})
		require.NoError(t, err)

		resp, err := admin.GetRecord(ctx, &pb.GetRecordRequest{Key: keys[0].Bytes()})
		require.NoError(t, err)
		require.Equal(t, records[keys[0]].EncryptedAccessGrant, resp.Record.EncryptedAccessGrant)
		require.False(t, resp.Record.Public)

		resp, err = admin.GetRecord(ctx, &pb.GetRecordRequest{Key: keys[1].Bytes()})
		require.NoError(t, err)
		require.Equal(t, records[keys[1]].EncryptedAccessGrant, resp.Record.EncryptedAccessGrant)
		require.True(t, resp.Record.Public)
	})
}

func TestNodeAdmin_DeleteRecord(t *testing.T) {
	badgerauthtest.RunSingleNode(t, badgerauth.Config{
		ID: badgerauth.NodeID{'a', 'd', 'm', 'd', 'e', 'l'},
	}, func(ctx *testcontext.Context, t *testing.T, node *badgerauth.Node) {
		admin := badgerauth.NewAdmin(node.DB)
		records, keys, entries := badgerauthtest.CreateFullRecords(ctx, t, node, 2)

		badgerauthtest.VerifyReplicationLog{
			Entries: entries,
		}.Check(ctx, t, node.UnderlyingDB().UnderlyingDB())

		_, err := admin.DeleteRecord(ctx, &pb.DeleteRecordRequest{Key: []byte{'a'}})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.NotFound)

		_, err = admin.DeleteRecord(ctx, &pb.DeleteRecordRequest{Key: keys[0].Bytes()})
		require.NoError(t, err)

		_, err = admin.DeleteRecord(ctx, &pb.DeleteRecordRequest{Key: make([]byte, 33)})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.InvalidArgument)

		badgerauthtest.VerifyReplicationLog{
			Entries: []badgerauthtest.ReplicationLogEntryWithTTL{entries[1]},
		}.Check(ctx, t, node.UnderlyingDB().UnderlyingDB())

		_, err = admin.GetRecord(ctx, &pb.GetRecordRequest{Key: keys[0].Bytes()})
		require.Equal(t, rpcstatus.Code(err), rpcstatus.NotFound)

		resp, err := admin.GetRecord(ctx, &pb.GetRecordRequest{Key: keys[1].Bytes()})
		require.NoError(t, err)
		require.Equal(t, resp.Record.EncryptedAccessGrant, records[keys[1]].EncryptedAccessGrant)
	})
}