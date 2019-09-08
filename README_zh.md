[EN](README.md) | 中文

支持跨平台多工具链交叉编译的包管理器

## 快速开始

安装软件包，并更新 package.yaml 文件

```shell
xcpm install <name> <version> [-t <toolchain>]
```

更新 package.yaml 定义的依赖软件包

```shell
xcpm update [-t <toolchain>]
```

下载软件包源码

```shell
xcpm download <name> <version>
```

编译软件源码

```shell
xcpm build <name> [-t <toolchain>]
```
