# SMSv2 independent parity candidate

This candidate is incomplete and must not be published. Track the remaining work in
[f5-sales-demo/terraform-provider-xcsh#2031](https://github.com/f5-sales-demo/terraform-provider-xcsh/issues/2031),
[enrichment #1738](https://github.com/f5-sales-demo/api-specs-enriched/issues/1738), and
[MCN #1097](https://github.com/f5-sales-demo/mcn/issues/1097).

## Independent baseline

Volterra 0.12.2 is pinned to commit
`22f029dbf14412c99502fe1daba829f7c3261017`. The extractor verifies the SMSv2 source
SHA-256 before parsing it and compares every extracted path against the installed
provider's exported schema. All 1,699 source paths match; the installed schema also
contains the SDK-injected resource ID. The manifest records both input digests.
This establishes a structural baseline, not protobuf mapping or lifecycle parity.

Build `tools/extract-legacy-smsv2.go` as an executable, then pass the pinned source,
`terraform providers schema -json` output for Volterra 0.12.2, and the output path.
Building first avoids Go interpreting the upstream `.go` argument as another source
file. A modified source or mismatched installed schema is rejected.

The refreshed matrix has 579 unclassified legacy paths. Deprecation, the existence
of a differently named cloud-site resource, and a successful write followed by
server stripping no longer justify automatic exclusions or platform-removal claims.
Generator gaps also fail the comparison. The comparison writes its diagnostic
matrix before returning failure so missing capabilities remain reviewable.

Existing `current_parity` and `modernized_semantics` entries describe structural
comparisons only. Independent request/protobuf mappings, responses, behavioral
redesign evidence, dependent-resource coverage and lifecycle receipts are still
required before functional acceptance. Do not treat these labels as release approval.

## Discovery configuration

The provider accepts omitted and empty node lists in `aws.not_managed` for
registration-time discovery. Previously, handwritten validation rejected this
configuration before any API request. Explicit nodes retain their validation.
An isolated live test of the local binary created an unregistered site and its
site-bound JWT, obtained no-change refresh plans before and after token import,
and deleted both objects. Subsequent GET requests returned 404 and the local
state was empty. This does not prove registered-node runtime discovery.

## Interface object discovery

Runtime and BGP readers resolve physical interface objects from the platform list
API, requesting the realized specification and ownership metadata. They correlate
the site's immutable UID, namespace, node and ethernet device with the configured
MAC, role and MTU. Names come from the matching object rather than a naming formula.
Missing, duplicated, foreign-owned, inconsistent and partially returned objects
fail discovery. This preserves a distinction between the physical transport object
and a tunnel interface reported by BGP.

The validator accepts discovered guest device names and rejects duplicate devices
within a node. The previous requirement that SLO and SLI use `eth0` and `eth1`
was unsupported by the API's EthernetInterfaceType contract.

A saved Terraform plan exposed a second limitation: the generated AWS rule forced
whole-site replacement for a device-only edit. An isolated, unregistered non-HA
site accepted a device-only PUT, returned the changed device, and was deleted.
Changing its node count was separately rejected by the API. The generator now
permits only device-only edits on explicitly non-HA, single-node sites to update
in place; other AWS topology edits retain replacement behavior. The rejected
replacement plan was not applied to the lab. Terraform also marks computed-only
`is_primary` and `is_management` observations unknown during edits. Topology
comparisons exclude those outputs without replacing their planned unknown values;
requested input changes still determine replacement behavior. This follows
HashiCorp's [plan modification process](https://developer.hashicorp.com/terraform/plugin/framework/resources/plan-modification).

The interface-list status arrays were empty in live observations. The existing
runtime `healthy` output still reflects global site provisioning and must not be
used as proof of per-interface health. Replacing that health contract and supporting
registered discovered nodes remain release blockers.

## BGP session behavior

`xcsh_site_bgp_status` accepts distinct remote IP addresses on one node/MAC. It
normalizes IP spelling and rejects duplicate node/address expectations before HTTP
requests, including duplicates that supply different MACs. MAC remains a transport
correlation input; `role` selects the payload route view.

The current BGP telemetry response exposes node and remote address but no routing
context or connector identifier. Ambiguous observations fail closed. Complete
context-aware session identity requires an authoritative observation contract.

The MCN candidate configures both AWS endpoints on each of six Connect peers and
expects twelve sessions. Its status output replaces `peer_count` with
`connect_peer_count` and `bgp_session_count`. AWS documents dual-session redundancy
and requires eBGP multihop TTL 2 in its
[Connect contract](https://docs.aws.amazon.com/vpc/latest/tgw/tgw-connect.html).
XC's realized multihop behavior has not been established.

## Live findings and release boundary

Scoped read-only checks on the existing three-site AWS lab returned all six
physical network-interface objects and all six BGP-referenced tunnel-interface
objects. Names matched configuration references. Each site still returned zero BGP
observations. This does not establish interface health or explain the missing peers.
The documented read-only Site CLI commands work through `exec-user` using the
configured short node name. On the first CE, BGP summary and neighbor commands
returned successfully with no substantive output. Its forwarding table correlated
the SLO MAC with physical `ens5`; the configured SLI MAC was absent. AWS still
reported both matching ENIs attached. The same instance's console boot output
independently correlated those MACs with `ens5` and `ens6`.
This is a testable device-discovery hypothesis,
not an established cause. No device name was guessed or applied to the lab.
See F5's [Site CLI reference](https://docs.cloud.f5.com/docs-v2/multi-cloud-network-connect/reference/ea-sitecli-ref).

The reference deployment's live status remains unverified; its checkout is pinned
to `15897bb91e06ef5ed5b6057c1ee73159f59a0af7`.

The existing MCN sites were inspected read-only. The isolated JWT test objects
were created and deleted; no AWS deployment was changed. Existing v7.4.1 remains
unchanged, and no PR or release has been published. The full generation command remains blocked: the pinned v6.1.1 manifest has
584 unresolved paths, while the corrected local enrichment manifest has 579.
The governed publication hold is not yet installed; substantive merges remain pending because the current release
workflows publish automatically. Complete parity, live traffic/redundancy/upgrades,
the second rebuild, refreshed no-change plan, final artifact receipts and direct
AGY review remain outstanding.
