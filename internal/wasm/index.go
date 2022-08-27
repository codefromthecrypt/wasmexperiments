package wasm

import "github.com/threadedstream/wasmexperiments/internal/pkg/reporter"

func (m *Module) GetFunction(i int) *Function {
	if i >= len(m.FunctionIndexSpace) || i < 0 {
		reporter.ReportError("module.GetFunction: attempting to use index out of bounds %d", i)
	}
	return nil
}
