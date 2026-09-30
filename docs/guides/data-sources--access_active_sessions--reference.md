---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_access_active_sessions."
xcsh_docs: {"aliases": [], "body_bytes": 8169, "body_sha256": "sha256:92617f3f7a4ed2c9facabf6a123498ce598844f107bbf7e8c684b35beab1339c", "canonical_id": "xcsh-docs:data-sources:access_active_sessions:reference", "child_ids": ["xcsh-docs:data-sources:access_active_sessions:properties:items"], "collection_id": "xcsh-docs:data-sources:access_active_sessions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_sessions:reference", "parent_id": "xcsh-docs:data-sources:access_active_sessions:fundamentals", "path": "docs/guides/data-sources--access_active_sessions--reference.md", "provider_name": "access_active_sessions", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_sessions/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_access_active_sessions.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md)
- Property reference

## Direct properties

<a id="schema-client_ip"></a>

### client_ip property

Type: `"string"`. Optional.

Filter sessions by client IP address. Supports a comma-separated list to match any of the specified
values (e.g. '192.0.2.106,192.0.2.233').

<a id="schema-cursor"></a>

### cursor property

Type: `"string"`. Optional.

Cursor for pagination. Contains encoded sequence\_id for next/previous page.

<a id="schema-direction"></a>

### direction property

Type: `"string"`. Optional.

\[Enum: FORWARD|BACKWARD\] Direction for pagination (forward or backward) GET next page of results
GET previous page of results. Possible values are \`FORWARD\`, \`BACKWARD\`. Defaults to
\`FORWARD\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("FORWARD",
    "BACKWARD"),
}
```

<a id="schema-expiration_time_from"></a>

### expiration_time_from property

Type: `"string"`. Optional.

Filter sessions that expire from this timestamp (inclusive) Format: RFC 3339 (e.g.,
2024-08-06T20:00:00Z)

<a id="schema-expiration_time_to"></a>

### expiration_time_to property

Type: `"string"`. Optional.

Filter sessions that expire up to this timestamp (inclusive) Format: RFC 3339 (e.g.,
2024-08-07T20:00:00Z)

<a id="schema-item_count"></a>

### item_count property

Type: `"number"`. Computed.

Total count of all active sessions (across all pages, after applying filters).

- [items](data-sources--access_active_sessions--properties--items.md): complete subsection reference.

<a id="schema-limit"></a>

### limit property

Type: `"number"`. Optional.

Limits the number of results to the specified number.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace Namespace of the App type for the current request.

<a id="schema-next_cursor"></a>

### next_cursor property

Type: `"string"`. Computed.

Next Cursor. Cursor for the next page of results.

<a id="schema-policy"></a>

### policy property

Type: `"string"`. Optional.

Filter sessions by policy. Supports a comma-separated list to match any of the specified values
(e.g. 'default-policy,strict-policy').

<a id="schema-previous_cursor"></a>

### previous_cursor property

Type: `"string"`. Computed.

Cursor for the previous page of results.

<a id="schema-site"></a>

### site property

Type: `"string"`. Optional.

Filter sessions by site. Supports a comma-separated list to match any of the specified values (e.g.
'us-west-2,eu-west-1').

<a id="schema-start_time_from"></a>

### start_time_from property

Type: `"string"`. Optional.

Filter sessions that started from this timestamp (inclusive) Format: RFC 3339 (e.g.,
2024-08-06T20:00:00Z)

<a id="schema-start_time_to"></a>

### start_time_to property

Type: `"string"`. Optional.

Filter sessions that started up to this timestamp (inclusive) Format: RFC 3339 (e.g.,
2024-08-07T20:00:00Z)

<a id="schema-status"></a>

### status property

Type: `"string"`. Optional.

\[Enum: ANY|PENDING\_ONLY|ESTABLISHED\_ONLY\] Filter sessions by status. If not specified, returns
sessions with all statuses. Return sessions with any status (default when no status filter is
specified) Filter for sessions with pending status Filter for sessions with established status.
Possible values are \`ANY\`, \`PENDING\_ONLY\`, \`ESTABLISHED\_ONLY\`. Defaults to \`ANY\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("ANY",
    "PENDING_ONLY",
    "ESTABLISHED_ONLY"),
}
```

<a id="schema-total_established"></a>

### total_established property

Type: `"number"`. Computed.

Total count of sessions with ESTABLISHED status (after applying filters).

<a id="schema-total_pending"></a>

### total_pending property

Type: `"number"`. Computed.

Total count of sessions with PENDING status (after applying filters).

<a id="schema-username"></a>

### username property

Type: `"string"`. Optional.

Filter sessions by username. Supports a comma-separated list to match any of the specified values
(e.g. 'joe,bob').

<a id="schema-virtual_server"></a>

### virtual_server property

Type: `"string"`. Optional.

Filter sessions by virtual server. Supports a comma-separated list to match any of the specified
values (e.g. 'web-app,API-server').

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `client_ip` | [client_ip](data-sources--access_active_sessions--reference.md#schema-client_ip) |
| `cursor` | [cursor](data-sources--access_active_sessions--reference.md#schema-cursor) |
| `direction` | [direction](data-sources--access_active_sessions--reference.md#schema-direction) |
| `expiration_time_from` | [expiration_time_from](data-sources--access_active_sessions--reference.md#schema-expiration_time_from) |
| `expiration_time_to` | [expiration_time_to](data-sources--access_active_sessions--reference.md#schema-expiration_time_to) |
| `item_count` | [item_count](data-sources--access_active_sessions--reference.md#schema-item_count) |
| `items` | [items](data-sources--access_active_sessions--properties--items.md#section) |
| `items.client_ip` | [items.client_ip](data-sources--access_active_sessions--properties--items.md#schema-items--client_ip) |
| `items.expiration_time` | [items.expiration_time](data-sources--access_active_sessions--properties--items.md#schema-items--expiration_time) |
| `items.id` | [items.id](data-sources--access_active_sessions--properties--items.md#schema-items--id) |
| `items.last_activity_time` | [items.last_activity_time](data-sources--access_active_sessions--properties--items.md#schema-items--last_activity_time) |
| `items.policy` | [items.policy](data-sources--access_active_sessions--properties--items.md#schema-items--policy) |
| `items.site` | [items.site](data-sources--access_active_sessions--properties--items.md#schema-items--site) |
| `items.start_time` | [items.start_time](data-sources--access_active_sessions--properties--items.md#schema-items--start_time) |
| `items.status` | [items.status](data-sources--access_active_sessions--properties--items.md#schema-items--status) |
| `items.username` | [items.username](data-sources--access_active_sessions--properties--items.md#schema-items--username) |
| `items.virtual_server` | [items.virtual_server](data-sources--access_active_sessions--properties--items.md#schema-items--virtual_server) |
| `limit` | [limit](data-sources--access_active_sessions--reference.md#schema-limit) |
| `namespace` | [namespace](data-sources--access_active_sessions--reference.md#schema-namespace) |
| `next_cursor` | [next_cursor](data-sources--access_active_sessions--reference.md#schema-next_cursor) |
| `policy` | [policy](data-sources--access_active_sessions--reference.md#schema-policy) |
| `previous_cursor` | [previous_cursor](data-sources--access_active_sessions--reference.md#schema-previous_cursor) |
| `site` | [site](data-sources--access_active_sessions--reference.md#schema-site) |
| `start_time_from` | [start_time_from](data-sources--access_active_sessions--reference.md#schema-start_time_from) |
| `start_time_to` | [start_time_to](data-sources--access_active_sessions--reference.md#schema-start_time_to) |
| `status` | [status](data-sources--access_active_sessions--reference.md#schema-status) |
| `total_established` | [total_established](data-sources--access_active_sessions--reference.md#schema-total_established) |
| `total_pending` | [total_pending](data-sources--access_active_sessions--reference.md#schema-total_pending) |
| `username` | [username](data-sources--access_active_sessions--reference.md#schema-username) |
| `virtual_server` | [virtual_server](data-sources--access_active_sessions--reference.md#schema-virtual_server) |

## Next pages

- [items](data-sources--access_active_sessions--properties--items.md)
- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md)
