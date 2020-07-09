// Copyright © VNG Realisatie 2020
// Licensed under the EUPL

package database

import (
	"context"
	"errors"
	"fmt"
	"path"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type AccessRequest struct {
	ID               string              `json:"id,omitempty"`
	OrganizationName string              `json:"organizationName,omitempty"`
	ServiceName      string              `json:"serviceName,omitempty"`
	Status           AccessRequestStatus `json:"status,omitempty"`
	CreatedAt        time.Time           `json:"createdAt,omitempty"`
	UpdatedAt        time.Time           `json:"updatedAt,omitempty"`
}

type AccessRequestStatus int

const (
	AccessRequestFailed AccessRequestStatus = iota
	AccessRequestCreated
	AccessRequestReceived
)

var ErrActiveAccessRequest = errors.New("already active access request")

func (db ETCDConfigDatabase) ListAllOutgoingAccessRequests(ctx context.Context) ([]*AccessRequest, error) {
	key := path.Join("access-requests", "outgoing")

	r := []*AccessRequest{}

	err := db.list(ctx, key, &r, clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend))
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db ETCDConfigDatabase) ListOutgoingAccessRequests(ctx context.Context, organizationName, serviceName string) ([]*AccessRequest, error) {
	key := path.Join("access-requests", "outgoing", organizationName, serviceName)

	r := []*AccessRequest{}

	err := db.list(ctx, key, &r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db ETCDConfigDatabase) CreateAccessRequest(ctx context.Context, accessRequest *AccessRequest) (*AccessRequest, error) {
	existing, err := db.GetLatestOutgoingAccessRequest(ctx, accessRequest.OrganizationName, accessRequest.ServiceName)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, ErrActiveAccessRequest
	}

	t := db.clock.Now()
	id := fmt.Sprintf("%x", t.UnixNano())

	key := path.Join("access-requests", "outgoing", accessRequest.OrganizationName, accessRequest.ServiceName, id)

	accessRequest.ID = id
	accessRequest.Status = AccessRequestCreated
	accessRequest.CreatedAt = t
	accessRequest.UpdatedAt = t

	err = db.put(ctx, key, accessRequest)
	if err != nil {
		return nil, err
	}

	return accessRequest, nil
}

func (db ETCDConfigDatabase) GetLatestOutgoingAccessRequest(ctx context.Context, organizationName, serviceName string) (*AccessRequest, error) {
	var r *AccessRequest

	key := path.Join("access-requests", "outgoing", organizationName, serviceName)

	err := db.get(ctx, key, &r, clientv3.WithLastKey()...)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (db ETCDConfigDatabase) ListAllLatestOutgoingAccessRequests(ctx context.Context) (map[string]*AccessRequest, error) {
	accessRequests, err := db.ListAllOutgoingAccessRequests(ctx)
	if err != nil {
		return nil, err
	}

	latestAccessRequests := make(map[string]*AccessRequest)

	for _, a := range accessRequests {
		key := path.Join(a.OrganizationName, a.ServiceName)
		if _, ok := latestAccessRequests[key]; !ok {
			latestAccessRequests[key] = a
		}
	}

	return latestAccessRequests, nil
}
