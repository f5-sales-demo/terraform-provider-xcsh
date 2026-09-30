---
page_title: "items"
subcategory: ""
description: "items for xcsh_access_active_sessions."
xcsh_docs: {"aliases": [], "body_bytes": 2357, "body_sha256": "sha256:fa674624ac9d0ed5923d567ecea341c18e9a226ddc7383433e6e1dfdc5ba17ca", "child_ids": [], "collection_id": "xcsh-docs:data-sources:access_active_sessions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_sessions:properties:items", "parent_id": "xcsh-docs:data-sources:access_active_sessions:reference", "path": "documentation/data-sources/access_active_sessions/properties/items/index.md", "provider_name": "access_active_sessions", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["items"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_sessions/properties/items/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items for xcsh_access_active_sessions.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# items

Breadcrumbs:

- [xcsh_access_active_sessions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/properties/)
- items

<a id="section"></a>

Type: `"list"`. Computed.

Sessions. List of active sessions.

## Direct properties

<a id="schema-items--client_ip"></a>

### client_ip property

Type: `"string"`. Computed.

Client IP of the user that connected via the session.

<a id="schema-items--expiration_time"></a>

### expiration_time property

Type: `"string"`. Computed.

Expiration time of the session in RFC 3339 format.

<a id="schema-items--id"></a>

### id property

Type: `"string"`. Computed.

Session ID. ID of the session.

<a id="schema-items--last_activity_time"></a>

### last_activity_time property

Type: `"string"`. Computed.

Last activity time of the session in RFC 3339 format.

<a id="schema-items--policy"></a>

### policy property

Type: `"string"`. Computed.

Policy. Policy associated with the session.

<a id="schema-items--site"></a>

### site property

Type: `"string"`. Computed.

Site. Site where the session is created.

<a id="schema-items--start_time"></a>

### start_time property

Type: `"string"`. Computed.

Start time of the session in RFC 3339 format.

<a id="schema-items--status"></a>

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

<a id="schema-items--username"></a>

### username property

Type: `"string"`. Computed.

Username used in authentication of this session.

<a id="schema-items--virtual_server"></a>

### virtual_server property

Type: `"string"`. Computed.

Virtual server associated with the session.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/properties/)
- [xcsh_access_active_sessions](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/)
