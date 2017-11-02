#/usr/bin/env bash

tag=$1
upload="github-release upload --user posener --repo eztables --tag ${tag}"

if [ -z ${tag} ]; then
	echo "usage: give git tag as a first argument" >&2
	exit 1
fi

if [ -z ${GITHUB_TOKEN} ]; then
	echo "usage: environment variable GITHUB_TOKEN must be exported" >&2
	exit 1
fi

mkdir -p build
rm -rf build/*

function build_and_upload() {
	name=$1
	echo $1 building...
	go build -o build/${name}
	echo $1 uploading...
	${upload} --name ${name} --file build/${name}
	echo $1 done...
}

for GOOS in darwin linux windows; do
	for GOARCH in 386 amd64; do
		build_and_upload eztables-$GOOS-$GOARCH &
	done
done

wait
