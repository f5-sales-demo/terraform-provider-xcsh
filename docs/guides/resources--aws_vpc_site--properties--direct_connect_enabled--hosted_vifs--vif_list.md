---
page_title: "direct_connect_enabled.hosted_vifs.vif_list"
subcategory: "Infrastructure"
description: "direct_connect_enabled.hosted_vifs.vif_list for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 6713, "body_sha256": "sha256:f4dfca0a942198bd119ad26e7bcf6db253696c5fdb84d6f95c02b1fed8a37378", "canonical_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "child_ids": ["xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:vif_list:same_as_site_region"], "collection_id": "xcsh-docs:resources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "parent_id": "xcsh-docs:resources:aws_vpc_site:properties:direct_connect_enabled:hosted_vifs", "path": "docs/guides/resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--vif_list.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs", "vif_list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_vpc_site/properties/direct_connect_enabled/hosted_vifs/vif_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "direct_connect_enabled.hosted_vifs.vif_list for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled.hosted_vifs.vif_list

Breadcrumbs:

- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
- [Property reference](resources--aws_vpc_site--reference.md)
- [direct_connect_enabled](resources--aws_vpc_site--properties--direct_connect_enabled.md)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs.md)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Hosted VIF Config. List of Hosted VIF Config.

Upstream description:

List of Hosted VIF Config.

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("vif_id"),
  validators.ConflictingListObjectAttributes("other_region",
    "same_as_site_region")}
```

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 30,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 30,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "30",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
vif_list {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-direct_connect_enabled--hosted_vifs--vif_list--other_region"></a>

### other_region property

Type: `"string"`. Optional.

\[Enum:
af-south-1|ap-east-1|ap-northeast-1|ap-northeast-2|ap-south-1|ap-southeast-1|ap-southeast-2|ap-southeast-3|ca-central-1|eu-central-1|eu-north-1|eu-south-1|eu-west-1|eu-west-2|eu-west-3|me-south-1|sa-east-1|us-east-1|us-east-2|us-west-1|us-west-2\]
Exclusive with \[same\_as\_site\_region\] Other Region. Possible values are \`af-south-1\`,
\`ap-east-1\`, \`ap-northeast-1\`, \`ap-northeast-2\`, \`ap-south-1\`, \`ap-southeast-1\`,
\`ap-southeast-2\`, \`ap-southeast-3\`, \`ca-central-1\`, \`eu-central-1\`, \`eu-north-1\`,
\`eu-south-1\`, \`eu-west-1\`, \`eu-west-2\`, \`eu-west-3\`, \`me-south-1\`, \`sa-east-1\`,
\`us-east-1\`, \`us-east-2\`, \`us-west-1\`, \`us-west-2\`.

Upstream description:

Exclusive with \[same\_as\_site\_region\] Other Region.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "enum": [
    "af-south-1",
    "ap-east-1",
    "ap-northeast-1",
    "ap-northeast-2",
    "ap-south-1",
    "ap-southeast-1",
    "ap-southeast-2",
    "ap-southeast-3",
    "ca-central-1",
    "eu-central-1",
    "eu-north-1",
    "eu-south-1",
    "eu-west-1",
    "eu-west-2",
    "eu-west-3",
    "me-south-1",
    "sa-east-1",
    "us-east-1",
    "us-east-2",
    "us-west-1",
    "us-west-2"
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
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.in": "[\\\"af-south-1\\\",\\\"ap-east-1\\\",\\\"ap-northeast-1\\\",\\\"ap-northeast-2\\\",\\\"ap-south-1\\\",\\\"ap-southeast-1\\\",\\\"ap-southeast-2\\\",\\\"ap-southeast-3\\\",\\\"ca-central-1\\\",\\\"eu-central-1\\\",\\\"eu-north-1\\\",\\\"eu-south-1\\\",\\\"eu-west-1\\\",\\\"eu-west-2\\\",\\\"eu-west-3\\\",\\\"me-south-1\\\",\\\"sa-east-1\\\",\\\"us-east-1\\\",\\\"us-east-2\\\",\\\"us-west-1\\\",\\\"us-west-2\\\"]"
  }
}
```

- [same_as_site_region](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--vif_list--same_as_site_region.md): complete subsection reference.

<a id="schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id"></a>

### vif_id property

Type: `"string"`. Optional.

AWS Direct Connect VIF ID that needs to be connected to the site.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthAtMost(1024),
}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.pattern": "^(dxvif-)([a-z0-9]{8}|[a-z0-9]{17})$"
  }
}
```

## Next pages

- [direct_connect_enabled.hosted_vifs.vif_list.same_as_site_region](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs--vif_list--same_as_site_region.md)
- [direct_connect_enabled.hosted_vifs](resources--aws_vpc_site--properties--direct_connect_enabled--hosted_vifs.md)
- [xcsh_aws_vpc_site](../resources/aws_vpc_site.md)
