---
page_title: "data_center_hosted"
subcategory: ""
description: "data_center_hosted for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 3418, "body_sha256": "sha256:9df713dc2b3c3eb7b2e8e05771032b4e9aafcc23cd1875dbe5f0214cd2c038f4", "child_ids": ["xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted:egress", "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted:ingress"], "collection_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:data_center_hosted", "parent_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:reference", "path": "documentation/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/index.md", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["data_center_hosted"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "data_center_hosted for xcsh_bot_defense_app_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

Infra F5 Hosted.

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

Upstream description:

Infra Host Name.

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

- [ingress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/): complete subsection reference.

<a id="schema-data_center_hosted--region"></a>

### region property

Type: `"string"`. Computed.

\[Enum: US|EU|ASIA\] Defines a selection for Bot Defense Advanced region - US: US US region - EU: EU
European Union region - ASIA: ASIA Asia region. Possible values are \`US\`, \`EU\`, \`ASIA\`.
Defaults to \`US\`.

Upstream description:

Defines a selection for Bot Defense Advanced region

&#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

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

- [data_center_hosted.egress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/egress/)
- [data_center_hosted.ingress](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/data_center_hosted/ingress/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/properties/)
- [xcsh_bot_defense_app_infrastructure](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_defense_app_infrastructure/)
