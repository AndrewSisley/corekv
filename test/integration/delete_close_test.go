package integration

import (
	"testing"

	"github.com/sourcenetwork/corekv/test/action"
	"github.com/sourcenetwork/corekv/test/state"
)

func TestDeleteClose_MemoryStoreDeleteOnClosedStore_Errors(t *testing.T) {
	test := &Test{
		SupportedStoreTypes: []state.StoreType{
			state.MemoryStoreType,
		},
		Actions: []action.Action{
			action.Close(),
			action.DeleteE([]byte("does not matter"), "datastore closed"),
		},
	}

	test.Execute(t)
}

func TestDeleteClose_BadgerStoreDeleteOnClosedStore_Errors(t *testing.T) {
	test := &Test{
		SupportedStoreTypes: []state.StoreType{
			state.BadgerStoreType,
		},
		Actions: []action.Action{
			action.Close(),
			action.DeleteE([]byte("does not matter"), "Writes are blocked, possibly due to DropAll or Close"),
		},
	}

	test.Execute(t)
}
