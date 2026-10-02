---
page_title: "s3_receiver"
subcategory: ""
description: "S3 Configuration for Global Log Receiver."
xcsh_docs: {"aliases": ["s3 receiver"], "body_bytes": 6175, "body_sha256": "sha256:7dad719b42f7c985116d2ab778ff4c3daf093c0a218f7f93a5ef5429f9014ba9", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:aws_cred", "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:compression", "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "documentation/data-sources/global_log_receiver/properties/s3_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3122321231311212-1111122011230100-1332330302001111-0100123100212313-2310312223130013-0122212111212000-1312212133311330-2113220302103202", "registry_path": "docs/guides/data-sources--global_log_receiver--reference--group-004.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["s3_receiver"], "schema_version": 1, "sections": [{"aliases": ["aws cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:aws_cred", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["s3_receiver", "aws_cred"], "syntax": "attribute", "type": "object"}, {"aliases": ["aws region"], "anchor": "schema-s3_receiver--aws_region", "description": "AWS Region Name.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["s3_receiver", "aws_region"], "syntax": "attribute", "type": "string"}, {"aliases": ["batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:batch", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["s3_receiver", "batch"], "syntax": "attribute", "type": "object"}, {"aliases": ["bucket"], "anchor": "schema-s3_receiver--bucket", "description": "S3 Bucket Name.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["s3_receiver", "bucket"], "syntax": "attribute", "type": "string"}, {"aliases": ["compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:compression", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["s3_receiver", "compression"], "syntax": "attribute", "type": "object"}, {"aliases": ["filename options"], "anchor": "section", "description": "Filename OPTIONS allow customization of filename and folder paths used by a destination endpoint bucket or file.", "document_id": "xcsh-docs:data-sources:global_log_receiver:properties:s3_receiver:filename_options", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["s3_receiver", "filename_options"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/s3_receiver/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "S3 Configuration for Global Log Receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# s3_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- s3_receiver

<a id="section"></a>

Type: `"single"`. Computed.

S3 Configuration for Global Log Receiver.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Direct properties

- [aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/aws_cred/): complete subsection reference.

<a id="schema-s3_receiver--aws_region"></a>

### aws_region property

Type: `"string"`. Computed.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
AWS Region. AWS Region Name. Possible values are \`ap-northeast-1\`, \`ap-southeast-1\`,
\`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`, \`us-east-2\`,
\`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`, \`ap-northeast-2\`,
\`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`, \`me-south-1\`, \`us-west-1\`,
\`ap-southeast-3\`.

Upstream description:

AWS Region Name.

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "ap-northeast-1",
    "ap-southeast-1",
    "eu-central-1",
    "eu-west-1",
    "eu-west-3",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-2",
    "ca-central-1",
    "af-south-1",
    "ap-east-1",
    "ap-south-1",
    "ap-northeast-2",
    "ap-southeast-2",
    "eu-south-1",
    "eu-north-1",
    "eu-west-2",
    "me-south-1",
    "us-west-1",
    "ap-southeast-3"
  ],
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.in": "[\\\"ap-northeast-1\\\",\\\"ap-southeast-1\\\",\\\"eu-central-1\\\",\\\"eu-west-1\\\",\\\"eu-west-3\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-2\\\",\\\"ca-central-1\\\",\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-south-1\\\",\\\"ap-northeast-2\\\",\\\"ap-southeast-2\\\",\\\"eu-south-1\\\",\\\"eu-north-1\\\",\\\"eu-west-2\\\",\\\"me-south-1\\\",\\\"us-west-1\\\",\\\"ap-southeast-3\\\"]"
  }
}
```

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/batch/): complete subsection reference.

<a id="schema-s3_receiver--bucket"></a>

### bucket property

Type: `"string"`. Computed.

S3 Bucket Name. S3 Bucket Name.

Upstream description:

S3 Bucket Name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 128,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 128,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 3,
    "pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "128",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[a-z0-9]+[a-z0-9\\\\.-]+[a-z0-9]$"
  }
}
```

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/compression/): complete subsection reference.

- [filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/filename_options/): complete subsection reference.

## Next pages

- [s3_receiver.aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/aws_cred/)
- [s3_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/batch/)
- [s3_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/compression/)
- [s3_receiver.filename_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/filename_options/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
