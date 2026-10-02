---
page_title: "items.system_metadata.initializers.result"
subcategory: ""
description: "Status is a return value for calls that don't return other objects."
xcsh_docs: {"aliases": ["items system metadata initializers result"], "body_bytes": 1957, "body_sha256": "sha256:86212f9defd9b950e38d2daa386b6baf3f312de682dbcacda014478151b46759", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:initializers:result", "parent_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:initializers", "path": "documentation/data-sources/site_registrations/properties/items/system_metadata/initializers/result/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0112212002013103-1113110302202100-0130100201332201-0112031220023102-2033211231310001-0313333212332121-0003131320033131-1212112011202321", "registry_path": "docs/guides/data-sources--site_registrations--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["items", "system_metadata", "initializers", "result"], "schema_version": 1, "sections": [{"aliases": ["code"], "anchor": "schema-items--system_metadata--initializers--result--code", "description": "Suggested HTTP return code for this status, 0 if not set.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:initializers:result", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "initializers", "result", "code"], "syntax": "attribute", "type": "number"}, {"aliases": ["reason"], "anchor": "schema-items--system_metadata--initializers--result--reason", "description": "Human-readable description of why this operation is in the 'Failure' status. If this value is empty there is no information available.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:initializers:result", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "initializers", "result", "reason"], "syntax": "attribute", "type": "string"}, {"aliases": ["status", "succeeded", "success", "successful"], "anchor": "schema-items--system_metadata--initializers--result--status", "description": "Status of the operation. One of: 'Success' or 'Failure'.", "document_id": "xcsh-docs:data-sources:site_registrations:properties:items:system_metadata:initializers:result", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["items", "system_metadata", "initializers", "result", "status"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/properties/items/system_metadata/initializers/result/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Status is a return value for calls that don't return other objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# items.system_metadata.initializers.result

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/)
- [items](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/)
- [items.system_metadata](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/system_metadata/)
- [items.system_metadata.initializers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/system_metadata/initializers/)
- items.system_metadata.initializers.result

<a id="section"></a>

Type: `"single"`. Computed.

Status is a return value for calls that don't return other objects.

## Direct properties

<a id="schema-items--system_metadata--initializers--result--code"></a>

### code property

Type: `"number"`. Computed.

Suggested HTTP return code for this status, 0 if not set.

<a id="schema-items--system_metadata--initializers--result--reason"></a>

### reason property

Type: `"string"`. Computed.

Human-readable description of why this operation is in the 'Failure' status. If this value is empty
there is no information available.

<a id="schema-items--system_metadata--initializers--result--status"></a>

### status property

Type: `"string"`. Computed.

Status of the operation. One of: 'Success' or 'Failure'.

## Next pages

- [items.system_metadata.initializers](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/properties/items/system_metadata/initializers/)
- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
