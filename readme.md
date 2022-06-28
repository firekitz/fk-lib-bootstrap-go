# dtx-lib-bootstrap-go

A library to increase the reliability of running and shutting down Go applications.

## Usage

```go
package main

import (
	"bitbucket.org/weltcorp/dtx-lib-bootstrap-go"
	"context"
)

func main() {
	ctx := context.Background()

	shutdowns := config.Setup(ctx)
	
	// Example
	go GrpcServer()
	go GatewayServer()

	// ...

	dtxBootstrap.GracefulShutdown(shutdowns)
}
```