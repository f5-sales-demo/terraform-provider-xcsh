---
page_title: "image_list.aws.image_id"
subcategory: ""
description: "Configuration for image_id."
xcsh_docs: {"aliases": ["image list aws image id"], "body_bytes": 1337, "body_sha256": "sha256:fcf5545b8493eab27b5c7ca23bbe5814a2083d088ed7211b514f76a4c8a3caac", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:aws:image_id", "parent_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:aws", "path": "documentation/data-sources/certified_hardware/properties/image_list/aws/image_id/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-1302333203022310-0211233233301012-3001210121223011-1203202333201212-1131222301032102-2010332103110012-2022112310232311-2123330102001302", "registry_path": "docs/guides/data-sources--certified_hardware--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["image_list", "aws", "image_id"], "schema_version": 1, "sections": [{"aliases": ["image id"], "anchor": "schema-image_list--aws--image_id--image_id", "description": "AWS ami image name. AWS ami image.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:aws:image_id", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_list", "aws", "image_id", "image_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["region"], "anchor": "schema-image_list--aws--image_id--region", "description": "AWS ami image region. AWS ami image region.", "document_id": "xcsh-docs:data-sources:certified_hardware:properties:image_list:aws:image_id", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["image_list", "aws", "image_id", "region"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/properties/image_list/aws/image_id/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration for image_id.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [image_list.aws](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/properties/image_list/aws/)
- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
