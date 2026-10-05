---
page_title: "access_logs_s3_params.aws_credentials"
subcategory: ""
description: "Configuration parameter for aws credentials."
xcsh_docs: {"aliases": ["access logs s3 params aws credentials", "authentication", "credential setup", "credentials"], "body_bytes": 1658, "body_sha256": "sha256:5294905dd03a6eec937f3a60163f4b23185603cc35ac5cc05fbeffeae5eceef1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials:secret_access_key"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:lma_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials", "parent_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params", "path": "documentation/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/index.md", "product": "distributed-cloud", "provider_name": "lma_region", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3333313333033032-0021303200000201-2311303202212332-1220022311211322-2312032121330332-3203003302132332-1333333332032311-2102222113212110", "registry_path": "docs/guides/data-sources--lma_region--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["access_logs_s3_params", "aws_credentials"], "schema_version": 1, "sections": [{"aliases": ["access logs s3 params aws credentials access key id"], "anchor": "schema-access_logs_s3_params--aws_credentials--access_key_id", "description": "AWS Access key ID. AWS Access key ID.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials", "access_key_id"], "syntax": "attribute", "type": "string"}, {"aliases": ["access logs s3 params aws credentials region"], "anchor": "schema-access_logs_s3_params--aws_credentials--region", "description": "AWS Region. AWS Region.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials", "region"], "syntax": "attribute", "type": "string"}, {"aliases": ["access logs s3 params aws credentials secret access key"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:lma_region:properties:access_logs_s3_params:aws_credentials:secret_access_key", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["access_logs_s3_params", "aws_credentials", "secret_access_key"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/lma_region/properties/access_logs_s3_params/aws_credentials/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Configuration parameter for aws credentials.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
