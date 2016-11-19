#!/bin/bash

function bump_version() {
	echo "$1" | awk -F '.' '{$3=$3+1; print $1 "." $2 "." $3 "-dev"}'
}

last_tag="$(git describe --tags --abbrev=0)"

if git describe --tags --exact-match &>/dev/null && git diff --exit-code &>/dev/null; then
	version="${last_tag}"
else
	version="$(bump_version ${last_tag})"
fi

echo "Building version ${version}"
go build -ldflags "-X main.version=${version}" -o dbtm
