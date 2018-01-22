package loader

import (
	"gopkg.in/nicksrandall/dataloader.v5"
	"fmt"
	"golang.org/x/net/context"
)

type key string

const (
	userLoaderKey     key = "user"
)

// Initialize a lookup map of context keys to batch functions.
//
// When Attach is called on the Collection, the batch functions are used to create new dataloader
// instances. The instances are attached to the request context at the provided keys.
//
// The keys are then used to extract the dataloader instances from the request context.
func NewLoaderCollection() LoaderCollection {
	return LoaderCollection{
		dataloaderFuncMap: map[key]dataloader.BatchFunc{
			userLoaderKey:     newUserLoader(),
		},
	}
}

// Collection holds an internal lookup of initialized batch data load functions.
type LoaderCollection struct {
	dataloaderFuncMap map[key]dataloader.BatchFunc
}


func (c LoaderCollection) Attach(ctx context.Context) context.Context {
	for k, batchFn := range c.dataloaderFuncMap {
		ctx = context.WithValue(ctx, k, dataloader.NewBatchedLoader(batchFn))
	}

	return ctx
}

// extract is a helper function to make common get-value, assert-type, return-error-or-value
// operations easier.
func extract(ctx context.Context, k key) (*dataloader.Loader, error) {
	ldr, ok := ctx.Value(k).(*dataloader.Loader)
	if !ok {
		return nil, fmt.Errorf("unable to find %s loader on the request context", k)
	}

	return ldr, nil
}