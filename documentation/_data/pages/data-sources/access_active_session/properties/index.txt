---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_access_active_session."
xcsh_docs: {"aliases": ["access active session"], "body_bytes": 4900, "body_sha256": "sha256:567908b9e490c2e5fecfb2ff9314f749b5a18ebe0431d382192e456fa103daaa", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:access_active_session:properties:variables"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:access_active_session:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_session:reference", "parent_id": "xcsh-docs:data-sources:access_active_session:fundamentals", "path": "documentation/data-sources/access_active_session/properties/index.md", "product": "distributed-cloud", "provider_name": "access_active_session", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3321011011311001-0000003030110313-1323202031311111-3202313331200301-3100012032100300-3213321201002033-1302101113321300-3113021210022031", "registry_path": "docs/guides/data-sources--access_active_session--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "reference", "schema_path": [], "schema_version": 1, "sections": [{"aliases": ["client ip"], "anchor": "schema-client_ip", "description": "Client IP of the user that connected via the session.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["client_ip"], "syntax": "attribute", "type": "string"}, {"aliases": ["expiration time"], "anchor": "schema-expiration_time", "description": "Expiration time of the session in RFC 3339 format.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["expiration_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["id"], "anchor": "schema-id", "description": "ID ID of the session.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["id"], "syntax": "attribute", "type": "string"}, {"aliases": ["last activity time"], "anchor": "schema-last_activity_time", "description": "Last activity time of the session in RFC 3339 format.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["last_activity_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["namespace"], "anchor": "schema-namespace", "description": "Namespace Namespace of the App type for the current request.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["required"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["namespace"], "syntax": "attribute", "type": "string"}, {"aliases": ["policy"], "anchor": "schema-policy", "description": "Policy. Policy associated with the session.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["policy"], "syntax": "attribute", "type": "string"}, {"aliases": ["site"], "anchor": "schema-site", "description": "Site. Site where the session is created.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["site"], "syntax": "attribute", "type": "string"}, {"aliases": ["start time"], "anchor": "schema-start_time", "description": "Start time of the session in RFC 3339 format.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["start_time"], "syntax": "attribute", "type": "string"}, {"aliases": ["status"], "anchor": "schema-status", "description": "The actual status of an active session Session status is pending Session status is established. Possible values are `PENDING`, `ESTABLISHED`. Defaults to `PENDING`.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["status"], "syntax": "attribute", "type": "string"}, {"aliases": ["username"], "anchor": "schema-username", "description": "Username used in authentication of this session.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["username"], "syntax": "attribute", "type": "string"}, {"aliases": ["variables"], "anchor": "section", "description": "Variables. Session variables as key-value pairs.", "document_id": "xcsh-docs:data-sources:access_active_session:properties:variables", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["variables"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual server"], "anchor": "schema-virtual_server", "description": "Virtual server associated with the session.", "document_id": "xcsh-docs:data-sources:access_active_session:reference", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["virtual_server"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_session/properties/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Property reference for xcsh_access_active_session.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
