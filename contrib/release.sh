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

echo "Releasing ${VERSION}"
read -n 1 -p "Do you want to proceed?"

CHANGES=$(git log "v1.8.6..HEAD" --pretty=full --grep=RELEASE_NOTES | grep RELEASE_NOTES | cut -d'=' -f2 | grep -v n/a | sort)
echo $CHANGES
# TODO prepend CHANGES to CHANGELOG.md
CHANGELOG=$(sed -n "/## ${VERSION}/,/##/{p}" CHANGELOG.md | head --lines=-1)
echo $CHANGELOG

echo "${VERSION}" > VERSION

# TODO git commit -am"Tag v${VERSION}"
# TODO git tag -s v${VERSION}
# TODO make completion
# TODO goreleaser --skip-publish
# TODO git push origin v${VERSION}
# TODO goreleaser --release-notes <(contrib/relnotes.sh)
