---
page_title: "cloud_hosted"
subcategory: ""
description: "Infra F5 Hosted."
xcsh_docs: {"aliases": ["cloud hosted"], "body_bytes": 4392, "body_sha256": "sha256:aa936cfe3488273caa1047a3e07fa9d2e86360dea7e33a3c43ecee933e432e4f", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:egress", "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:ingress"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted", "parent_id": "xcsh-docs:resources:bot_defense_app_infrastructure:reference", "path": "documentation/resources/bot_defense_app_infrastructure/properties/cloud_hosted/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0203000233203130-2220133113021311-0112223313212301-0002122322321023-1010110301202032-2030213223220002-0131021011102110-3122333203011333", "registry_path": "docs/guides/resources--bot_defense_app_infrastructure--reference--group-001.md", "relationships": [{"anchor": "schema-cloud_hosted--infra_host_name", "enforcement": "provider-schema", "group": "cloud_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloud_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:egress", "type": "requires"}, {"anchor": "section", "enforcement": "provider-schema", "group": "cloud_hosted:RequiredObjectAttributes:egress,infra_host_name,ingress", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:ingress", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["cloud_hosted"], "schema_version": 1, "sections": [{"aliases": ["egress"], "anchor": "section", "description": "Egress", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:egress", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-cloud_hosted--egress--ip_address", "enforcement": "provider-schema", "group": "cloud_hosted.egress:RequiredListObjectAttributes:ip_address", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:egress", "type": "requires"}], "schema_path": ["cloud_hosted", "egress"], "syntax": "block", "type": "object"}, {"aliases": ["infra host name"], "anchor": "schema-cloud_hosted--infra_host_name", "description": "Infra Host Name.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloud_hosted", "infra_host_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["ingress"], "anchor": "section", "description": "Ingress", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:ingress", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-cloud_hosted--ingress--host_name", "enforcement": "provider-schema", "group": "cloud_hosted.ingress:ConflictingListObjectAttributes:host_name,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:ingress", "type": "conflicts"}, {"anchor": "schema-cloud_hosted--ingress--ip_address", "enforcement": "provider-schema", "group": "cloud_hosted.ingress:ConflictingListObjectAttributes:host_name,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:ingress", "type": "conflicts"}], "schema_path": ["cloud_hosted", "ingress"], "syntax": "block", "type": "object"}, {"aliases": ["region"], "anchor": "schema-cloud_hosted--region", "description": "Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region.", "document_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloud_hosted", "region"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/properties/cloud_hosted/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Infra F5 Hosted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloud_hosted

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/)
- cloud_hosted

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: cloud\_hosted, data\_center\_hosted\] F5 Hosted. Infra F5 Hosted.

Upstream description:

Infra F5 Hosted.

Provider validators and defaults (from schema source):

```go
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

OneOf alternatives in this subsection:

- [cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/#section)
- [data_center_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/data_center_hosted/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
cloud_hosted {
  # Configure direct properties listed below.
}
```

## Direct properties

- [egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/egress/): complete subsection reference.

<a id="schema-cloud_hosted--infra_host_name"></a>

### infra_host_name property

Type: `"string"`. Optional.

Infra Host Name. Infra Host Name.

Upstream description:

Infra Host Name.

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
    "format": "hostname",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
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
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

- [ingress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/ingress/): complete subsection reference.

<a id="schema-cloud_hosted--region"></a>

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

- [cloud_hosted.egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/egress/)
- [cloud_hosted.ingress](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/cloud_hosted/ingress/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/properties/)
- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bot_defense_app_infrastructure/)
