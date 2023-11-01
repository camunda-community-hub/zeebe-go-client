# Zeebe Go Client

## Development

If we had a gateway-protocol change we need to make sure that we regenerate the protobuf file, which is used by the go client.
In order to do this please follow [this guide](../../gateway-protocol-impl/README.md).

## Testing

### gRPC Mock

To regenerate the gateway mock `internal/mock_pb/mock_gateway.go` run [`mockgen`](https://github.com/golang/mock#installation):

```
GO111MODULE=off mockgen github.com/camunda/zeebe/clients/go/v8/pkg/pb GatewayClient,Gateway_ActivateJobsClient > internal/mock_pb/mock_gateway.go
```

If you see errors regarding packages which are not found then as alternative you can try this:

```
GO111MODULE=off mockgen -source $GOPATH/src/github.com/camunda/zeebe/clients/go/v8/pkg/pb/gateway.pb.go GatewayClient,Gateway_ActivateJobsClient > internal/mock_pb/mock_gateway.go
```

### Integration tests

To run the integration tests, a Docker image for Zeebe must be built with the tag 'current-test'. To do that you can run (in the camunda/zeebe dir):

```
docker build --build-arg DISTBALL=dist/target/camunda-zeebe*.tar.gz -t camunda/zeebe:current-test --target app .
```

To add new zbctl tests, you must generate a golden file with the expected output of the command you are testing. The tests ignore numbers so you can leave any keys or timestamps in your golden file, even though these will most likely be different from test command's output. However, non-numeric variables are not ignored. For instance, the help menu contains:

```
--clientCache string    Specify the path to use for the OAuth credentials cache. If omitted, will read from the environment variable 'ZEEBE_CLIENT_CONFIG_PATH' (default "YOUR_HOME/.camunda/credentials")
```

To make them host-independent, the tests replace the HOME environment variable with `/tmp` which means you must do the same in your golden file.

## Dependencies

After making changes to the Go client, you can vendor the new dependencies with:

```
go mod vendor
```

This command will also remove or download dependencies as needed. To do that without vendoring them, you can run `go mod tidy`.
