<!-- [![Go Reference](https://pkg.go.dev/badge/github.com/grafana/xk6-kubernetes.svg)](https://github.com/saniyar-dev/xk6-net) -->
<!-- [![Version Badge](https://img.shields.io/github/v/release/grafana/xk6-kubernetes?style=flat-square)](https://github.com/grafana/xk6-kubernetes/releases) -->
<!-- ![Build Status](https://img.shields.io/github/actions/workflow/status/grafana/xk6-kubernetes/ci.yml?style=flat-square) -->

# xk6-net

A k6 extension that recreate HTTP requests, which you can make more complicated requests.
Phase 1 implementation of new-http api design for k6.
For more information see [this issue](https://github.com/grafana/k6/issues/3038) and [this one](https://github.com/grafana/k6/issues/2461)

## Build

To build a custom `k6` binary with this extension, first ensure you have the prerequisites:

- [Go toolchain](https://go101.org/article/go-toolchain.html)
- Git

1. Download [xk6](https://github.com/grafana/xk6):

   ```bash
   go install go.k6.io/xk6/cmd/xk6@latest
   ```

2. [Build the k6 binary](https://github.com/grafana/xk6#command-usage):

   ```bash
   xk6 build --with github.com/saniyar-dev/xk6-net
   ```

   The `xk6 build` command creates a k6 binary that includes the xk6-net extension in your local folder. This k6 binary can now run a k6 test using [xk6-net APIs](#apis).

### Development

To make development a little smoother, use the `Makefile` in the root folder. The default target will format your code<!-- , run tests, --> and create a `k6` binary with your local code rather than from GitHub.

```shell
git clone https://github.com/saniyar-dev/xk6-net.git
cd xk6-net
make
```

<!---->
<!-- Using the `k6` binary with `xk6-kubernetes`, run the k6 test as usual: -->
<!---->
<!-- ```bash -->
<!-- ./k6 run k8s-test-script.js -->
<!---->
<!-- ``` -->

# APIs

## Generic API

### Examples
