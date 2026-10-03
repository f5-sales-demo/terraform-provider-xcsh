---
page_title: "access_logs_s3_params"
subcategory: ""
description: "Configuration parameter for access logs s3 params."
xcsh_docs: {"aliases": ["access logs s3 params"], "body_bytes": 1230, "body_sha256": "sha256:b9352bcea6e934c0cf0c87544d085ee60c71e1bb039bf79f36521ed0b42833bb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "parent_id": "xcsh-docs:data-sources:lma_region:reference", "path": "documentation/data-sources/lma_region/properties/access_logs_s3_params/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3003330033321322-1111330132110212-0220302231333132-1331112112201111-1330231001000000-1202301331222333-2313111100322211-2320323330310311", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_logs_s3_params"], "schema_version": 1, "sections": [{"aliases": ["access logs s3 params aws credentials", "authentication", "credential setup", "credentials"], "anchor": "section", "description": "Configuration parameter for aws credentials.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials"], "syntax": "attribute", "type": "object"}, {"aliases": ["access logs s3 params bucket"], "anchor": "schema-access_logs_s3_params--bucket", "description": "S3 Bucket Name. S3 Bucket Name.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_logs_s3_params", "bucket"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/access_logs_s3_params/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Configuration parameter for access logs s3 params.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": [], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
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

## Next pages

- [access_logs_s3_params.aws_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
