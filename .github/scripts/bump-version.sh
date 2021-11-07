#!/bin/bash

set -e

version=$1
tag=${2:-none} # give default so it's not required to pass in

# configure git repositories
cd "$repoDirectory"
git config user.name "$GITHUB_ACTOR"
git config user.email "$GITHUB_ACTOR@users.noreply.github.com"

echo "Github actor: $GITHUB_ACTOR@users.noreply.github.com"

# if it's a release, then checkout the specified tag
if [[ $version == "release" ]]
then
  echo "checking out tag $tag for release"
  git checkout tags/"$tag" -b release-prod
fi

# bump the version
bumpversion "$version" --verbose

git push --tags
git push origin main
