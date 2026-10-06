---
page_title: "items.object.system_metadata.initializers.result"
subcategory: ""
description: "Status is a return value for calls that don't return other objects."
xcsh_docs: {"aliases": ["items object system metadata initializers result"], "body_bytes": 1908, "body_sha256": "sha256:09a6e929f8d324b7fdfaf752af0b808b22af4851a582decd7cad0ab50c2ebf21", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers:result", "parent_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers", "path": "documentation/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/result/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2233223021000222-1201110001112012-3010212330103120-0020303232133213-0223130011113110-1231311030013001-0200020112022003-3312120221021300", "registry_path": "docs/guides/data-sources--site_registrations_by_state--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "object", "system_metadata", "initializers", "result"], "schema_version": 1, "sections": [{"aliases": ["items object system metadata initializers result code"], "anchor": "schema-items--object--system_metadata--initializers--result--code", "description": "Suggested HTTP return code for this status, 0 if not set.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers:result", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "result", "code"], "syntax": "attribute", "type": "number"}, {"aliases": ["items object system metadata initializers result reason"], "anchor": "schema-items--object--system_metadata--initializers--result--reason", "description": "Human-readable description of why this operation is in the 'Failure' status. If this value is empty there is no information available.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers:result", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "result", "reason"], "syntax": "attribute", "type": "string"}, {"aliases": ["items object system metadata initializers result status", "succeeded", "success", "successful"], "anchor": "schema-items--object--system_metadata--initializers--result--status", "description": "Status of the operation. One of: 'Success' or 'Failure'.", "document_id": "xcsh-docs:data-sources:site_registrations_by_state:properties:items:object:system_metadata:initializers:result", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "object", "system_metadata", "initializers", "result", "status"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/result/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Status is a return value for calls that don't return other objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.object.system_metadata.initializers.result

Breadcrumbs:

- [xcsh_site_registrations_by_state](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/)
- [items.object](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/)
- [items.object.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/)
- [items.object.system_metadata.initializers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_state/properties/items/object/system_metadata/initializers/)
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
