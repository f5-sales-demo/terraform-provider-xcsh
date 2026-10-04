---
page_title: "access_logs_s3_params.aws_credentials"
subcategory: ""
description: "Configuration parameter for aws credentials."
xcsh_docs: {"aliases": ["access logs s3 params aws credentials", "authentication", "credential setup", "credentials"], "body_bytes": 1658, "body_sha256": "sha256:5294905dd03a6eec937f3a60163f4b23185603cc35ac5cc05fbeffeae5eceef1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials:secret_access_key"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials", "parent_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "path": "documentation/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3333313333033032-0021303200000201-2311303202212332-1220022311211322-2312032121330332-3203003302132332-1333333332032311-2102222113212110", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_logs_s3_params", "aws_credentials"], "schema_version": 1, "sections": [{"aliases": ["access logs s3 params aws credentials access key id"], "anchor": "schema-access_logs_s3_params--aws_credentials--access_key_id", "description": "AWS Access key ID. AWS Access key ID.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials", "access_key_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["access logs s3 params aws credentials region"], "anchor": "schema-access_logs_s3_params--aws_credentials--region", "description": "AWS Region. AWS Region.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials", "region"], "syntax": "attribute", "type": "string"}, {"aliases": ["access logs s3 params aws credentials secret access key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials:secret_access_key", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials", "secret_access_key"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Configuration parameter for aws credentials.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_logs_s3_params.aws_credentials

Breadcrumbs:

- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/)
- [access_logs_s3_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/)
- access_logs_s3_params.aws_credentials

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for aws credentials.

## Direct properties

<a id="schema-access_logs_s3_params--aws_credentials--access_key_id"></a>

### access_key_id property

Type: `"string"`. Computed.

AWS Access key ID. AWS Access key ID.

<a id="schema-access_logs_s3_params--aws_credentials--region"></a>

### region property

Type: `"string"`. Computed.

AWS Region. AWS Region.

- [secret_access_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/): complete subsection reference.

## Next pages

- [access_logs_s3_params.aws_credentials.secret_access_key](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/secret_access_key/)
- [access_logs_s3_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/properties/access_logs_s3_params/)
- [xcsh_lma_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/lma_region/)
