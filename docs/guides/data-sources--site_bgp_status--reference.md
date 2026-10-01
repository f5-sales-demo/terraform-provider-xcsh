---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_site_bgp_status."
xcsh_docs: {"aliases": [], "body_bytes": 5727, "body_sha256": "sha256:f59775e123ffa88337551721c6f75cf5c1885b4dd8c812bb736f3161a9a0317f", "canonical_id": "xcsh-docs:data-sources:site_bgp_status:reference", "child_ids": ["xcsh-docs:data-sources:site_bgp_status:properties:expected_peers", "xcsh-docs:data-sources:site_bgp_status:properties:peers"], "collection_id": "xcsh-docs:data-sources:site_bgp_status:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_bgp_status:reference", "parent_id": "xcsh-docs:data-sources:site_bgp_status:fundamentals", "path": "docs/guides/data-sources--site_bgp_status--reference.md", "provider_name": "site_bgp_status", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_bgp_status/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_site_bgp_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md)
- Property reference

## Direct properties

<a id="schema-bgp_routes_json"></a>

### bgp_routes_json property

Type: `"string"`. Computed.

<a id="schema-converged"></a>

### converged property

Type: `"bool"`. Computed.

<a id="schema-expected_exported_routes"></a>

### expected_exported_routes property

Type: `["set", "string"]`. Required.

Exact prefixes that every expected node must export and carry in the selected SLO or SLI route view.

- [expected_peers](data-sources--site_bgp_status--properties--expected_peers.md): complete subsection reference.

<a id="schema-id"></a>

### id property

Type: `"string"`. Computed.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.OneOf("system")}
```

- [peers](data-sources--site_bgp_status--properties--peers.md): complete subsection reference.

<a id="schema-poll_interval_seconds"></a>

### poll_interval_seconds property

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 60)}
```

<a id="schema-site"></a>

### site property

Type: `"string"`. Required.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{stringvalidator.LengthAtLeast(1)}
```

<a id="schema-sli_routes_json"></a>

### sli_routes_json property

Type: `"string"`. Computed.

<a id="schema-slo_routes_json"></a>

### slo_routes_json property

Type: `"string"`. Computed.

<a id="schema-timeout_seconds"></a>

### timeout_seconds property

Type: `"number"`. Optional, Computed.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{int64validator.Between(1, 1800)}
```

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `bgp_routes_json` | [bgp_routes_json](data-sources--site_bgp_status--reference.md#schema-bgp_routes_json) |
| `converged` | [converged](data-sources--site_bgp_status--reference.md#schema-converged) |
| `expected_exported_routes` | [expected_exported_routes](data-sources--site_bgp_status--reference.md#schema-expected_exported_routes) |
| `expected_peers` | [expected_peers](data-sources--site_bgp_status--properties--expected_peers.md#section) |
| `expected_peers.expected_imported_routes` | [expected_peers.expected_imported_routes](data-sources--site_bgp_status--properties--expected_peers.md#schema-expected_peers--expected_imported_routes) |
| `expected_peers.mac` | [expected_peers.mac](data-sources--site_bgp_status--properties--expected_peers.md#schema-expected_peers--mac) |
| `expected_peers.node` | [expected_peers.node](data-sources--site_bgp_status--properties--expected_peers.md#schema-expected_peers--node) |
| `expected_peers.peer_address` | [expected_peers.peer_address](data-sources--site_bgp_status--properties--expected_peers.md#schema-expected_peers--peer_address) |
| `expected_peers.role` | [expected_peers.role](data-sources--site_bgp_status--properties--expected_peers.md#schema-expected_peers--role) |
| `id` | [id](data-sources--site_bgp_status--reference.md#schema-id) |
| `namespace` | [namespace](data-sources--site_bgp_status--reference.md#schema-namespace) |
| `peers` | [peers](data-sources--site_bgp_status--properties--peers.md#section) |
| `peers.advertised_prefix_count` | [peers.advertised_prefix_count](data-sources--site_bgp_status--properties--peers.md#schema-peers--advertised_prefix_count) |
| `peers.established` | [peers.established](data-sources--site_bgp_status--properties--peers.md#schema-peers--established) |
| `peers.interface_name` | [peers.interface_name](data-sources--site_bgp_status--properties--peers.md#schema-peers--interface_name) |
| `peers.mac` | [peers.mac](data-sources--site_bgp_status--properties--peers.md#schema-peers--mac) |
| `peers.node` | [peers.node](data-sources--site_bgp_status--properties--peers.md#schema-peers--node) |
| `peers.peer_address` | [peers.peer_address](data-sources--site_bgp_status--properties--peers.md#schema-peers--peer_address) |
| `peers.received_prefix_count` | [peers.received_prefix_count](data-sources--site_bgp_status--properties--peers.md#schema-peers--received_prefix_count) |
| `peers.role` | [peers.role](data-sources--site_bgp_status--properties--peers.md#schema-peers--role) |
| `peers.state` | [peers.state](data-sources--site_bgp_status--properties--peers.md#schema-peers--state) |
| `peers.state_changed_at` | [peers.state_changed_at](data-sources--site_bgp_status--properties--peers.md#schema-peers--state_changed_at) |
| `poll_interval_seconds` | [poll_interval_seconds](data-sources--site_bgp_status--reference.md#schema-poll_interval_seconds) |
| `site` | [site](data-sources--site_bgp_status--reference.md#schema-site) |
| `sli_routes_json` | [sli_routes_json](data-sources--site_bgp_status--reference.md#schema-sli_routes_json) |
| `slo_routes_json` | [slo_routes_json](data-sources--site_bgp_status--reference.md#schema-slo_routes_json) |
| `timeout_seconds` | [timeout_seconds](data-sources--site_bgp_status--reference.md#schema-timeout_seconds) |

## Next pages

- [expected_peers](data-sources--site_bgp_status--properties--expected_peers.md)
- [peers](data-sources--site_bgp_status--properties--peers.md)
- [xcsh_site_bgp_status](../data-sources/site_bgp_status.md)
