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
The first recovered session reports an external-neighbor hop limit of 255.
It establishes, but that observation does not prove the documented TTL 2 contract;
packet-level evidence and a supported configuration path remain outstanding.

## Independent request captures

`tools/capture-legacy-smsv2.py` runs the pinned Linux amd64 Volterra 0.12.2
binary against an isolated loopback HTTPS API. It checks the binary digest,
uses a temporary test certificate and state, and never loads cloud credentials.
The twelve discovery fixtures in `tools/testdata/legacy-smsv2/` cover every
legacy platform branch. Repeating the AWS capture is byte-identical.

Eleven discovery/software-default requests also matched the current locally
built provider requests exactly. The twelfth, `rseries`, is absent from the
current Terraform schema and fails before sending a request. A subsequent live request received an explicit unsupported-platform rejection;
its evidence now establishes platform removal. In that fixture, legacy `enable_ha = false` is
omitted rather than serialized as a disable-HA choice. These captures prove
request serialization only; discovery, realized nodes and platform lifecycles
still require their separate acceptance evidence.

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
A reviewed saved Terraform plan then changed only the first CE SLI device from
`eth1` to the MAC-verified `ens6`, in place, using the locally built provider
candidate. Site UID was preserved, API read-back matched, and the SLI MAC appeared
as a physical user-visible forwarding interface. BGP observations remained empty.
This confirms an interface-selection defect without establishing the complete
BGP failure cause. The regenerated physical object has a new name; the existing
SLI connector initially referenced the removed old object. A second reviewed
saved plan updated only that binding and its recorded runtime observation. The
first SLI BGP session then established, independently confirmed by node neighbor
output and XC telemetry. A third reviewed plan added only the second AWS endpoint
on that tunnel. Both sessions established; the first remained up while the second
was added. Both initially reported zero accepted prefixes. AWS lacked Connect route-table
associations and propagation. Separate reviewed plans enabled those settings for
the SLI attachment. Both sessions then accepted the workload prefix, AWS reported
both sessions up, and the local provider confirmed both same-MAC sessions and
workload route agreement. HTTP to the VIP still timed out: the virtual site
selected zero CEs and the expected VIP route was absent. The first CE now carries the explicit MCN topology label. Updating the virtual
site selector exposed a second defect: its ReplaceSpecType is empty and the
server ignored the spec change. Enrichment now marks selector/site type immutable;
the generator preserves mutability through references and emits block replacement
modifiers. The regenerated and reviewed provider recreated the virtual site; independent
read-back confirms its selector and selection of CE01. The other tunnels and traffic remain unverified. Separate AWS boot-console reads correlated all three
sites' ENI MACs with guest devices `ens5` and `ens6`; only the first SLI has been
changed so far.
See F5's [Site CLI reference](https://docs.cloud.f5.com/docs-v2/multi-cloud-network-connect/reference/ea-sitecli-ref).

The reference deployment's live status remains unverified; its checkout is pinned
to `15897bb91e06ef5ed5b6057c1ee73159f59a0af7`.

The first existing MCN site received the single SLI device update described above;
its remote Terraform state was refreshed by apply. The other two sites were
inspected read-only. The isolated JWT test objects were created and deleted. The first SLI Connect
attachment association and propagation were added; CE instances remain intact. Existing v7.4.1 remains
unchanged, and no PR or release has been published. The full generation command remains blocked: the pinned v6.1.1 manifest has
584 unresolved paths. The corrected local manifest and verified rSeries removal
now leave 467 unresolved paths.
An operational publication hold is installed and verified in GitHub: provider
`on-merge.yml`, `release-manual.yml`, `_tag-release.yml` and enrichment
`sync-and-enrich.yml` are disabled, with no active runs at installation.
PR checks remain enabled. The final digest-bound governed promotion gate remains
outstanding; the hold must stay in place until its acceptance conditions are met. Complete parity, live traffic/redundancy/upgrades,
the second rebuild, refreshed no-change plan, final artifact receipts and direct
AGY review remain outstanding.

## Consolidated delivery

Issue #2031 will deliver the provider parity repairs in one consolidated PR.
Keep the failing-regression, implementation, regeneration and verification loop
local until the complete candidate is reviewable. Enrichment #1738 and MCN #1097
remain linked changes in their owning repositories; do not split individual
parity repairs into additional PRs or push each local iteration into CI.

This is a clean-break prerelease. Compatibility shims and rollback support are
not acceptance requirements. The authorized MCN development environment is
ephemeral and may be destroyed and rebuilt to verify idempotence. Exact target
identity, reviewed saved plans and preservation of unrelated resources still
apply.

## Interface-addressing evidence

The installed legacy binary accepts explicit-node `static-dns` and `dhcp-server`
fixtures. Their captures retain the synthetic Terraform inputs alongside the
outgoing request, so omitted enum defaults can be distinguished from unconfigured
fields. In the AWS capture, the deprecated static `dns_server` is transmitted.
The DHCP capture transmits option 82, fixed MAC assignments and excluded address
pools. The explicitly configured include-pools enum is omitted on the wire as its
protobuf default. These observations do not establish that a cloud API accepts or
realizes these settings.

The current [F5 SMSv2 FAQ](https://docs.cloud.f5.com/docs-v2/multi-cloud-network-connect/faqs/smsv2-faqs),
checked on September 6, 2026, lists DHCP server support on baremetal as a feature
gap. Its scope does not establish removal across all legacy platform branches.
The [KVM deployment instructions](https://docs.cloud.f5.com/docs-v2/multi-cloud-network-connect/how-to/site-management/deploy-sms-kvm-clickops)
list DHCP client and static IP as interface-address choices. These documents,
legacy serialization and current schema absence must be evaluated separately;
they do not justify silently classifying the repeated missing paths as parity.

## Verified platform removal and virtual-site recovery

The current XC API rejected the exact pinned legacy rSeries discovery request
with HTTP 400: `Rseries provider is not supported for SecureMeshSite`. The
sanitized enrichment evidence binds that result to the legacy fixture and
protected live receipt digests. All 112 rSeries paths are recorded as verified
platform removal; rSeries is not carried forward. The regenerated matrix leaves
467 unresolved paths, with no other removal inferred from deprecation alone.

Provider candidate `9fd9cc817230` passed local tests, affected race checks, lint,
pre-commit and direct AGY review. Its reviewed saved Terraform plan replaced
only the VIP virtual site. Independent API reads confirmed a new UID, the desired
`mcn-topology` selector and selection of CE01. This establishes selector lifecycle
behavior and placement selection; traffic and complete-topology acceptance remain
separate requirements.

Following the selector replacement, the bounded HTTP request from the owned AWS
workload instance to the configured VIP and Host returned HTTP 200. Both first-SLI
BGP sessions remained established and each accepted the workload prefix. This is
first-tunnel recovery evidence, not acceptance of the remaining topology,
redundancy, upgrades or rebuilds.

## Current API addressing correction

Isolated current-API create/read probes retained the configured static DNS and
DHCP-server fields, including option 82 and excluded pools. Each unregistered
probe object was deleted and its absence verified. Enrichment commit `9c93baac`
corrects the shared interface schema with those observed fields and choice
membership. Enrichment regressions and pre-commit passed. A provider regression
first failed on all eleven supported platform branches and then passed after
regeneration; it also asserts that rSeries remains absent. The regenerated matrix
now has 236 unresolved paths. Realized DHCP behavior remains unverified.

A separate reviewed saved plan changed only CE01's SLO device from `eth0` to
`ens5`, in place, using provider candidate `9fd9cc817230`. The site UID was
preserved. API read-back and physical user-visible forwarding interfaces now
agree on both MAC-correlated devices, `ens5` and `ens6`. Two BGP observations
remain; SLO connector and session recovery continues serially.

Topology labels were applied in separate reviewed plans to CE02 and CE03 and
verified on both SMSv2 and registered-site objects. A fresh plan then changed only
CE01's SLO connector reference to the actual `ens5` interface object plus its
runtime observation state. Node neighbor output subsequently showed three
established BGP sessions: the first SLO endpoint and both SLI endpoints. The SLI
sessions still received the workload prefix; SLO route advertisement and its
second endpoint remain separate verification steps.

A reviewed BGP-only update added the second AWS-assigned SLO endpoint without
changing existing peer inputs. Node neighbor output now confirms four established
sessions on CE01, two per tunnel. Both SLI sessions accept the workload prefix;
both SLO sessions currently accept zero prefixes. SLO route-table attachment
configuration and recovery of CE02/CE03 remain outstanding.

The addressing candidate passed full internal and tooling suites, affected
provider race tests and deterministic regeneration (679 provider/client Go files
were byte-identical). Enrichment commit `9c93baac` passed direct AGY review.

## Twelve-session recovery evidence

Fresh independent observations on September 6 confirmed all twelve sessions
Established in XC and up in AWS across the three CEs and six Connect peers. Each
session's imported route path contains the expected workload prefix and its exact
AWS endpoint. A bounded request from the owned workload instance to the shared
VIP returned HTTP 200. The evidence receipt has SHA256
`a3915c42a4733c67d6b480bac22db36ecb09855669bf8ab7c496765e15c7c290`;
its protected content-addressed archive preserves the observations and local
provider artifact `55914ce5b92bd391840d4802b073869387f7d7b6`.

CE02 and CE03 recovered through separate reviewed device, connector and BGP
updates. Their site UIDs were preserved, and physical user-visible interfaces
matched the verified devices by MAC. This supersedes the partial topology status
above. Shared VIP availability does not identify which CE carried that request.
Per-session redundancy, CE failover, VIP route identity, the multihop contract,
serial upgrades, the second rebuild and a refresh-enabled no-change plan remain
outstanding. The 236 unresolved parity paths still block publication.

The route checker previously pooled imported and exported prefixes across a
node. A route learned through another session could therefore satisfy an
expected peer's route requirement. The new regression requires an imported path
whose remote address matches that session; missing path identity and exported
routes alone cannot satisfy it. This behavioral correction is part of the same
consolidated provider PR. The twelve-session receipt predates this correction
and does not validate its final artifact.

## Physical transport readiness

The runtime reader now requires registered-site interface status in addition to
site provisioning and resolved interface objects. Each requested transport must
match its configured node, MAC, device and physical SLO/SLI network role in one
published interface-status document that is not marked stale. Its Ethernet link
must report up. Missing, ambiguous, stale and down observations fail bounded
readiness polling with a physical-link diagnostic. Backup interfaces remain
healthy when their physical link is up.

Read-only observations of the recovered three-CE lab show all six physical links
up with the expected MAC and device identities. These observations establish the
API response shape. The API's status publication metadata identifies the node; the reader
does not construct a VER instance name or infer physical health from tunnel status.

A targeted live refresh with clean provider commit `2134873dd3d3` subsequently
passed all three runtime readers and all three BGP readers: six healthy physical
interfaces and twelve established sessions, each with its own imported workload
route path. The binary SHA256 is
`4a9df8a3dedb2101124dcf59d5383bafa1e630c16662126777f335ae4439703d`.
The saved plan contains no managed-resource changes. This targeted refresh does
not establish a full refresh-enabled no-change plan or release acceptance.

## Logging contract investigation

Two reproducible captures from the pinned Volterra 0.12.2 binary exercise the
legacy top-level `log_receiver` and the network-aware `log_receiver_with_net`
configuration. An isolated current-API experiment accepted both requests and
returned identical logging configuration: `log_receiver_with_net.log_receiver`
with `use_slo_sli`. All three probe objects were deleted and their absence verified.
The sanitized enrichment evidence binds the captures and live receipt by digest.

The new provider retains the network-aware interface without a legacy alias.
Live log delivery remains unverified, so the four legacy logging paths remain
unresolved. Configuration normalization alone does not close the capability gap.

## Geographic selection and upgrade drain settings

The pinned legacy binary sends `re_select.specific_geography` and
`drain_max_unavailable_node_percentage`. Isolated current-API create/read probes
retained both configured values. The drain response also added the default
`disable_vega_upgrade_mode` choice. Source enrichment regressions first failed
for the missing properties, then passed after restoring the properties and their
mutually exclusive choice membership.

Regeneration from enrichment commit `f54e7874` leaves 234 unresolved legacy
paths. The parity gate continues to fail until those capabilities are classified
with adequate evidence; rSeries remains absent.

A separate legacy private-ADN request was accepted, but the field was absent
from read-back. That result does not establish an equivalent replacement or
verified removal. Private ADN remains unresolved and is not added to the provider.
All probe objects were removed and their absence verified. Actual Regional Edge
selection and percentage-based draining during an upgrade remain unverified.
