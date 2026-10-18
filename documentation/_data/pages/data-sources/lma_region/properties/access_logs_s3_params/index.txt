---
page_title: "access_logs_s3_params"
subcategory: ""
description: "Configuration parameter for access logs s3 params."
xcsh_docs: {"aliases": ["access logs s3 params"], "body_bytes": 823, "body_sha256": "sha256:3ab372c42854b6eb7fb7b8bbf4c0b06fb4df509a066acddec8af1162514be906", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "parent_id": "xcsh-docs:data-sources:lma_region:reference", "path": "documentation/data-sources/lma_region/properties/access_logs_s3_params/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3003330033321322-1111330132110212-0220302231333132-1331112112201111-1330231001000000-1202301331222333-2313111100322211-2320323330310311", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_logs_s3_params"], "schema_version": 1, "sections": [{"aliases": ["access logs s3 params aws credentials", "authentication", "credential setup", "credentials"], "anchor": "section", "description": "Configuration parameter for aws credentials.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials"], "syntax": "attribute", "type": "object"}, {"aliases": ["access logs s3 params bucket"], "anchor": "schema-access_logs_s3_params--bucket", "description": "S3 Bucket Name. S3 Bucket Name.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_logs_s3_params", "bucket"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/access_logs_s3_params/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Configuration parameter for access logs s3 params.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": [], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
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
