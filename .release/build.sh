#!/usr/bin/env bash
if [ "$1" = "test" ]; then
    tag="v0.0.0"
    echo "test mode: building current branch as v0.0.0"
else
    workingBranch=$(git rev-parse --abbrev-ref HEAD)

    git diff --exit-code > /dev/null
    if [ $? -ne 0 ]; then
        echo "git workspace is not clean, commit or stash your changes"
        exit 1
    fi

    git fetch --tags
    tag=$(git describe --tags "$(git rev-list --tags --max-count=1)")
    if [ -z "${tag}" ]; then
        echo "could not find latest tag"
        exit 1
    fi

    git checkout "${tag}"
fi

target_folder=".release"
package="github.com/Oppodelldog/droxy"
ldflags=-ldflags="-X github.com/Oppodelldog/droxy/version.Number=${tag}"

package_split=("${package//\// }")
package_name="${package_split[-1]}"

platforms=("linux/amd64" "windows/amd64" "windows/386" "linux/arm/7")

for platform in "${platforms[@]}"
do
    platform_split=("${platform//\// }")
    GOOS="${platform_split[0]}"
    GOARCH="${platform_split[1]}"
    GOARM=""

    if [ "${GOARCH}" = "arm" ]; then
        GOARM=${platform_split[2]}

    fi

    output_folder="${GOOS}-${GOARCH}${GOARM}"
    output_name=${package_name}
    if [ "${GOOS}" = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS="${GOOS}" GOARCH="${GOARCH}" GOARM="${GOARM}" go build "${ldflags}" -o "${target_folder}/${tag}/${output_folder}/${output_name}" ${package}
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi

    # compress binary and remove builds
    currentWd=$(pwd)
    cp LICENSE "${target_folder}/${tag}/${output_folder}/LICENSE"
    cd "${target_folder}/${tag}/${output_folder}" || exit

    tar -cvzf "../${output_name}-${output_folder}.tar.gz" "${output_name}" LICENSE

    cd "${currentWd}" || exit
    rm -rf "${target_folder:?}/${tag}/${output_folder}"
done

if [ "$1" = "test" ]; then
    echo "test release done"
else
    git checkout "${workingBranch}"
fi