# GitHub automation security

## Canonical mutation credential

Repository automation has one credential for branch pushes, pull requests, and
events that must start downstream workflows: the `REPO_SYNC_TOKEN` repository
secret. There is no `GITHUB_TOKEN` fallback and no alternate legacy secret name.
Every mutating workflow checks that this secret is present before it performs
work.

GitHub intentionally prevents events created with a workflow's `GITHUB_TOKEN`
from starting most new workflow runs. The fine-grained personal access token is
therefore required for generated pull requests whose required checks must run.

The token is passed only to steps that need it. Workflows keep top-level
permissions empty and grant the automatic `GITHUB_TOKEN` job-level permissions
only for separate read-only or GitHub-release operations.

## Required access

`REPO_SYNC_TOKEN` is a fine-grained token restricted to this repository with:

- Contents: read and write
- Pull requests: read and write
- Actions: read and write when the automation operation requires workflow
  dispatch or workflow-file mutation

The repository secret must never be printed, copied into an artifact, committed,
or passed through an untrusted pull-request context. Credential lifecycle changes
are an external administrative operation; repository workflows never rotate or
revoke the token.

## Workflow contract

The canonical token is used by:

- `auto-merge.yml` for same-repository automatic merges;
- `discover-defaults.yml` for generated default-discovery pull requests;
- `on-merge.yml` for regeneration branches, pull requests, and delivery receipts;
- `sync-openapi.yml` for immutable specification delivery branches and pull
  requests;
- `enforce-repo-settings.yml` when invoking the fleet settings controller.

Fork pull requests never receive repository secrets. Same-repository jobs that
require downstream-triggering mutations fail closed when `REPO_SYNC_TOKEN` is
absent.

## Live API credentials

Live verification uses `XCSH_API_URL` and `XCSH_API_TOKEN`. The URL identifies
infrastructure and is treated as sensitive along with the token. Workflows may
check only whether these values are present; they must not print either value or
upload raw tenant output. Published evidence contains aggregate results with
failure payloads and infrastructure identifiers removed.

## Verification

Workflow contract tests enforce:

- immutable action and reusable-workflow references;
- the absence of obsolete alternate secret names and authentication fallbacks;
- required `REPO_SYNC_TOKEN` preflight failures;
- fork isolation and least-privilege permissions;
- sanitized live-tenant artifacts;
- fail-closed publication and regeneration result matrices.

Run the workflow contract tests, Actionlint, Zizmor, and ShellCheck after any
automation change.
