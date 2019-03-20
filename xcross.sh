docker run --rm -it -v "$PWD":/go/src/awesomegostars -w /go/src/awesomegostars dockercore/golang-cross:1.11.0 sh -c '
    for GOOS in darwin linux windows; do
      for GOARCH in 386 amd64; do
        echo "Building $GOOS-$GOARCH"
        export GOOS=$GOOS
        export GOARCH=$GOARCH
        go build -o awesomegostars-$GOOS-$GOARCH
      done
    done
    '
