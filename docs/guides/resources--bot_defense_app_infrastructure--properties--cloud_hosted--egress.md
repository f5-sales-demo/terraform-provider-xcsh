---
page_title: "cloud_hosted.egress"
subcategory: ""
description: "cloud_hosted.egress for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 8368, "body_sha256": "sha256:6c2cf9dd2c52792bb51ab65af72e5b1c0375ff37298a66e6dfa00e30b27d181b", "canonical_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:egress", "child_ids": [], "collection_id": "xcsh-docs:resources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted:egress", "parent_id": "xcsh-docs:resources:bot_defense_app_infrastructure:properties:cloud_hosted", "path": "docs/guides/resources--bot_defense_app_infrastructure--properties--cloud_hosted--egress.md", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloud_hosted", "egress"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bot_defense_app_infrastructure/properties/cloud_hosted/egress/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloud_hosted.egress for xcsh_bot_defense_app_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloud_hosted.egress

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md)
- [Property reference](resources--bot_defense_app_infrastructure--reference.md)
- [cloud_hosted](resources--bot_defense_app_infrastructure--properties--cloud_hosted.md)
- cloud_hosted.egress

<a id="section"></a>

Type: `"object"`. list nested block, Optional.

Egress. Egress

Upstream description:

Egress

Provider validators and defaults (from schema source):

```go
Validators: []validator.List{validators.RequiredListObjectAttributes("ip_address")}
```

Receipt-pinned upstream constraints:

```json
{
  "minItems": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
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
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.repeated.min_items": "1",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

Terraform syntax:

```terraform
egress {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-cloud_hosted--egress--ip_address"></a>

### ip_address property

Type: `"string"`. Optional.

IP Address. Egress IP address.

Upstream description:

Egress IP address.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(7, 1024),
  validators.IPValidator(),
}
```

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 7,
    "pattern": "^((25[0-5]|(2[0-4]|1\\d|[1-9]|)\\d)\\.?\\b){4}$"
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.ip": "true"
  }
}
```

<a id="schema-cloud_hosted--egress--location"></a>

### location property

Type: `"string"`. Optional.

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

Upstream description:

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

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.OneOf("AWS_AP_NORTHEAST_1",
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
    "GCP_US_WEST_2"),
}
```

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

## Next pages

- [cloud_hosted](resources--bot_defense_app_infrastructure--properties--cloud_hosted.md)
- [xcsh_bot_defense_app_infrastructure](../resources/bot_defense_app_infrastructure.md)
