---
page_title: "items.object.system_metadata.initializers.result"
subcategory: ""
description: "Status is a return value for calls that don't return other objects."
xcsh_docs: {"aliases": ["items object system metadata initializers result"], "body_bytes": 1901, "body_sha256": "sha256:2f05988e155fbcee7db0e97849a76675f105ba0141fe040d85c57dbd0a9bbba0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:initializers:result", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:initializers", "path": "documentation/data-sources/site_registrations_by_site/properties/items/object/system_metadata/initializers/result/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3123213313301213-3212220113203033-0201212311303212-0313113200010131-3213121300121131-2133011001230202-1120210212021310-0033222122200330", "registry_path": "docs/guides/data-sources--site_registrations_by_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "system_metadata", "initializers", "result"], "schema_version": 1, "sections": [{"aliases": ["items object system metadata initializers result code"], "anchor": "schema-items--object--system_metadata--initializers--result--code", "description": "Suggested HTTP return code for this status, 0 if not set.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:initializers:result", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "result", "code"], "syntax": "attribute", "type": "number"}, {"aliases": ["items object system metadata initializers result reason"], "anchor": "schema-items--object--system_metadata--initializers--result--reason", "description": "Human-readable description of why this operation is in the 'Failure' status. If this value is empty there is no information available.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:initializers:result", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "result", "reason"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata initializers result status", "succeeded", "success", "successful"], "anchor": "schema-items--object--system_metadata--initializers--result--status", "description": "Status of the operation. One of: 'Success' or 'Failure'.", "document_id": "xcsh-docs:data-sources:site_registrations_by_site:properties:items:object:system_metadata:initializers:result", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "result", "status"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/properties/items/object/system_metadata/initializers/result/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Status is a return value for calls that don't return other objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.system_metadata.initializers.result

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/)
- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/system_metadata/)
- [items.object.system_metadata.initializers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/items/object/system_metadata/initializers/)
- items.object.system_metadata.initializers.result

<a id="section"></a>

Type: `"single"`. Computed.

Status is a return value for calls that don't return other objects.

## Direct properties

<a id="schema-items--object--system_metadata--initializers--result--code"></a>

### code property

Type: `"number"`. Computed.

Suggested HTTP return code for this status, 0 if not set.

<a id="schema-items--object--system_metadata--initializers--result--reason"></a>

### reason property

Type: `"string"`. Computed.

Human-readable description of why this operation is in the 'Failure' status. If this value is empty
there is no information available.

<a id="schema-items--object--system_metadata--initializers--result--status"></a>

### status property

Type: `"string"`. Computed.

Status of the operation. One of: 'Success' or 'Failure'.
