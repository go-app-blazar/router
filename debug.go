package router

// Debug toggles debug logging for this package.
func Debug(debug bool) {
	debugLayoutWrapper = debug
	debugRegister = debug
	debugRouterView = debug
}

// debugLayoutWrapper toggles debug logging for LayoutWrapper.
var debugLayoutWrapper = false

// DebugLayoutWrapper toggles debug logging for LayoutWrapper.
func DebugLayoutWrapper(debug bool) {
	debugLayoutWrapper = debug
}

// debugRegister toggles debug logging for Register.
var debugRegister = false

// DebugRegister toggles debug logging for Register.
func DebugRegister(debug bool) {
	debugRegister = debug
}

// debugRouterView toggles debug logging for RouterView.
var debugRouterView = false

// DebugRouterView toggles debug logging for RouterView.
func DebugRouterView(debug bool) {
	debugRouterView = debug
}
