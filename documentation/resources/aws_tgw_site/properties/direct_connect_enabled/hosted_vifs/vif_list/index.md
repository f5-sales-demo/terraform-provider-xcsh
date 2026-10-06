---
page_title: "direct_connect_enabled.hosted_vifs.vif_list"
subcategory: ""
description: "List of Hosted VIF Config."
xcsh_docs: {"aliases": ["direct connect enabled hosted vifs vif list"], "body_bytes": 7058, "body_sha256": "sha256:50f87c7ff02d41e99f1c74e9ab349b4c21a025b3e26f51b02c2e6284f1325189", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list:same_as_site_region"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs", "path": "documentation/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0011221312000023-0333022123120120-2230132022303332-2133231000121113-0201013333223313-1013022112012032-3330202203123231-0023030233200303", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [{"anchor": "schema-direct_connect_enabled--hosted_vifs--vif_list--other_region", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.vif_list:ConflictingListObjectAttributes:other_region,same_as_site_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.vif_list:ConflictingListObjectAttributes:other_region,same_as_site_region", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list:same_as_site_region", "type": "conflicts"}, {"anchor": "schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id", "enforcement": "provider-schema", "group": "direct_connect_enabled.hosted_vifs.vif_list:RequiredListObjectAttributes:vif_id", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["direct_connect_enabled", "hosted_vifs", "vif_list"], "schema_version": 1, "sections": [{"aliases": ["direct connect enabled hosted vifs vif list other region"], "anchor": "schema-direct_connect_enabled--hosted_vifs--vif_list--other_region", "description": "Exclusive with Other Region.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["af-south-1", "ap-east-1", "ap-northeast-1", "ap-northeast-2", "ap-south-1", "ap-southeast-1", "ap-southeast-2", "ap-southeast-3", "ca-central-1", "eu-central-1", "eu-north-1", "eu-south-1", "eu-west-1", "eu-west-2", "eu-west-3", "me-south-1", "sa-east-1", "us-east-1", "us-east-2", "us-west-1", "us-west-2"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "vif_list", "other_region"], "syntax": "attribute", "type": "string"}, {"aliases": ["direct connect enabled hosted vifs vif list same as site region"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list:same_as_site_region", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "vif_list", "same_as_site_region"], "syntax": "attribute", "type": "object"}, {"aliases": ["direct connect enabled hosted vifs vif list vif id"], "anchor": "schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id", "description": "AWS Direct Connect VIF ID that needs to be connected to the site.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:direct_connect_enabled:hosted_vifs:vif_list", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["direct_connect_enabled", "hosted_vifs", "vif_list", "vif_id"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "List of Hosted VIF Config.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# direct_connect_enabled.hosted_vifs.vif_list

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [direct_connect_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/)
- [direct_connect_enabled.hosted_vifs](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/)
- direct_connect_enabled.hosted_vifs.vif_list

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

List of Hosted VIF Config. List of Hosted VIF Config.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["af-south-1","ap-east-1","ap-northeast-1","ap-northeast-2","ap-south-1","ap-southeast-1","ap-southeast-2","ap-southeast-3","ca-central-1","eu-central-1","eu-north-1","eu-south-1","eu-west-1","eu-west-2","eu-west-3","me-south-1","sa-east-1","us-east-1","us-east-2","us-west-1","us-west-2"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [same_as_site_region](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/direct_connect_enabled/hosted_vifs/vif_list/same_as_site_region/): complete subsection reference.

<a id="schema-direct_connect_enabled--hosted_vifs--vif_list--vif_id"></a>

### vif_id property

Type: `"string"`. Optional.

AWS Direct Connect VIF ID that needs to be connected to the site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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
