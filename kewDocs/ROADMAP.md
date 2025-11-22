**Golang Learning Roadmap**

**1. Fundamentals**
- Install Go & set up environment (GOROOT, GOPATH)
- Basic syntax: variables, constants, functions, types
- Control structures: if/else, switch, for
- Arrays, slices, maps, structs
- Pointers & value/reference semantics

**2. Core Concepts**
- Packages & module management (`go mod`)
- Error handling idioms (`error`, wrapping, custom errors)
- Interfaces, type embedding
- Methods & receivers
- Concurrency: goroutines, channels, select, sync primitives

**3. Standard Library**
- File I/O (`os`, `io`, `bufio`)
- Networking (`net`, `net/http`)
- JSON, XML parsing (`encoding/json`)
- Contexts (`context.Context`)
- Time, date, random, sorting, regexp

**4. Testing & Tooling**
- Unit testing (`testing` package)
- Benchmarking, profiling (`pprof`)
- Static analysis: `go vet`, `golint`, `staticcheck`
- Dependency management: `go get`, `replace`, `tidy`

**5. Applied Projects**
- CLI tools (cobra, urfave/cli)
- REST APIs (standard `net/http`, gorilla/mux, chi)
- Database usage (`database/sql`, GORM, sqlx)
- Websockets, gRPC basics

**6. Advanced Topics**
- Generics
- Unsafe, reflection (`reflect`)
- Build tags, custom build pipelines
- Advanced concurrency: worker pools, context cancellation
- FFI, calling C code

**7. Ecosystem**
- Popular frameworks: Gin, Echo (API); Viper (config)
- Go modules, vendoring
- Deployment: Docker, cross-compiling, architecture-specific builds

**8. Patterns & Best Practices**
- Idiomatic Go (`Effective Go`, Go Proverbs)
- Organizing large codebases
- Error handling patterns, logging (zap, zerolog)

**9. Real-World: OSS Contribution**
- Read & debug others' Go code (e.g., Docker, kubernetes)
- Make PRs, review issues

**10. Stay Up-to-Date**
- Follow Go blogs/news
- Attend Go conferences/meetups (virtual or in-person)

**Non-Obvious Tips**
- Write/contribute to Go linters or code generators
- Explore custom Go runtime tweaks (speculative, for advanced users)
- Experiment with WebAssembly targets (`GOARCH=wasm`)
- Dive into Go compiler internals

Minimal viable order: fundamentals → core language → stdlib → concurrency → applied projects → larger codebases/OSS → advanced performance.