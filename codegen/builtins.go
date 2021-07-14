package codegen

import (
	"github.com/llir/llvm/ir"
	llir "github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func (state *State) MakeBuiltin(s string, in types.Type, out types.Type) *llir.Func {
	state.fns[s] = state.module.NewFunc(s, out, llir.NewParam(param, in))
	return state.fns[s]
}

func (state *State) BuiltinDouble() {
	fn := state.MakeBuiltin("double", int_t, int_t)
	block := fn.NewBlock("entry")
	block.NewRet(block.NewMul(constant.NewInt(&int_bt, 2), fn.Params[0]))
}

func (state *State) BuiltinPuts() {
	state.MakeBuiltin("puts", types.I8Ptr, types.I32)
}

func (state *State) BuiltinCalloc() {
	state.fns["calloc"] = state.module.NewFunc(
		"calloc",
		types.I8Ptr,
		ir.NewParam("size", types.I64),
		ir.NewParam("count", types.I64))
}
