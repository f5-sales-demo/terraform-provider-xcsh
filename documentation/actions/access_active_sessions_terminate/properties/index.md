---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_access_active_sessions_terminate."
xcsh_docs: {"aliases": [], "body_bytes": 1432, "body_sha256": "sha256:94b3cae7df3431b6ef80d8649b05be77b16d4054dae008fedbf676acb6e369eb", "child_ids": [], "collection_id": "xcsh-docs:actions:access_active_sessions_terminate:collection", "completeness": "complete", "id": "xcsh-docs:actions:access_active_sessions_terminate:reference", "parent_id": "xcsh-docs:actions:access_active_sessions_terminate:fundamentals", "path": "documentation/actions/access_active_sessions_terminate/properties/index.md", "provider_name": "access_active_sessions_terminate", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_sessions_terminate/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_access_active_sessions_terminate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Property reference

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/)
- Property reference

## Direct properties

<a id="schema-ids"></a>

### ids property

Type: `["list", "string"]`. Optional.

List of session IDs to terminate (maximum 100 per request).

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{
  listvalidator.SizeBetween(1, 100),
}
```

<a id="schema-namespace"></a>

### namespace property

Type: `"string"`. Required.

Namespace Namespace of the App type for the current request (path parameter)

## All schema paths

Each exact path has one authoritative reference destination. Collection element indexes are runtime positions; schema paths name the subsection.

| Schema path | Complete reference |
| --- | --- |
| `ids` | [ids](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/properties/#schema-ids) |
| `namespace` | [namespace](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/properties/#schema-namespace) |

## Next pages

- [xcsh_access_active_sessions_terminate](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/)
