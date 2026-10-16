# SMSv2 capability ledger

This ledger tracks incremental qualification under
[#2586](https://github.com/f5-sales-demo/terraform-provider-xcsh/issues/2586). Schema presence does not
establish lifecycle correctness or traffic enforcement. Historical AWS firewall receipts remain
scoped to v15.9.0. Frozen v15.10.0 currently qualifies AWS and Azure certified-hardware reads only.
GCP and Oracle live qualification is deferred.

`capabilities-v1.json` enumerates 7,173 attribute, block and structural object paths across 53 related
operation surfaces, including 1,362 paths in the SMSv2 resource. `selections` contains description-level
groups. `source_audit` additionally records 976 source expressions for selections, conflicting inputs and
replacement behavior across 100 source files. These are declarations to qualify, not proof that validation or
lifecycle works.

Each capability records schema presence, classification, validation receipts, lifecycle receipts, per-cloud
runtime receipts and outstanding work separately. Supported classifications are implemented, defective,
unqualified, contract-blocked, deprecated and deferred. Use a runtime receipt for the specific applicable
cloud, path and revision before marking runtime qualified. `implemented` alone never means traffic acceptance.
The current operation ownership paths are candidates; verify the recorded source files and owning generator or
API before repairing a defect. Related surfaces such as AWS TGW and KVM remain in the inventory even though
the initial fixtures are single-node AWS and Azure.

The frozen baseline is provider v15.9.0, API v12.0.8 and consumer commit
`947d3f60098eec5f4c1536aa2d338523e53f8687`. Exact provider source, audit source, API source, installed Linux
binary and schema digests are in `pins-v1.json`. Audit source can differ from the published artifact source;
do not attribute an audit-source constraint to another installed release without verification.

## Reproduce or refresh

Use an isolated directory and an explicit installation configuration:

```hcl
terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = "= 15.9.0"
    }
  }
}
```

```hcl
provider_installation {
  direct {}
}
```

Save the second block to `direct.tfrc`, then run:

```sh
export TF_CLI_CONFIG_FILE="$PWD/direct.tfrc"
terraform init -backend=false -input=false
terraform providers schema -json > schema.json
sha256sum schema.json
```

Verify the selected version and installed binary checksum. From the provider checkout at the audit source
revision, generate twice and compare:

```sh
python3 scripts/smsv2_capabilities.py \
  --schema /path/to/schema.json --source-root . \
  --pins coverage/smsv2/pins-v1.json --output /path/to/first.json
python3 scripts/smsv2_capabilities.py \
  --schema /path/to/schema.json --source-root . \
  --pins coverage/smsv2/pins-v1.json --output /path/to/second.json
cmp /path/to/first.json /path/to/second.json
```

For a successor, create new exact pins and pass `--previous` with the prior ledger. Matching IDs retain
acceptance evidence; removed IDs stay visible as contract-blocked. A qualified runtime field without a receipt
fails generation. Retained receipts remain evidence only for their recorded revisions; independently qualify a
successor before declaring current acceptance. Keep plans, state, credentials, raw registrations and
identity-bearing runtime evidence private; commit only sanitized receipt summaries and hashes.

## Reconciliation and current gaps

| Capability | Classification | Current evidence and remaining work |
| --- | --- | --- |
| Enhanced Firewall attachment | Partially qualified on AWS | [AWS fixture receipt](https://github.com/f5-sales-demo/multi-cloud-networking/pull/1502) proves scoped TCP 80 allow/deny with HTTPS control, ordered rules, reference replacement, policy import/data-source equality, no-op/drift and deletion/cleanup on v15.9.0. Three reference/detachment paths have scoped AWS runtime receipts; classification and complete parent lifecycle remain unqualified. Both orders of the allow/deny policy pair deny HTTP; cross-policy precedence and Azure remain unqualified. |
| Forward Proxy attachment | Unqualified | Parent attachment and policy surfaces are present. Prove HTTP URL and TLS/SNI filtering, order, `NEXT_POLICY`, removal and combined firewall operation. Classify TLS decryption separately. |
| Registration certified hardware | Published v15.10.0; AWS/Azure reads accepted | [#2593](https://github.com/f5-sales-demo/terraform-provider-xcsh/pull/2593) added the computed-only lookup. The signed installed v15.10.0 artifact returns the exact authoritative registration value on both clouds, with zero managed changes and XC writes. Separate [AWS receipt](certified-hw-aws-v15.10.0-receipt.json) and [Azure receipt](certified-hw-v15.10.0-receipt.json) retain saved-plan hashes. #2129 is closed; newer artifacts and other capabilities require separate receipts. |
| Omitted MTU runtime read | Defective, contract unqualified | [#2600](https://github.com/f5-sales-demo/terraform-provider-xcsh/issues/2600): v15.9.0 runtime lookup rejects omitted MTU read as 0 even while fresh HTTP/HTTPS traverse the CE. Establish supported defaults and authoritative physical reads; no invented MTU or broader import exclusions. |
| In-place version unpinning | Contract-blocked | [#1387](https://github.com/f5-sales-demo/terraform-provider-xcsh/issues/1387) remains open. Audit source adds replacement for changed `software_settings`. Require supported update/clear and authoritative reads; replacement is not an unpinning workaround. |
| Azure Route Server multihop | Contract-blocked | Compiled contract reports `no_schema_valid_ebgp_multihop_request_control`. Retain rejection until an authoritative successor contract is published. |
| AWS multi-node site HA | Contract-blocked | Published node-configuration contract is single-node, two interfaces and HA disabled. Establish a multi-node contract before claiming site HA; independent sites are a different resilience topology. |
| GCP and Oracle runtime | Deferred | Retain static/mock coverage. No AWS or Azure result establishes their live acceptance. |

The historical consumer `coverage/smsv2-coverage-matrix.md` describes provider v3.x, including API
configuration probes without backing Azure VMs. Its counts and successful applies are historical evidence, not
current image boot, readiness, routing or enforcement acceptance. Consumer issues #646, #647 and #648 retain
networking and advanced-mode follow-ups; this ledger does not close them or expand import exclusions.

## Disposable fixture gates

Use separate task-owned AWS and Azure state and resource identities. Keep the substrate stable within a
capability cycle. The fixture needs a real CE with SLO and SLI, a private traffic client, and controlled HTTP
and HTTPS origins. Before policy acceptance, record image identity and boot, registration, ONLINE readiness,
MAC-to-device joins, routing and baseline traffic.

For every security case, generate a fresh traffic identifier and prove both the CE path and origin
observation. Cloud security groups must permit the test flows so they cannot imitate CE policy enforcement.
The existing `xcsh_securemesh_site_v2` attachment blocks remain the sole site-association writer.

Review and hash saved apply, upgrade and destroy plans. Verify identity, unrelated settings and concurrency
preservation. Mutable cases require create/read/update, explicit and omitted defaults, contradictory choices,
list order, import/data-source consistency, repeated no-op plans, drift and deletion. Preserve first failures
separately from corrective runs and classify fixture, provider, platform and prerequisite failures.

Each child delivery needs applicable focused and full regression gates, deterministic generation,
documentation validation, public PII/secret checks, PR review/CI, regeneration, immutable publication and
installed AWS/Azure acceptance. Destroy and verify task-owned cloud and XC absence at cycle completion. Close
live child issues only with receipts; the umbrella remains open while any deferred or contract-blocked
qualifications remain.

AWS receipt references are pinned to the exact consumer commit and content hash. They qualify only
the recorded v15.9.0 scenario. They do not promote all firewall matcher paths, ordered policy
precedence, a newer provider, Azure, TLS decryption or later SMSv2 families. First failures and
corrective runs are preserved in the receipt; all raw state and identity-bearing evidence stays private.

The current Azure parent import probe on v15.9.0 did not converge: create-only/write-only
`software_settings` was absent from imported state and interface labels changed from null to empty
maps. The pinned baseline then proposed replacement. This was a read-only probe with zero XC
mutations and no excluded attributes; it does not establish supported in-place unpinning.

`capabilities-v2.json` freezes provider v15.10.0 and API v12.0.9 with 7,174 paths and 53
operation surfaces. It retains prior receipts while resetting their current-artifact runtime
status to unqualified. Only the signed-artifact AWS and Azure `certified_hw` reads have successor
qualification. Exact pins and the public-safe signed-artifact receipt are adjacent to the ledger.
