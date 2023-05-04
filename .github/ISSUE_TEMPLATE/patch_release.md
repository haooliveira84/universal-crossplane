---
name: Patch Release
about: Cut a UXP patch release
labels: release
---

<!--
Issue title should be in the following format:

    Cut vX.Y.Z-up.K Release on DATE

For example:

    Cut v1.3.1-up.1 on June 29, 2021.

Please assign the release manager to the issue.
-->

This issue can be closed when we have completed the following steps (in order).
Please ensure all artifacts (PRs, workflow runs, Tweets, etc) are linked from
this issue for posterity. Assuming `vX.Y.Z-up.K` is being cut, after upstream
[crossplane/crossplane][upstream-xp] `vX.Y.Z` has been released
according to the declared schedule, you should have:

- [ ] Synced the `release-vX.Y` release branch in [upbound/crossplane][upbound-xp-fork], with upstream [crossplane/crossplane][upstream-xp] release branch, up to the `vX.Y.Z` tag, adding any required change specific to the fork, see [here][sync-xp-fork] for more details.
- [ ] Tagged [upbound/crossplane][upbound-xp-fork] `vX.Y.Z-up.K` from the `release-X.Y` branch by:
  - [ ] Running the [Tag workflow][tag-xp-fork] on the `release-vX.Y` branch with the proper release version, `vX.Y.Z-up.K`. Use `Release vX.Y.Z-up.K` as message (FYI: the format suggested is only for consistency, there is no actual dependency on it).
  - [ ] Running the [CI workflow][ci-xp-fork] on the `release-vX.Y` branch to build and publish the latest tagged artifacts.
  - [ ] Confirm the image has been successfully published using: `docker buildx imagetools inspect docker.io/upbound/crossplane:vX.Y.Z-up.K`
- [ ] Created and merged a PR for [upbound/universal-crossplane][uxp] to either the `main` branch, if cutting a patch for the latest supported release, **taking care to label it `backport release-X.Y`**, or directly to the `release-X.Y` branch, if cutting a patch for an older supported release. With the following changes:
  - [ ] Update any reference to the old latest release to `vX.Y.Z-up.K`, such as `CROSSPLANE_TAG` and `CROSSPLANE_COMMIT` in the `Makefile`.
  - [ ] Run `make generate` to import any changes in the [upstream Helm chart][upstream-helm-chart].
- [ ] Cut [UXP][uxp] `vX.Y.Z-up.K` release from the `release-X.Y` branch by:
  - [ ] Running the [Tag workflow][tag-uxp] on the `release-vX.Y` branch with the proper release version, `vX.Y.Z-up.K`. Use `Release vX.Y.Z-up.K` as message (FYI: the format suggested is only for consistency, there is no actual dependency on it).
  - [ ] Running the [CI workflow][ci-uxp] on the `release-vX.Y` branch to build and publish the latest tagged artifacts.
- [ ] Run the [Promote workflow][promote-uxp] on the `release-vX.Y` branch to promote `vX.Y.Z-up.K` to [stable][uxp-stable-channel], it should contain `universal-crossplane-X.Y.Z-up.K.tgz`. Verify everything is correctly working by running `up uxp install` against an empty Kubernetes cluster, e.g. `kind create cluster`, which should result in an healthy UXP installation with expected image versions.
- [ ] Created and published well authored release notes for [UXP][uxp-releases] `vX.Y.Z-up.K`. See the previous release for an example, these should at least:
  - [ ] enumerate relevant updates that were merged in [u/xp][upbound-xp-fork] and [u/uxp][uxp].
  - [ ] mention the [xp/xp][upstream-xp] version it refers to.
  - [ ] list new contributors to [u/uxp][uxp].
  - [ ] have the links to the full changelog of [u/xp][upbound-xp-fork] and [u/uxp][uxp].
- [ ] Ensured that users have been notified of the release on all communitcation channels:
  - [ ] Slack: crossposting on Crossplane's Slack workspace channels `#announcements`, `#upbound` and `#squad-crossplane` on Upbound's Slack.
  - [ ] Twitter: ask `#marketing` on Upbound's Slack to do so.


<!-- Named Links -->
[ci-uxp]: https://github.com/upbound/universal-crossplane/actions/workflows/ci.yml
[ci-xp-fork]: https://github.com/upbound/crossplane/actions/workflows/ci.yml
[promote-uxp]: https://github.com/upbound/universal-crossplane/actions/workflows/promote.yml
[sync-xp-fork]: https://github.com/upbound/universal-crossplane/blob/main/CONTRIBUTING.md#crossplane-fork-sync
[tag-uxp]: https://github.com/upbound/universal-crossplane/actions/workflows/tag.yml
[tag-xp-fork]: https://github.com/upbound/crossplane/actions/workflows/tag.yml
[upbound-xp-fork]: https://github.com/upbound/crossplane
[upstream-helm-chart]: https://github.com/crossplane/crossplane/tree/master/cluster/charts/crossplane
[upstream-xp-values]: https://github.com/crossplane/crossplane/blob/master/cluster/charts/crossplane/values.yaml.tmpl
[upstream-xp]: https://github.com/crossplane/crossplane
[uxp-main-channel]: https://charts.upbound.io/main
[uxp-releases]: https://github.com/upbound/universal-crossplane/releases
[uxp-schedule]: https://github.com/upbound/universal-crossplane/blob/main/README.md#releases
[uxp-stable-channel]: https://charts.upbound.io/stable
[uxp-values]: https://github.com/upbound/universal-crossplane/blob/main/cluster/charts/universal-crossplane/values.yaml.tmpl
[uxp]: https://github.com/upbound/universal-crossplane
