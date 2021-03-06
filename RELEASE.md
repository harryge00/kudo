# Development and Release Process

## Development Processes

### Master

The HEAD of master is consider the most active development, to which, new features and bug fixes are applied. The expectation is that master will always build and pass tests. Development is incremental. It is expected that pull-requests are either complete, are not fully integrated into code execution path or have a feature switch which must be enabled until code completion is reached. HEAD on master is expected to be in a state that it could be released. The next major release will be tagged and released from master. If the current released version is `v0.3.1` and the next major release is `v0.4.0` when it is deemed appropriate `v0.4.0` will be tagged off master, followed by a release. From that tag work will continue on master. IF it is necessary in the future to have a patch release for `v0.4.1`, a branch will be created off the `v0.4.0` tag with the nomenclature of `releases\{major.minor}`, in this example `releases\0.4`. Branches with `releases/*` are protected branches in the repository.

### Bug Fixes

Bug fixes are expected to be worked on and applied to `master`.  If the fix is needed for a previous supported release version of KUDO, then a back port is expected. The bug pull request is expected to be marked with a `back port` label. It is easiest for the developer working on the bug to back port to a previous version. It is expected when possible that the back port is a `git cherry-pick` to a previous version, updating as necessary to conform to previous code and architecture. On occasion a cherry-pick is too much of a burden. In this case, fix the bug as a new pull request against the release branch. Provide the same title and details of the original pull request for traceability.

## Release Request Process

The KUDO Project is released on an as-needed basis. The process is as follows:

1. An issue is proposing a new release with a changelog since the last release
2. All [OWNERS](OWNERS) must LGTM this release
3. Release process below is followed
5. The release issue is closed

### Release Process

The official binaries for KUDO are created using [goreleaser](https://goreleaser.com/) for the release process through the circle-ci release job. The [.goreleaser.yml](.goreleaser.yml) defines the binaries which are supported for each release.  The circle-ci release process has not been provided with credentials for Github and Docker hub yet.  The current process to initiate a release is:

1. Ensure you have credential `GITHUB_TOKEN` set. The env must include `export GITHUB_TOKEN=<personal access token>`. [Help](https://help.github.com/en/articles/creating-a-personal-access-token-for-the-command-line) provided from Github. The token must grant full access to: repo, write:packages, read:packages.
2. Ensure you are logged into Docker hub and have rights to push to kudobuilder.
3. Tag repo with expected release `git tag -a v0.2.0 -m "v0.2.0"`  && push tag `git push --tags`.
4. Invoke goreleaser `goreleaser --rm-dist`.
5. Update the GH release with Release high-levels.
6. An announcement email is sent to `kudobuilder@googlegroups.com` with the subject `[ANNOUNCE] Kudo $VERSION is released`
7. Run `./hack/generate_krew.sh` and submit the generated `kudo.yaml` to https://github.com/kubernetes-sigs/krew-index/.

**Note:** If there are issues with the release, any changes to the repository will result in it being considered "dirty" and not in a state to be released.
It is possible outside of the standard release process to build a "snapshot" release using the following command: `goreleaser release --skip-publish --snapshot --rm-dist`
This process will create a "dist" folder with all the build artifacts. The changelog is not created unless a full release is executed. If you are looking to get a "similar" changelog, install [github-release-notes](https://github.com/buchanae/github-release-notes) and execute `github-release-notes -org kudobuilder -repo kudo -since-latest-release`.

### Cutting a Release Branch

As outlined above, when it is necessary to create a new release branch, it is necessary to update the [circle-ci config](https://github.com/kudobuilder/kudo/blob/master/.circle-ci/config.yml#L13) to test merges against the correct branch. It is necessary replace all references to `master` with the appropriate release branch.

### Cutting a Patch Release

When cutting a patch release, for example `v0.3.3`, it is necessary to ensure that all bugs fixed on master after `v0.3.2` have landed on the release branch, `releases/0.3` in this case.
