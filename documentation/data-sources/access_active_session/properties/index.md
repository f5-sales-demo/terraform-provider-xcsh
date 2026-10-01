---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_access_active_session."
xcsh_docs: {"aliases": [], "body_bytes": 4900, "body_sha256": "sha256:567908b9e490c2e5fecfb2ff9314f749b5a18ebe0431d382192e456fa103daaa", "child_ids": ["xcsh-docs:data-sources:access_active_session:properties:variables"], "collection_id": "xcsh-docs:data-sources:access_active_session:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_session:reference", "parent_id": "xcsh-docs:data-sources:access_active_session:fundamentals", "path": "documentation/data-sources/access_active_session/properties/index.md", "provider_name": "access_active_session", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_session/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_access_active_session.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_access_active_session](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/)
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

- [variables](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/variables/): complete subsection reference.

<a id="schema-virtual_server"></a>

### virtual_server property

Type: `"string"`. Computed.

Virtual server associated with the session.

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `client_ip` | [client_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-client_ip) |
| `expiration_time` | [expiration_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-expiration_time) |
| `id` | [id](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-id) |
| `last_activity_time` | [last_activity_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-last_activity_time) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-namespace) |
| `policy` | [policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-policy) |
| `site` | [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-site) |
| `start_time` | [start_time](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-start_time) |
| `status` | [status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-status) |
| `username` | [username](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-username) |
| `variables` | [variables](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/variables/#section) |
| `variables.value` | [variables.value](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/variables/#schema-variables--value) |
| `variables.variable` | [variables.variable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/variables/#schema-variables--variable) |
| `virtual_server` | [virtual_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/#schema-virtual_server) |

## Next pages

- [variables](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/properties/variables/)
- [xcsh_access_active_session](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_session/)
