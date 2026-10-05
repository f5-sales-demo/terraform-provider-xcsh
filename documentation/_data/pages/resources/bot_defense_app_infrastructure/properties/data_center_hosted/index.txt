---
page_title: "data_center_hosted"
subcategory: ""
description: "Infra F5 Hosted."
xcsh_docs: {"aliases": ["data center hosted"], "body_bytes": 4309, "body_sha256": "sha256:d00c221be8aeba72a6bbb10a4fcd1b7f70a40fd09662f75a955b2c580f1763c2", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:egress", "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted", "parent_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "path": "documentation/resources/bot_defense_app_infrastructure/properties/data_center_hosted/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1133100200313302-2010010001030132-2021321011023121-2030100303300000-2303100013223101-1012010312223212-3320321330320131-2132230031013001", "registry_path": "docs/guides/resources--bot_defense_app_infrastructure--reference--group-001.md", "relationships": [{"anchor": "schema-data_center_hosted--infra_host_name", "enforcement": "provider-schema", "group": "data_center_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "data_center_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:egress", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "data_center_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["data_center_hosted"], "schema_version": 1, "sections": [{"aliases": ["data center hosted egress"], "anchor": "section", "description": "Egress", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:egress", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-data_center_hosted--egress--ip_address", "enforcement": "provider-schema", "group": "data_center_hosted.egress:RequiredListObjectAttributes:ip_address", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:egress", "type": "requires"}], "schema_path": ["data_center_hosted", "egress"], "syntax": "block", "type": "object"}, {"aliases": ["data center hosted infra host name"], "anchor": "schema-data_center_hosted--infra_host_name", "description": "Infra Host Name.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted", "enum_extraction_complete": true, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_center_hosted", "infra_host_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["data center hosted ingress"], "anchor": "section", "description": "Ingress", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-data_center_hosted--ingress--host_name", "enforcement": "provider-schema", "group": "data_center_hosted.ingress:ConflictingListObjectAttributes:host_name,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress", "type": "conflicts"}, {"anchor": "schema-data_center_hosted--ingress--ip_address", "enforcement": "provider-schema", "group": "data_center_hosted.ingress:ConflictingListObjectAttributes:host_name,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress", "type": "conflicts"}], "schema_path": ["data_center_hosted", "ingress"], "syntax": "block", "type": "object"}, {"aliases": ["data center hosted region"], "anchor": "schema-data_center_hosted--region", "description": "Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:data_center_hosted", "enum_extraction_complete": true, "enum_validators": [{"case_sensitive": true, "complete": true, "source": "ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf", "validator": "OneOf", "values": ["ASIA", "EU", "US"], "version": 1}], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_center_hosted", "region"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/properties/data_center_hosted/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Infra F5 Hosted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data_center_hosted

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/)
- data_center_hosted

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

F5 Hosted. Infra F5 Hosted.

Upstream description:

Infra F5 Hosted.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("egress",
    "infra_host_name",
    "ingress")}
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

Terraform syntax:

```terraform
data_center_hosted {
  # Configure direct properties listed below.
}
```

## Direct properties

- [egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/egress/): complete subsection reference.

<a id="schema-data_center_hosted--infra_host_name"></a>

### infra_host_name property

Type: `"string"`. Optional.

Infra Host Name. Infra Host Name.

Upstream description:

Infra Host Name.

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
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-03T05:10:26+00:00"
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [ingress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/): complete subsection reference.

<a id="schema-data_center_hosted--region"></a>

### region property

Type: `"string"`. Optional.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense Advanced region

&#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: true
EnumValidators: [{"version":1,"validator":"OneOf","values":["ASIA","EU","US"],"case_sensitive":true,"complete":true,"source":"ast-validator:github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator.OneOf"}]
Validators: []validator.String{
  stringvalidator.OneOf("US",
    "EU",
    "ASIA"),
}
```

Receipt-pinned upstream constraints:

```json
{
  "default": "US",
  "enum": [
    "US",
    "EU",
    "ASIA"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

## Next pages

- [data_center_hosted.egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/egress/)
- [data_center_hosted.ingress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/)
- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/)
