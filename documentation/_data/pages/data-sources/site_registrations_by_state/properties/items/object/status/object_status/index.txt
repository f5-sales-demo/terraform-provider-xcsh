---
page_title: "items.object.status.object_status"
subcategory: ""
description: "items.object.status.object_status for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 1920, "body_sha256": "sha256:2fc10d6e4515a2956b76b1844bc9333e51fb8e809b2322f0e427a4745b1e4d31", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:status:object_status", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:status", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/status/object_status/index.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["items", "object", "status", "object_status"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/status/object_status/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "items.object.status.object_status for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.status.object_status

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- [items.object.status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/status/)
- items.object.status.object_status

<a id="section"></a>

Type: `"single"`. Computed.

Status is a return value for calls that don't return other objects.

## Direct properties

<a id="schema-items--object--status--object_status--code"></a>

### code property

Type: `"number"`. Computed.

Suggested HTTP return code for this status, 0 if not set.

<a id="schema-items--object--status--object_status--reason"></a>

### reason property

Type: `"string"`. Computed.

Human-readable description of why this operation is in the 'Failure' status. If this value is empty
there is no information available.

<a id="schema-items--object--status--object_status--status"></a>

### status property

Type: `"string"`. Computed.

Status of the operation. One of: 'Success' or 'Failure'.

## Next pages

- [items.object.status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/status/)
- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
