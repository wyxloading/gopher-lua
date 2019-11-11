package lua

func OpenBaseStrip(funcNames []string) func(*LState) int {
	return func(L *LState) int {
		if len(funcNames) < 1 {
			return 0
		}
		funcs := make(map[string]LGFunction)
		for _, funcName := range funcNames {
			if f, ok := baseFuncs[funcName]; ok {
				funcs[funcName] = f
			}
		}
		if len(funcs) < 1 {
			return 0
		}
		global := L.Get(GlobalsIndex).(*LTable)
		L.SetGlobal("_G", global)
		L.SetGlobal("_VERSION", LString(LuaVersion))
		L.SetGlobal("_GOPHER_LUA_VERSION", LString(PackageName+" "+PackageVersion))
		basemod := L.RegisterModule("_G", funcs)
		global.RawSetString("ipairs", L.NewClosure(baseIpairs, L.NewFunction(ipairsaux)))
		global.RawSetString("pairs", L.NewClosure(basePairs, L.NewFunction(pairsaux)))
		L.Push(basemod)
		return 1
	}
}
