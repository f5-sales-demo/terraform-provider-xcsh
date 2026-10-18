---
page_title: "image_list.aws.image_id"
subcategory: ""
description: "Configuration for image_id."
xcsh_docs: {"aliases": ["image list aws image id"], "body_bytes": 1068, "body_sha256": "sha256:8b35b68234401b30493bd765d90911476f6a8817a1cb81af94cd8fa3f6e36f85", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:aws:image_id", "parent_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:aws", "path": "documentation/data-sources/certified_hardware/properties/image_list/aws/image_id/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1302333203022310-0211233233301012-3001210121223011-1203202333201212-1131222301032102-2010332103110012-2022112310232311-2123330102001302", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["image_list", "aws", "image_id"], "schema_version": 1, "sections": [{"aliases": ["image list aws image id image id"], "anchor": "schema-image_list--aws--image_id--image_id", "description": "AWS ami image name. AWS ami image.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:aws:image_id", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_list", "aws", "image_id", "image_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["image list aws image id region"], "anchor": "schema-image_list--aws--image_id--region", "description": "AWS ami image region. AWS ami image region.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:aws:image_id", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_list", "aws", "image_id", "region"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/image_list/aws/image_id/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Configuration for image_id.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": [], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# image_list.aws.image_id

Breadcrumbs:

- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/)
- [image_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/)
- [image_list.aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/aws/)
- image_list.aws.image_id

<a id="section"></a>

Type: `"single"`. Computed.

Configuration for image\_id.

## Direct properties

<a id="schema-image_list--aws--image_id--image_id"></a>

### image_id property

Type: `"string"`. Computed.

AWS ami image name. AWS ami image.

<a id="schema-image_list--aws--image_id--region"></a>

### region property

Type: `"string"`. Computed.

AWS ami image region. AWS ami image region.
