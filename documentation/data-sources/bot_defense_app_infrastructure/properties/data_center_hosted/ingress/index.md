---
page_title: "data_center_hosted.ingress"
subcategory: ""
description: "Ingress"
xcsh_docs: {"aliases": ["data center hosted ingress"], "body_bytes": 7683, "body_sha256": "sha256:4f35842df1baecf88dda027f31e50461fc0b0b5d623140b9fff211dc57a1f03a", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress", "parent_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted", "path": "documentation/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/index.md", "product": "distributed-cloud", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2332132211113123-0320021203313232-3302130030312112-3002010012133101-1132123033332312-1022033303130002-2000100333021131-2121111210222313", "registry_path": "docs/guides/data-sources--bot_defense_app_infrastructure--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["data_center_hosted", "ingress"], "schema_version": 1, "sections": [{"aliases": ["data center hosted ingress host name"], "anchor": "schema-data_center_hosted--ingress--host_name", "description": "Exclusive with Ingress Host Name.", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_center_hosted", "ingress", "host_name"], "syntax": "attribute", "type": "string"}, {"aliases": ["data center hosted ingress ip address"], "anchor": "schema-data_center_hosted--ingress--ip_address", "description": "Exclusive with Ingress IP Address.", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_center_hosted", "ingress", "ip_address"], "syntax": "attribute", "type": "string"}, {"aliases": ["data center hosted ingress location"], "anchor": "schema-data_center_hosted--ingress--location", "description": "Region location AWS_AP_NORTHEAST_1 AWS_AP_NORTHEAST_3 AWS_AP_SOUTH_1 AWS_AP_SOUTH_2 AWS_AP_SOUTHEAST_1 AWS_AP_SOUTHEAST_2 AWS_AP_SOUTHEAST_3 AWS_EU_CENTRAL_1 AWS_EU_NORTH_1 AWS_EU_WEST_1 AWS_ME_SOUTH_1 AWS_SA_EAST_1 AWS_US_EAST_1 AWS_US_EAST_2 AWS_US_WEST_1 AWS_US_WEST_2 GCP_ASIA_EAST_1 GCP_ASIA_EAST_2", "document_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["data_center_hosted", "ingress", "location"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Ingress", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# data_center_hosted.ingress

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/)
- [data_center_hosted](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/)
- data_center_hosted.ingress

<a id="section"></a>

Type: `"list"`. Computed.

Ingress. Ingress

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 3,
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 3,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minItems": 1,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.max_items": "3",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

<a id="schema-data_center_hosted--ingress--host_name"></a>

### host_name property

Type: `"string"`. Computed.

Exclusive with \[ip\_address\] Ingress Host Name.

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
      "validatedAt": "2026-10-09T12:34:59+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.hostname": "true"
  }
}
```

<a id="schema-data_center_hosted--ingress--ip_address"></a>

### ip_address property

Type: `"string"`. Computed.

Exclusive with \[host\_name\] Ingress IP Address.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "ip",
    "formatDescription": "IPv4 dotted-decimal notation (e.g., 192.168.1.1)",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-09T12:34:59+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-data_center_hosted--ingress--location"></a>

### location property

Type: `"string"`. Computed.

\[Enum:
AWS\_AP\_NORTHEAST\_1|AWS\_AP\_NORTHEAST\_3|AWS\_AP\_SOUTH\_1|AWS\_AP\_SOUTH\_2|AWS\_AP\_SOUTHEAST\_1|AWS\_AP\_SOUTHEAST\_2|AWS\_AP\_SOUTHEAST\_3|AWS\_EU\_CENTRAL\_1|AWS\_EU\_NORTH\_1|AWS\_EU\_WEST\_1|AWS\_ME\_SOUTH\_1|AWS\_SA\_EAST\_1|AWS\_US\_EAST\_1|AWS\_US\_EAST\_2|AWS\_US\_WEST\_1|AWS\_US\_WEST\_2|GCP\_ASIA\_EAST\_1|GCP\_ASIA\_EAST\_2|GCP\_ASIA\_NORTHEAST\_1|GCP\_ASIA\_NORTHEAST\_2|GCP\_ASIA\_NORTHEAST\_3|GCP\_ASIA\_SOUTH\_1|GCP\_ASIA\_SOUTHEAST\_1|GCP\_ASIA\_SOUTHEAST\_2|GCP\_AUSTRALIA\_SOUTHEAST\_1|GCP\_EUROPE\_WEST\_1|GCP\_EUROPE\_WEST\_2|GCP\_EUROPE\_WEST\_3|GCP\_NORTHAMERICA\_NORTHEAST\_1|GCP\_NORTHAMERICA\_NORTHEAST\_2|GCP\_SOUTHAMERICA\_EAST\_1|GCP\_SOUTHAMERICA\_WEST\_1|GCP\_US\_CENTRAL\_1|GCP\_US\_EAST\_1|GCP\_US\_EAST\_4|GCP\_US\_WEST\_1|GCP\_US\_WEST\_2\]
Region location AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1.. Possible values are
\`AWS\_AP\_NORTHEAST\_1\`, \`AWS\_AP\_NORTHEAST\_3\`, \`AWS\_AP\_SOUTH\_1\`, \`AWS\_AP\_SOUTH\_2\`,
\`AWS\_AP\_SOUTHEAST\_1\`, \`AWS\_AP\_SOUTHEAST\_2\`, \`AWS\_AP\_SOUTHEAST\_3\`,
\`AWS\_EU\_CENTRAL\_1\`, \`AWS\_EU\_NORTH\_1\`, \`AWS\_EU\_WEST\_1\`, \`AWS\_ME\_SOUTH\_1\`,
\`AWS\_SA\_EAST\_1\`, \`AWS\_US\_EAST\_1\`, \`AWS\_US\_EAST\_2\`, \`AWS\_US\_WEST\_1\`,
\`AWS\_US\_WEST\_2\`, \`GCP\_ASIA\_EAST\_1\`, \`GCP\_ASIA\_EAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_1\`,
\`GCP\_ASIA\_NORTHEAST\_2\`, \`GCP\_ASIA\_NORTHEAST\_3\`, \`GCP\_ASIA\_SOUTH\_1\`,
\`GCP\_ASIA\_SOUTHEAST\_1\`, \`GCP\_ASIA\_SOUTHEAST\_2\`, \`GCP\_AUSTRALIA\_SOUTHEAST\_1\`,
\`GCP\_EUROPE\_WEST\_1\`, \`GCP\_EUROPE\_WEST\_2\`, \`GCP\_EUROPE\_WEST\_3\`,
\`GCP\_NORTHAMERICA\_NORTHEAST\_1\`, \`GCP\_NORTHAMERICA\_NORTHEAST\_2\`,
\`GCP\_SOUTHAMERICA\_EAST\_1\`, \`GCP\_SOUTHAMERICA\_WEST\_1\`, \`GCP\_US\_CENTRAL\_1\`,
\`GCP\_US\_EAST\_1\`, \`GCP\_US\_EAST\_4\`, \`GCP\_US\_WEST\_1\`, \`GCP\_US\_WEST\_2\`. Defaults to
\`AWS\_AP\_NORTHEAST\_1\`.

Additional upstream details:

Region location

AWS\_AP\_NORTHEAST\_1 AWS\_AP\_NORTHEAST\_3 AWS\_AP\_SOUTH\_1 AWS\_AP\_SOUTH\_2
AWS\_AP\_SOUTHEAST\_1 AWS\_AP\_SOUTHEAST\_2 AWS\_AP\_SOUTHEAST\_3 AWS\_EU\_CENTRAL\_1
AWS\_EU\_NORTH\_1 AWS\_EU\_WEST\_1 AWS\_ME\_SOUTH\_1 AWS\_SA\_EAST\_1 AWS\_US\_EAST\_1
AWS\_US\_EAST\_2 AWS\_US\_WEST\_1 AWS\_US\_WEST\_2 GCP\_ASIA\_EAST\_1 GCP\_ASIA\_EAST\_2
GCP\_ASIA\_NORTHEAST\_1 GCP\_ASIA\_NORTHEAST\_2 GCP\_ASIA\_NORTHEAST\_3 GCP\_ASIA\_SOUTH\_1
GCP\_ASIA\_SOUTHEAST\_1 GCP\_ASIA\_SOUTHEAST\_2 GCP\_AUSTRALIA\_SOUTHEAST\_1 GCP\_EUROPE\_WEST\_1
GCP\_EUROPE\_WEST\_2 GCP\_EUROPE\_WEST\_3 GCP\_NORTHAMERICA\_NORTHEAST\_1
GCP\_NORTHAMERICA\_NORTHEAST\_2 GCP\_SOUTHAMERICA\_EAST\_1 GCP\_SOUTHAMERICA\_WEST\_1
GCP\_US\_CENTRAL\_1 GCP\_US\_EAST\_1 GCP\_US\_EAST\_4 GCP\_US\_WEST\_1 GCP\_US\_WEST\_2.

Receipt-pinned upstream constraints:

```json
{
  "default": "AWS_AP_NORTHEAST_1",
  "enum": [
    "AWS_AP_NORTHEAST_1",
    "AWS_AP_NORTHEAST_3",
    "AWS_AP_SOUTH_1",
    "AWS_AP_SOUTH_2",
    "AWS_AP_SOUTHEAST_1",
    "AWS_AP_SOUTHEAST_2",
    "AWS_AP_SOUTHEAST_3",
    "AWS_EU_CENTRAL_1",
    "AWS_EU_NORTH_1",
    "AWS_EU_WEST_1",
    "AWS_ME_SOUTH_1",
    "AWS_SA_EAST_1",
    "AWS_US_EAST_1",
    "AWS_US_EAST_2",
    "AWS_US_WEST_1",
    "AWS_US_WEST_2",
    "GCP_ASIA_EAST_1",
    "GCP_ASIA_EAST_2",
    "GCP_ASIA_NORTHEAST_1",
    "GCP_ASIA_NORTHEAST_2",
    "GCP_ASIA_NORTHEAST_3",
    "GCP_ASIA_SOUTH_1",
    "GCP_ASIA_SOUTHEAST_1",
    "GCP_ASIA_SOUTHEAST_2",
    "GCP_AUSTRALIA_SOUTHEAST_1",
    "GCP_EUROPE_WEST_1",
    "GCP_EUROPE_WEST_2",
    "GCP_EUROPE_WEST_3",
    "GCP_NORTHAMERICA_NORTHEAST_1",
    "GCP_NORTHAMERICA_NORTHEAST_2",
    "GCP_SOUTHAMERICA_EAST_1",
    "GCP_SOUTHAMERICA_WEST_1",
    "GCP_US_CENTRAL_1",
    "GCP_US_EAST_1",
    "GCP_US_EAST_4",
    "GCP_US_WEST_1",
    "GCP_US_WEST_2"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```
