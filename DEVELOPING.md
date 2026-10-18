# Provider development

## Latest upstream specifications

This prerelease provider consumes the latest published API specifications.
Run `make download-specs` before generation. The command resolves the latest
release, verifies its publication receipt and every downloaded asset, and records
the source used by this build.

`tools/spec-version.txt` and `tools/spec-release.json` record the resolved source.
They provide reproducible generation and publication evidence; they do not select
an older release for future updates. Historical delivery ledgers remain evidence
of completed publications.

Sync OpenAPI Specs resolves latest on every upstream notification, manual run,
and hourly reconciliation. A newer release supersedes an unpublished older
delivery. An identical pending delivery resumes, and a completed delivery is a
no-op. Invalid receipts and conflicting identities fail.

## Generate and verify

```sh
make download-specs
make generate
make build
make test
make lint
```

Read `CONTRIBUTING.md` for the issue, worktree, review, and cleanup workflow.
Source PRs carry authored changes and verified source receipts. On Merge owns
the generated provider, documentation, and regeneration receipt, then publishes
and records the completed delivery.

The network allowlist check binds reviewed content independently of its release
label. New releases containing identical reviewed content flow automatically;
changed content requires review. Keep downloaded source checksums, contract
validation, generated content ownership, and publication checks effective.
