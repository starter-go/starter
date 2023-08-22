# starter

这是 starter 框架的核心模块，提供自动配置、日志初始化等功能.


## 常用属性

| 名称                  | 用途                       | 备注                            |
| --------------------- | -------------------------- | ------------------------------- |
| debug.enabled         | 启用调试功能               |                                 |
| debug.log-properties  | 把属性打印到日志中         | 需要配置 "debug.enabled" 为true |
| debug.log-environment | 把 环境变量 打印到日志中   | 需要配置 "debug.enabled" 为true |
| debug.log-arguments   | 把 命令行参数 打印到日志中 | 需要配置 "debug.enabled" 为true |
