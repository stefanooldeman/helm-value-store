package zookeeper_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/skuid/helm-value-store/internal/zookeeper"
	"github.com/skuid/helm-value-store/store"
)

var releaseStoreType = reflect.TypeOf(new(store.ReleaseStore)).Elem()

func TestZookeeperStore(t *testing.T) {

	zk, err := zookeeper.NewReleaseStore("localhost:2181")

	zkType := reflect.TypeOf(zk)

	if err != nil {
		fmt.Errorf("zookeeper.NewReleaseStore gave unexpected error: %v", err)
	}

	if !zkType.Implements(releaseStoreType) {
		fmt.Errorf("NewReleaseStore does not implement the ReleaseStore interface. Type info: %v", t)
	}

}
