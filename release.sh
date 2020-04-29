#!/bin/bash
if [[ $(git diff --stat) != '' ]]; then
    echo "Git is dirty"
    #exit 1
fi
#git co master
#git pull origin master
if [[ $(git diff --stat) != '' ]]; then
    echo "Git is dirty"
    #exit 1
fi
VERSION_CHANGELOG=$(head -1 CHANGELOG.md | sed -n "s/^##\s*\(\S*\).*$/\1/p")
VERSION_GIT=$(git describe | sed -n "s/^v\([0-9.]*\)-.*$/\1/p")
LAST_VERSION=$(cat VERSION)
if [[ $VERSION_CHANGELOG != $VERSION_GIT || $VERSION_GIT != $LAST_VERSION ]]; then
    echo "Version mismatch. Git, Changelog and VERSION have to match"
    exit 1
fi
if [[ $VERSION == '' ]]; then
    MAJMIN=$(echo $LAST_VERSION | cut -d'.' -f1,2)
    PATCH=$(echo $LAST_VERSION | cut -d'.' -f3)
    NEXT=$((PATCH+1))
    VERSION="$MAJMIN.$NEXT"
fi
# TODO: ask to proceeed
echo $VERSION
git log "v${LAST_VERSION}..HEAD" --pretty=full --grep=RELEASE_NOTES
CHANGELOG=$(sed -n "/## ${VERSION}/,/##/{p}" CHANGELOG.md | head --lines=-1)
echo $CHANGELOG
