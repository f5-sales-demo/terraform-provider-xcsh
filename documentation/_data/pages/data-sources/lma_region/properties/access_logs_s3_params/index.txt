---
page_title: "access_logs_s3_params"
subcategory: ""
description: "Configuration parameter for access logs s3 params."
xcsh_docs: {"aliases": ["access logs s3 params"], "body_bytes": 823, "body_sha256": "sha256:3ab372c42854b6eb7fb7b8bbf4c0b06fb4df509a066acddec8af1162514be906", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "parent_id": "xcsh-docs:data-sources:lma_region:reference", "path": "documentation/data-sources/lma_region/properties/access_logs_s3_params/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3003330033321322-1111330132110212-0220302231333132-1331112112201111-1330231001000000-1202301331222333-2313111100322211-2320323330310311", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_logs_s3_params"], "schema_version": 1, "sections": [{"aliases": ["access logs s3 params aws credentials", "authentication", "credential setup", "credentials"], "anchor": "section", "description": "Configuration parameter for aws credentials.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials"], "syntax": "attribute", "type": "object"}, {"aliases": ["access logs s3 params bucket"], "anchor": "schema-access_logs_s3_params--bucket", "description": "S3 Bucket Name. S3 Bucket Name.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_logs_s3_params", "bucket"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/access_logs_s3_params/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Configuration parameter for access logs s3 params.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_logs_s3_params

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- access_logs_s3_params

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for access logs s3 params.

## Direct properties

- [aws_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/): complete subsection reference.

<a id="schema-access_logs_s3_params--bucket"></a>

### bucket property

Type: `"string"`. Computed.

S3 Bucket Name. S3 Bucket Name.
