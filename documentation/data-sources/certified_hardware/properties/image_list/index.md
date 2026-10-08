---
page_title: "image_list"
subcategory: ""
description: "List of image names with providers for this certified hardware, e.g. AWS ami-0f99d090261d2acd5."
xcsh_docs: {"aliases": ["image list"], "body_bytes": 1329, "body_sha256": "sha256:be61831a4e1b43e70533416c762703ad8211a05578335cff69549f53418c206b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:certified_hardware:properties:image_list:aws", "xcsh-docs:data-sources:certified_hardware:properties:image_list:azure", "xcsh-docs:data-sources:certified_hardware:properties:image_list:gcp"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:image_list", "parent_id": "xcsh-docs:data-sources:certified_hardware:reference", "path": "documentation/data-sources/certified_hardware/properties/image_list/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3301201231202101-0111203210233310-3010321103121133-1110131201111123-3310102031021120-0121212002303313-1220201232030313-2301130001301220", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["image_list"], "schema_version": 1, "sections": [{"aliases": ["image list aws"], "anchor": "section", "description": "AWS. AWS specific information.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:aws", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["image_list", "aws"], "syntax": "attribute", "type": "object"}, {"aliases": ["image list azure"], "anchor": "section", "description": "Azure. Azure specific information.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:azure", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["image_list", "azure"], "syntax": "attribute", "type": "object"}, {"aliases": ["image list gcp"], "anchor": "section", "description": "GCP. GCP specific information.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:gcp", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["image_list", "gcp"], "syntax": "attribute", "type": "object"}, {"aliases": ["image list name"], "anchor": "schema-image_list--name", "description": "Name. Image name to use for this hardware.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_list", "name"], "syntax": "attribute", "type": "string"}, {"aliases": ["image list provider ref"], "anchor": "schema-image_list--provider_ref", "description": "Image provider F5 Distributed Cloud, Cloud provider like AWS or Azure.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_list", "provider_ref"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/image_list/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "List of image names with providers for this certified hardware, e.g. AWS ami-0f99d090261d2acd5.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# image_list

Breadcrumbs:

- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/)
- image_list

<a id="section"></a>

Type: `"list"`. Computed.

List of image names with providers for this certified hardware, e.g. AWS ami-0f99d090261d2acd5.

## Direct properties

- [aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/aws/): complete subsection reference.

- [azure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/azure/): complete subsection reference.

- [gcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/gcp/): complete subsection reference.

<a id="schema-image_list--name"></a>

### name property

Type: `"string"`. Computed.

Name. Image name to use for this hardware.

<a id="schema-image_list--provider_ref"></a>

### provider_ref property

Type: `"string"`. Computed.

Image provider F5 Distributed Cloud, Cloud provider like AWS or Azure.
