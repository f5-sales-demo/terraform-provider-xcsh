---
page_title: "cloud_hosted"
subcategory: ""
description: "Infra F5 Hosted."
xcsh_docs: {"aliases": ["cloud hosted"], "body_bytes": 2956, "body_sha256": "sha256:114f9800ad2049ed4b99acc4d3f43213e302a6c76387f36ba327c64b81c5692c", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted:egress", "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted:ingress"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted", "parent_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:reference", "path": "documentation/data-sources/bot_defense_app_infrastructure/properties/cloud_hosted/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1103310101020302-1303233332110101-0231320222231323-0233221013020013-0333122021130002-1312103123303001-3021222211302231-2313310212322103", "registry_path": "docs/guides/data-sources--bot_defense_app_infrastructure--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["cloud_hosted"], "schema_version": 1, "sections": [{"aliases": ["cloud hosted egress"], "anchor": "section", "description": "Egress", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted:egress", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloud_hosted", "egress"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloud hosted infra host name"], "anchor": "schema-cloud_hosted--infra_host_name", "description": "Infra Host Name.", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloud_hosted", "infra_host_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["cloud hosted ingress"], "anchor": "section", "description": "Ingress", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted:ingress", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["cloud_hosted", "ingress"], "syntax": "attribute", "type": "object"}, {"aliases": ["cloud hosted region"], "anchor": "schema-cloud_hosted--region", "description": "Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region.", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["cloud_hosted", "region"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_defense_app_infrastructure/properties/cloud_hosted/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Infra F5 Hosted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloud_hosted

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/)
- cloud_hosted

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: cloud\_hosted, data\_center\_hosted\] F5 Hosted. Infra F5 Hosted.

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

- [cloud_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/cloud_hosted/#section)
- [data_center_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/cloud_hosted/egress/): complete subsection reference.

<a id="schema-cloud_hosted--infra_host_name"></a>

### infra_host_name property

Type: `"string"`. Computed.

Infra Host Name. Infra Host Name.

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
      "validatedAt": "2026-10-06T12:36:10+00:00"
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

- [ingress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/cloud_hosted/ingress/): complete subsection reference.

<a id="schema-cloud_hosted--region"></a>

### region property

Type: `"string"`. Computed.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

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
