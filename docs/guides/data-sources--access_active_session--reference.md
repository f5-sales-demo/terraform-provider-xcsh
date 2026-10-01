---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_access_active_session."
xcsh_docs: {"aliases": [], "body_bytes": 3960, "body_sha256": "sha256:40a06ab855f9ce0a095a86a8ca91c9ac5d6c874217cd0d415eb7fb52b6f33137", "canonical_id": "xcsh-docs:data-sources:access_active_session:reference", "child_ids": ["xcsh-docs:data-sources:access_active_session:properties:variables"], "collection_id": "xcsh-docs:data-sources:access_active_session:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_session:reference", "parent_id": "xcsh-docs:data-sources:access_active_session:fundamentals", "path": "docs/guides/data-sources--access_active_session--reference.md", "provider_name": "access_active_session", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_session/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_access_active_session.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_access_active_session](../data-sources/access_active_session.md)
- Property reference

## Direct properties

<a id="schema-client_ip"></a>

### client_ip property

Type: `"string"`. Computed.

Client IP of the user that connected via the session.

<a id="schema-expiration_time"></a>

### expiration_time property

Type: `"string"`. Computed.

Expiration time of the session in RFC 3339 format.

<a id="schema-id"></a>

### id property

Type: `"string"`. Required.

ID ID of the session.

<a id="schema-last_activity_time"></a>

### last_activity_time property

Type: `"string"`. Computed.

Last activity time of the session in RFC 3339 format.

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace Namespace of the App type for the current request.

<a id="schema-policy"></a>

### policy property

Type: `"string"`. Computed.

Policy. Policy associated with the session.

<a id="schema-site"></a>

### site property

Type: `"string"`. Computed.

Site. Site where the session is created.

<a id="schema-start_time"></a>

### start_time property

Type: `"string"`. Computed.

Start time of the session in RFC 3339 format.

<a id="schema-status"></a>

### status property

Type: `"string"`. Computed.

\[Enum: PENDING|ESTABLISHED\] The actual status of an active session Session status is pending
Session status is established. Possible values are \`PENDING\`, \`ESTABLISHED\`. Defaults to
\`PENDING\`.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("PENDING",
    "ESTABLISHED"),
}
```

<a id="schema-username"></a>

### username property

Type: `"string"`. Computed.

Username used in authentication of this session.

- [variables](data-sources--access_active_session--properties--variables.md): complete subsection reference.

<a id="schema-virtual_server"></a>

### virtual_server property

Type: `"string"`. Computed.

Virtual server associated with the session.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `client_ip` | [client_ip](data-sources--access_active_session--reference.md#schema-client_ip) |
| `expiration_time` | [expiration_time](data-sources--access_active_session--reference.md#schema-expiration_time) |
| `id` | [id](data-sources--access_active_session--reference.md#schema-id) |
| `last_activity_time` | [last_activity_time](data-sources--access_active_session--reference.md#schema-last_activity_time) |
| `namespace` | [namespace](data-sources--access_active_session--reference.md#schema-namespace) |
| `policy` | [policy](data-sources--access_active_session--reference.md#schema-policy) |
| `site` | [site](data-sources--access_active_session--reference.md#schema-site) |
| `start_time` | [start_time](data-sources--access_active_session--reference.md#schema-start_time) |
| `status` | [status](data-sources--access_active_session--reference.md#schema-status) |
| `username` | [username](data-sources--access_active_session--reference.md#schema-username) |
| `variables` | [variables](data-sources--access_active_session--properties--variables.md#section) |
| `variables.value` | [variables.value](data-sources--access_active_session--properties--variables.md#schema-variables--value) |
| `variables.variable` | [variables.variable](data-sources--access_active_session--properties--variables.md#schema-variables--variable) |
| `virtual_server` | [virtual_server](data-sources--access_active_session--reference.md#schema-virtual_server) |

## Next pages

- [variables](data-sources--access_active_session--properties--variables.md)
- [xcsh_access_active_session](../data-sources/access_active_session.md)
