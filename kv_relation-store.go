// Copyright (c) 2023 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package kv

import (
	"context"
)

type Bytesy interface {
	~[]byte | ~string
}

/*
TODO:
StreamIDs(ctx context.Context, ch chan<- string) error
// StreamRelatedIDs returns all RelatedIDs
StreamRelatedIDs(ctx context.Context, ch chan<- string) error
*/
type RelationStoreAdder[KEY Bytesy, RELATED_KEY Bytesy] interface {
	Add(ctx context.Context, key KEY, relatedKeys ...RELATED_KEY) error
}

type RelationStoreRemover[KEY Bytesy, RELATED_KEY Bytesy] interface {
	Delete(ctx context.Context, key KEY) error
	Remove(ctx context.Context, key KEY, relatedKeys ...RELATED_KEY) error
}

type RelationStoreGetter[KEY Bytesy, RELATED_KEY Bytesy] interface {
	RelatedKeys(ctx context.Context, key KEY) error
	Keys(ctx context.Context, relatedKey RELATED_KEY) error
}

type RelationStore[KEY Bytesy, RELATED_KEY Bytesy] interface {
	RelationStoreAdder[KEY, RELATED_KEY]
	RelationStoreRemover[KEY, RELATED_KEY]
	RelationStoreGetter[KEY, RELATED_KEY]
}

func NewRelationStore[KEY Bytesy, RELATED_KEY Bytesy](db DB, bucketNameBase string) RelationStore[KEY, RELATED_KEY] {
	return &relationStore[KEY, RELATED_KEY]{
		db,
		//NewRelationStoreTx[KEY, RELATED_KEY](bucketName)
	}
}

type relationStore[KEY Bytesy, RELATED_KEY Bytesy] struct {
	db DB
	//store RelationStoreTx
}
