// Package api provides utility functions for making API calls to AI providers.
//
// This file contains helper functions for making API calls with context support
// and handling timeouts and cancellation gracefully.
package api

import (
	"context"
	"reflect"
)

// ApiCallWithContext executes an API call with context support for timeout and cancellation.
//
// Parameters:
//   - ctx: Context for timeout and cancellation
//   - apiFunc: The API function to call (must be a callable)
//   - args: Variable number of arguments to pass to the API function
//
// Returns:
//   - interface{}: The result of the API call
//   - error: Any error that occurred during the call
//
// The function uses reflection to handle different API function signatures and
// provides context-aware execution with proper cleanup on cancellation.
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
