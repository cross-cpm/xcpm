[EN](README.md) | 中文

支持跨平台多工具链交叉编译的包管理器

## 快速开始

安装 package.yaml 定义的所有依赖软件包

```shell
xcpm install [-t <toolchain>]
```

安装单个依赖软件包

```shell
xcpm install <name> [-t <toolchain>]
```

## 常用命令

下载软件包源码

```shell
xcpm download <name> <version>
```

编译软件源码

```shell
xcpm build <name> [-t <toolchain>]
```

安装软件包，并更新 package.yaml 文件

```shell
xcpm install <name> <version> [-t <toolchain>]
```

更新软件包数据仓库

```shell
xcpm update
```

