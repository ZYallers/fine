package f

import (
	"context"
	"io"

	"github.com/ZYallers/fine/container/fvar"
	"github.com/ZYallers/fine/internal/empty"
	"github.com/ZYallers/fine/util/futil"
)

// Go creates a new asynchronous goroutine function with specified recover function.
//
// The parameter `recoverFunc` is called when any panic during executing of `goroutineFunc`.
// If `recoverFunc` is given nil, it ignores the panic from `goroutineFunc` and no panic will
// throw to parent goroutine.
//
// But, note that, if `recoverFunc` also throws panic, such panic will be thrown to parent goroutine.
func Go(
	ctx context.Context,
	goroutineFunc func(ctx context.Context),
	recoverFunc func(ctx context.Context, exception error),
) {
	futil.Go(ctx, goroutineFunc, recoverFunc)
}

// NewVar returns a gvar.Var.
func NewVar(i interface{}, safe ...bool) *Var {
	return fvar.New(i, safe...)
}

// Dump dumps a variable to stdout with more manually readable.
func Dump(values ...interface{}) {
	futil.Dump(values...)
}

// DumpTo writes variables `values` as a string in to `writer` with more manually readable
func DumpTo(writer io.Writer, value interface{}, option futil.DumpOption) {
	futil.DumpTo(writer, value, option)
}

// DumpWithType acts like Dump, but with type information.
// Also see Dump.
func DumpWithType(values ...interface{}) {
	futil.DumpWithType(values...)
}

// DumpWithOption returns variables `values` as a string with more manually readable.
func DumpWithOption(value interface{}, option futil.DumpOption) {
	futil.DumpWithOption(value, option)
}

// DumpJson pretty dumps json content to stdout.
func DumpJson(value interface{}) {
	futil.DumpJson(value)
}

// Throw throws an exception, which can be caught by TryCatch function.
func Throw(exception interface{}) {
	futil.Throw(exception)
}

// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.
func Try(ctx context.Context, try func(ctx context.Context)) (err error) {
	return futil.Try(ctx, try)
}

// TryCatch implements try...catch... logistics using internal panic...recover.
// It automatically calls function `catch` if any exception occurs and passes the exception as an error.
//
// But, note that, if function `catch` also throws panic, the current goroutine will panic.
func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error)) {
	futil.TryCatch(ctx, try, catch)
}

// IsNil checks whether given `value` is nil.
// Parameter `traceSource` is used for tracing to the source variable if given `value` is type
// of pointer that also points to a pointer. It returns nil if the source is nil when `traceSource`
// is true.
// Note that it might use reflect feature which affects performance a little.
func IsNil(value interface{}, traceSource ...bool) bool {
	return empty.IsNil(value, traceSource...)
}

// IsEmpty checks whether given `value` empty.
// It returns true if `value` is in: 0, nil, false, "", len(slice/map/chan) == 0.
// Or else it returns true.
//
// The parameter `traceSource` is used for tracing to the source variable if given `value` is type of pointer
// that also points to a pointer. It returns true if the source is empty when `traceSource` is true.
// Note that it might use reflect feature which affects performance a little.
func IsEmpty(value interface{}, traceSource ...bool) bool {
	return empty.IsEmpty(value, traceSource...)
}
