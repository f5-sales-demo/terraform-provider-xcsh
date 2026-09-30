---
page_title: "Property reference"
subcategory: ""
description: "Property reference for xcsh_access_active_sessions_terminate."
xcsh_docs: {"aliases": [], "body_bytes": 1125, "body_sha256": "sha256:94f6cd5c2c9a5b987f40e24bdd429add6c41675bb21fbe34ce33dbe93c1e22f4", "canonical_id": "xcsh-docs:actions:access_active_sessions_terminate:reference", "child_ids": [], "collection_id": "xcsh-docs:actions:access_active_sessions_terminate:collection", "completeness": "complete", "id": "xcsh-docs:actions:access_active_sessions_terminate:reference", "parent_id": "xcsh-docs:actions:access_active_sessions_terminate:fundamentals", "path": "docs/guides/actions--access_active_sessions_terminate--reference.md", "provider_name": "access_active_sessions_terminate", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "publishing_destination": "registry", "role": "reference", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_sessions_terminate/properties/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Property reference for xcsh_access_active_sessions_terminate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Property reference

Breadcrumbs:

- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md)
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
| `ids` | [ids](actions--access_active_sessions_terminate--reference.md#schema-ids) |
| `namespace` | [namespace](actions--access_active_sessions_terminate--reference.md#schema-namespace) |

## Next pages

- [xcsh_access_active_sessions_terminate](../actions/access_active_sessions_terminate.md)
