module github.com/vividvilla/simplesessions/examples/fasthttp-redis

go 1.18

require (
	github.com/redis/go-redis/v9 v9.5.1
	github.com/valyala/fasthttp v1.54.0
	github.com/vividvilla/simplesessions/stores/redis/v3 v3.0.0
	github.com/vividvilla/simplesessions/v3 v3.0.0
)

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/klauspost/compress v1.17.7 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
)

replace (
	github.com/vividvilla/simplesessions/stores/redis/v3 => ../../stores/redis
	github.com/vividvilla/simplesessions/v3 => ../..
)
