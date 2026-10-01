---
page_title: "aws_cloud_watch_receiver"
subcategory: ""
description: "aws_cloud_watch_receiver for xcsh_global_log_receiver."
xcsh_docs: {"aliases": [], "body_bytes": 9597, "body_sha256": "sha256:8d90b901854a3331e88bca85be54861b50cb1b362df86a50c503b1d73c124f98", "child_ids": ["xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:aws_cred", "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:batch", "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver:compression"], "collection_id": "xcsh-docs:data-sources:global_log_receiver:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:global_log_receiver:properties:aws_cloud_watch_receiver", "parent_id": "xcsh-docs:data-sources:global_log_receiver:reference", "path": "documentation/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/index.md", "provider_name": "global_log_receiver", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["aws_cloud_watch_receiver"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "aws_cloud_watch_receiver for xcsh_global_log_receiver.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["global_log_receiverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# aws_cloud_watch_receiver

Breadcrumbs:

- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- aws_cloud_watch_receiver

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: aws\_cloud\_watch\_receiver, azure\_event\_hubs\_receiver, azure\_receiver,
datadog\_receiver, gcp\_bucket\_receiver, http\_receiver, kafka\_receiver, new\_relic\_receiver,
qradar\_receiver, s3\_receiver, splunk\_receiver, sumo\_logic\_receiver\] AWS Cloudwatch Logs
Configuration for Global Log Receiver.

Upstream description:

AWS Cloudwatch Logs Configuration for Global Log Receiver.

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

- [aws_cloud_watch_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/#section)
- [azure_event_hubs_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_event_hubs_receiver/#section)
- [azure_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/azure_receiver/#section)
- [datadog_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/datadog_receiver/#section)
- [gcp_bucket_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/gcp_bucket_receiver/#section)
- [http_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/http_receiver/#section)
- [kafka_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/kafka_receiver/#section)
- [new_relic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/new_relic_receiver/#section)
- [qradar_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/qradar_receiver/#section)
- [s3_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/s3_receiver/#section)
- [splunk_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/splunk_receiver/#section)
- [sumo_logic_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/sumo_logic_receiver/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/aws_cred/): complete subsection reference.

<a id="schema-aws_cloud_watch_receiver--aws_region"></a>

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

- [batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/): complete subsection reference.

- [compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/): complete subsection reference.

<a id="schema-aws_cloud_watch_receiver--group_name"></a>

### group_name property

Type: `"string"`. Computed.

The group name of the target Cloudwatch Logs stream.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Type: `"string"`. Computed.

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

Upstream description:

The stream name of the target Cloudwatch Logs stream. Note that there can only be one writer to a
log stream at a time.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [aws_cloud_watch_receiver.aws_cred](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/aws_cred/)
- [aws_cloud_watch_receiver.batch](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/batch/)
- [aws_cloud_watch_receiver.compression](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/aws_cloud_watch_receiver/compression/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/properties/)
- [xcsh_global_log_receiver](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/global_log_receiver/)
