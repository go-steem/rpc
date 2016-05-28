# rpc

Golang RPC client library for [Steem](https://steem.io).

## Package Organisation

Once you create a `Client` object, you can start calling its methods,
which correspond to the methods exported via `steemd`'s RPC endpoint.

There are two versions for every method. The regular method and the raw method.
The difference is that the raw method returns `*json.RawMessage`, so it is not
trying to unmarshall the response into the right object. The reason for this
distinction is to be able to start using the `rpc` package even though not all
methods are specified properly yet.

## Status

This package is still under rapid development and it is by no means complete.
There is no promise considering API stability yet.

For now there are raw methods specified for most of the RPC methods available.

## License

MIT, see the `LICENSE` file.
