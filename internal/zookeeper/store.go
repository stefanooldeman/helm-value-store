package zookeeper

import (
	"context"
	"time"

	"github.com/skuid/helm-value-store/store"

	"github.com/samuel/go-zookeeper/zk"
)

type ZookeeperStore struct {
	c zk.Connect
}

// NewReleaseStore creates a new ReleaseStore
func NewReleaseStore(zookeeperConnection string) (store.ReleaseStore, error) {

	c, _, err := zk.Connect([]string{"127.0.0.1"}, time.Second*10)
	if err != nil {
		return err
	}

	// TODO: defer zk.Close()
	return &ZookeeperStore{c}, nil
}

func (zs *ZookeeperStore) Delete(ctx context.Context, uniqueID string) error {
	return nil
}

func (zs *ZookeeperStore) Get(ctx context.Context, uniqueID string) (*store.Release, error) {
	return &store.Release{}, nil
}

func (zs *ZookeeperStore) Put(context.Context, store.Release) error {
	return nil
}

func (zs *ZookeeperStore) List(ctx context.Context, selector map[string]string) (store.Releases, error) {
	return store.Releases{}, nil
}

func (zs *ZookeeperStore) Load(context.Context, store.Releases) error {
	return nil
}

func (zs *ZookeeperStore) Setup(context.Context) error {
	_, err = zs.c.Create("/helmvaluestore/", "0", 0, gozk.WorldACL(gozk.PERM_ALL))
	if err != nil {
		println(err.String())
	} else {
		println("Created!")
	}

	return err
}
