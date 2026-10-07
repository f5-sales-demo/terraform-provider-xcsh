---
page_title: "data_center_hosted"
subcategory: ""
description: "Infra F5 Hosted."
xcsh_docs: {"aliases": ["data center hosted"], "body_bytes": 2525, "body_sha256": "sha256:bff910b21d8dfdc232c6ecd0cc81669ba3dd511651eceed96f7166f52edfb91a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted:egress", "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted", "parent_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:reference", "path": "documentation/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-3020301332100300-3133221002233332-3131022122123002-3023220030220010-3123110321022223-3103033023022121-0002110311100232-1102230222113001", "registry_path": "docs/guides/data-sources--bot_defense_app_infrastructure--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["data_center_hosted"], "schema_version": 1, "sections": [{"aliases": ["data center hosted egress"], "anchor": "section", "description": "Egress", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted:egress", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["data_center_hosted", "egress"], "syntax": "attribute", "type": "object"}, {"aliases": ["data center hosted infra host name"], "anchor": "schema-data_center_hosted--infra_host_name", "description": "Infra Host Name.", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_center_hosted", "infra_host_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["data center hosted ingress"], "anchor": "section", "description": "Ingress", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["data_center_hosted", "ingress"], "syntax": "attribute", "type": "object"}, {"aliases": ["data center hosted region"], "anchor": "schema-data_center_hosted--region", "description": "Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU European Union region - ASIA: ASIA Asia region.", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_center_hosted", "region"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Infra F5 Hosted.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data_center_hosted

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/)
- data_center_hosted

<a id="section"></a>

Type: `"single"`. Computed.

F5 Hosted. Infra F5 Hosted.

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

- [egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/egress/): complete subsection reference.

<a id="schema-data_center_hosted--infra_host_name"></a>

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
      "validatedAt": "2026-10-07T18:46:18+00:00"
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

- [ingress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/): complete subsection reference.

<a id="schema-data_center_hosted--region"></a>

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
