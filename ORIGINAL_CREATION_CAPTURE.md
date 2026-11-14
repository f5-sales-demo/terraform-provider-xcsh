# Original creation capture

This opt-in Linux producer records original Secure Mesh Site v2 and Token creation
identities during Terraform apply. It does not establish ownership, registration,
guest access, health, routing or traffic acceptance by itself.

The owning executor prepares a private scope before starting the original provider.
Set `XCSH_ORIGINAL_CREATION_SCOPE_FILE` to its absolute path. Leaving it unset
preserves ordinary provider behavior on every platform. Enabled capture currently
requires Linux process evidence; other platforms reject it before submission.

The scope uses the following versioned fields:

| Field | Meaning |
| --- | --- |
| `schema_version` | Integer `1` |
| `scope_sha256` | Lowercase SHA-256 digest of the executor's reviewed operation |
| `directory` | Absolute, existing receipt directory |
| `provider_executable_sha256` | Independently verified digest of the original provider executable |
| `resources` | Exact target inventory of `type_name`, `namespace` and `name` |

Participating types are `xcsh_securemesh_site_v2` and `xcsh_token`. Target names and
namespaces must be nonempty lowercase DNS-style identifiers. The scope admits at
most 128 unique targets and is bounded to 64 KiB. Duplicate JSON keys, unknown
fields, trailing values, duplicate targets and unsupported versions fail admission.

The scope must be a regular file owned by the provider user with mode `0400` and
one hard link. The receipt directory must belong to that user with mode `0700`.
Paths must be absolute and canonical; symlink components are rejected. Capture
retains the original directory descriptor and checks the scope bytes, file identities,
permissions and executable identity again during persistence.

Before POST, capture exclusively creates and synchronizes a target-specific
`claim-<digest>.json`. The claim remains after success, error, interruption or
descriptor closure. A second invocation cannot reserve the same target in that
directory. An unresolved claim requires independent reconciliation; never delete
it to permit an unchanged retry.

After a successful POST, the provider first retains its existing resource ID and
creation-identity state. It then exclusively creates a random
`identity-<digest>.json` reservation and then a random
`receipt-<nonce>.json`, changes it to mode `0400`, synchronizes the file and directory,
and verifies the original bindings. Follow-up reads occur only after this checkpoint.
If receipt persistence fails or the POST has no immutable UID, an explicit diagnostic
stops the operation while retaining the successful creation state for owned cleanup.

Receipts contain only the version, reviewed operation and exact scope-file digests,
target projection, hashed original UID, random nonce and original process identity.
Process identity includes PID, parent PID, process start ticks, hashed boot ID,
executable digest and executable device/inode. No API credential, JWT, bootstrap
material, cloud-init or raw Terraform state is exported. Incomplete files are failed
captures; a consumer must reject truncation and malformed output.

A PID in JSON is self-reported evidence. Consumer admission must independently
observe the original provider child, parent relationship, Linux process start and
boot identity, executable digest/inode and the retained original sink descriptor.
It must match the expected private scope bytes and target inventory, verify exactly
one complete receipt per original claim and immutable identity reservation, and reject replaced identities. Terminal
applied-state corroboration remains a separate check. Neither a receipt nor later
read-back grants a renewed observation budget or authorizes another mutation.

The marketplace consumer remains held until this producer completes normal reviewed
source delivery, generated publication and downloaded artifact verification, followed
by independent installed and live acceptance.

## Version 2 bootstrap release

Version 1 continues to record creation only. Version 2 adds a
`bootstrap_authorization` object with integer `schema_version: 1`, an absolute
private `directory`, a canonical unpadded base64 Ed25519 `public_key`, and
`original_start_ms`/`original_deadline_ms` from the owning Linux monotonic clock.
The original operation is exactly two hours. The authorization directory must be
separate from the receipt directory, with no nesting between them.

After successful Token creation and its immutable receipt, the provider retains
created identity and a private held marker. It withholds computed token credentials
until the owning consumer finishes its original Platform inspection and durable
single-use checkout. The consumer writes `grant-<nonce>.json` with mode `0400`
and one hard link in the authorization directory.

The grant envelope has exactly `payload` and `signature` fields. The signature
is canonical unpadded base64 Ed25519 over the exact UTF-8 payload bytes prefixed
with `XCSH_BOOTSTRAP_RELEASE_V1\n`. The payload contains:

| Field | Required binding |
| --- | --- |
| `schema_version`, `kind` | Integer `1`, `original-bootstrap-release` |
| `scope_sha256`, `scope_file_sha256` | Exact original receipt values |
| `target`, `process`, `nonce` | Exact original receipt objects and nonce |
| `creation_identity_sha256` | Exact original created UID digest |
| `content_sha256` | SHA-256 of the original token content, without exporting it |
| `issued_monotonic_ms` | Original operation time; no more than five minutes old |
| `expires_at_epoch_ms` | Future token expiry independently verified by the issuer |

Unknown, duplicate, missing, malformed or differently cased fields refuse.
The provider checks the original scope, process, file and directory bindings,
grant signature, content digest and both clocks before release. It exclusively
creates `used-<nonce>.json`, flushes that file and directory, and rechecks the
original bindings before returning credentials. Cancellation, expiry, failed
persistence or an ambiguous marker never permits another release.

The private held marker keeps credentials unavailable during a later refresh
after failed authorization. Created resource identity remains available for
owned cleanup without registration or health. Legacy resources without this
private marker retain their prior read behavior. A released token can reconcile
the original unchanged configuration; a new VM bootstrap still needs the owning
consumer's reviewed plan and fresh, node-bound single-use material.

Neither a grant nor a receipt alone proves guest, routing, traffic, failover or
installation acceptance. Marketplace consumer enablement remains held until the
new producer's generated publication, downloaded artifact, installed runtime and
live gates are independently qualified.
