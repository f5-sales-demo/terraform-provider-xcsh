---
page_title: "items.object.status.object_status"
subcategory: ""
description: "Status is a return value for calls that don't return other objects."
xcsh_docs: {"aliases": ["items object status object status"], "body_bytes": 1911, "body_sha256": "sha256:14d1992315aafdabc2b4b5a3e03b599bc68fb62d42d1165a32b583de3039ed90", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status:object_status", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/status/object_status/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0211302121122020-2100302003213322-2033323302131313-2210322133332310-0211112303301201-1321320013121000-2012203200023013-2331203103212023", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "status", "object_status"], "schema_version": 1, "sections": [{"aliases": ["code"], "anchor": "schema-items--object--status--object_status--code", "description": "Suggested HTTP return code for this status, 0 if not set.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status:object_status", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "status", "object_status", "code"], "syntax": "attribute", "type": "number"}, {"aliases": ["reason"], "anchor": "schema-items--object--status--object_status--reason", "description": "Human-readable description of why this operation is in the 'Failure' status. If this value is empty there is no information available.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status:object_status", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "status", "object_status", "reason"], "syntax": "attribute", "type": "string"}, {"aliases": ["login success", "status", "succeeded", "success", "successful"], "anchor": "schema-items--object--status--object_status--status", "description": "Status of the operation. One of: 'Success' or 'Failure'.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:status:object_status", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "status", "object_status", "status"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/status/object_status/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Status is a return value for calls that don't return other objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.status.object_status

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/status/)
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

- [items.object.status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/status/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
