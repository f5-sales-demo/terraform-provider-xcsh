---
page_title: "aws_cloud_watch_receiver"
subcategory: ""
description: "AWS Cloudwatch Logs Configuration for Global Log Receiver."
xcsh_docs: {"aliases": ["aws cloud watch receiver"], "body_bytes": 10225, "body_sha256": "sha256:e941279e708b243c076a1fd08d13ca7b0501e73ad03f61f3787db4b8efd62a7b", "capabilities": ["monitoring"], "category": "monitoring", "child_ids": ["xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:aws_cred", "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch", "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:compression"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "parent_id": "xcsh-docs:resources:global_log_receiver:reference", "path": "documentation/resources/global_log_receiver/properties/aws_cloud_watch_receiver/index.md", "product": "distributed-cloud", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3221021130232030-0231000300303232-3323332003022321-1200333231022020-1131300231330322-3302031000232121-0000022002032222-1001121332110333", "registry_path": "docs/guides/resources--global_log_receiver--reference--group-001.md", "relationships": [{"anchor": "schema-aws_cloud_watch_receiver--aws_region", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver:RequiredObjectAttributes:aws_region,group_name,stream_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "type": "requires"}, {"anchor": "schema-aws_cloud_watch_receiver--group_name", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver:RequiredObjectAttributes:aws_region,group_name,stream_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "type": "requires"}, {"anchor": "schema-aws_cloud_watch_receiver--stream_name", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver:RequiredObjectAttributes:aws_region,group_name,stream_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["aws_cloud_watch_receiver"], "schema_version": 1, "sections": [{"aliases": ["aws cloud watch receiver aws cred"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:aws_cred", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_cloud_watch_receiver--aws_cred--name", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.aws_cred:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:aws_cred", "type": "requires"}], "schema_path": ["aws_cloud_watch_receiver", "aws_cred"], "syntax": "block", "type": "object"}, {"aliases": ["aws cloud watch receiver aws region"], "anchor": "schema-aws_cloud_watch_receiver--aws_region", "description": "AWS Region Name.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ca-central-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-west-1", "us-west-2"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_cloud_watch_receiver", "aws_region"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws cloud watch receiver batch"], "anchor": "section", "description": "Batch OPTIONS allow tuning for how batches of logs are sent to an endpoint.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-aws_cloud_watch_receiver--batch--max_bytes", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch", "type": "conflicts"}, {"anchor": "schema-aws_cloud_watch_receiver--batch--max_events", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch", "type": "conflicts"}, {"anchor": "schema-aws_cloud_watch_receiver--batch--timeout_seconds", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.batch:ConflictingObjectAttributes:max_bytes,max_bytes_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch:max_bytes_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.batch:ConflictingObjectAttributes:max_events,max_events_disabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch:max_events_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.batch:ConflictingObjectAttributes:timeout_seconds,timeout_seconds_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:batch:timeout_seconds_default", "type": "conflicts"}], "schema_path": ["aws_cloud_watch_receiver", "batch"], "syntax": "block", "type": "object"}, {"aliases": ["aws cloud watch receiver compression"], "anchor": "section", "description": "Compression Type.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:compression", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.compression:ConflictingObjectAttributes:compression_default,compression_gzip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_gzip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.compression:ConflictingObjectAttributes:compression_default,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_none", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "aws_cloud_watch_receiver.compression:ConflictingObjectAttributes:compression_gzip,compression_none", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver:compression:compression_none", "type": "conflicts"}], "schema_path": ["aws_cloud_watch_receiver", "compression"], "syntax": "block", "type": "object"}, {"aliases": ["aws cloud watch receiver group name"], "anchor": "schema-aws_cloud_watch_receiver--group_name", "description": "The group name of the target Cloudwatch Logs stream.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_cloud_watch_receiver", "group_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["aws cloud watch receiver stream name"], "anchor": "schema-aws_cloud_watch_receiver--stream_name", "description": "The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a log stream at a time.", "document_id": "xcsh-docs:resources:global_log_receiver:properties:aws_cloud_watch_receiver", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["aws_cloud_watch_receiver", "stream_name"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/global_log_receiver/properties/aws_cloud_watch_receiver/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "AWS Cloudwatch Logs Configuration for Global Log Receiver.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_cloud_watch_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/)
- aws_cloud_watch_receiver

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: aws\_cloud\_watch\_receiver, azure\_event\_hubs\_receiver, azure\_receiver,
datadog\_receiver, gcp\_bucket\_receiver, http\_receiver, kafka\_receiver, new\_relic\_receiver,
qradar\_receiver, s3\_receiver, splunk\_receiver, sumo\_logic\_receiver\] AWS Cloudwatch Logs
Configuration for Global Log Receiver.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("aws_region",
    "group_name",
    "stream_name")}
```

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

OneOf alternatives in this subsection:

- [aws_cloud_watch_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/#section)
- [azure_event_hubs_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_event_hubs_receiver/#section)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/azure_receiver/#section)
- [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/datadog_receiver/#section)
- [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/gcp_bucket_receiver/#section)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/http_receiver/#section)
- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/kafka_receiver/#section)
- [new_relic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/new_relic_receiver/#section)
- [qradar_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/qradar_receiver/#section)
- [s3_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/s3_receiver/#section)
- [splunk_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/splunk_receiver/#section)
- [sumo_logic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/sumo_logic_receiver/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
aws_cloud_watch_receiver {
  # Configure direct properties listed below.
}
```

## Direct properties

- [aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/aws_cred/): complete subsection reference.

<a id="schema-aws_cloud_watch_receiver--aws_region"></a>

### aws_region property

Type: `"string"`. Optional.

\[Enum:
ap-northeast-1|ap-southeast-1|eu-central-1|eu-west-1|eu-west-3|sa-east-1|us-east-1|us-east-2|us-west-2|ca-central-1|af-south-1|ap-east-1|ap-south-1|ap-northeast-2|ap-southeast-2|eu-south-1|eu-north-1|eu-west-2|me-south-1|us-west-1|ap-southeast-3\]
AWS Region. AWS Region Name. Possible values are \`ap-northeast-1\`, \`ap-southeast-1\`,
\`eu-central-1\`, \`eu-west-1\`, \`eu-west-3\`, \`sa-east-1\`, \`us-east-1\`, \`us-east-2\`,
\`us-west-2\`, \`ca-central-1\`, \`af-south-1\`, \`ap-east-1\`, \`ap-south-1\`, \`ap-northeast-2\`,
\`ap-southeast-2\`, \`eu-south-1\`, \`eu-north-1\`, \`eu-west-2\`, \`me-south-1\`, \`us-west-1\`,
\`ap-southeast-3\`.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["af-south-1","ap-east-1","ap-northeast-1","ap-northeast-2","ap-south-1","ap-southeast-1","ap-southeast-2","ap-southeast-3","ca-central-1","eu-central-1","eu-north-1","eu-south-1","eu-west-1","eu-west-2","eu-west-3","me-south-1","sa-east-1","us-east-1","us-east-2","us-west-1","us-west-2"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("ap-northeast-1",
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
    "ap-southeast-3"),
}
```

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/): complete subsection reference.

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/): complete subsection reference.

<a id="schema-aws_cloud_watch_receiver--group_name"></a>

### group_name property

Type: `"string"`. Optional.

The group name of the target Cloudwatch Logs stream.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 3,
    "pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[\\\\.\\\\-_/#A-Za-z0-9]+$"
  }
}
```

<a id="schema-aws_cloud_watch_receiver--stream_name"></a>

### stream_name property

Type: `"string"`. Optional.

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
Validators: []validator.String{
  stringvalidator.LengthBetween(3, 512),
}
```

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 512,
  "minLength": 3,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 512,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-08T03:45:36+00:00"
    },
    "minLength": 3,
    "pattern": "^[^:*]*$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[^:*]*$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.max_len": "512",
    "ves.io.schema.rules.string.min_len": "3",
    "ves.io.schema.rules.string.pattern": "^[^:*]*$"
  }
}
```
