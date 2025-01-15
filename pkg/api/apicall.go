package api

import (
	"context"
	"reflect"
)

func ApiCallWithContext(ctx context.Context, apiFunc interface{}, args ...interface{}) (interface{}, error) {
	resultChan := make(chan interface{})
	errChan := make(chan error)

	go func() {
		funcValue := reflect.ValueOf(apiFunc)
		params := make([]reflect.Value, len(args))
		for i, arg := range args {
			params[i] = reflect.ValueOf(arg)
		}

		results := funcValue.Call(params)

		if len(results) > 1 && !results[1].IsNil() {
			errChan <- results[1].Interface().(error)
		} else {
			resultChan <- results[0].Interface()
		}
	}()

	select {
	case result := <-resultChan:
		return result, nil
	case err := <-errChan:
		return nil, err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
