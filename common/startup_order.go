package common

// 定义初始化顺序，数值越小越靠前
const (
	StartupOrder = iota - 1000

	StartupOrderLogger
	StartupOrderBanner
	StartupOrderDebug
)
