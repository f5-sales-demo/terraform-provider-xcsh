---
page_title: "cloud_hosted"
subcategory: ""
description: "cloud_hosted for xcsh_bot_defense_app_infrastructure."
xcsh_docs: {"aliases": [], "body_bytes": 3315, "body_sha256": "sha256:dfe360a7f53497c6b52d4691e699cbdbcd4d7f61d0c60d7abbb52a79d4053927", "canonical_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted", "child_ids": ["xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted:egress", "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted:ingress"], "collection_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:properties:cloud_hosted", "parent_id": "xcsh-docs:data-sources:bot_defense_app_infrastructure:reference", "path": "docs/guides/data-sources--bot_defense_app_infrastructure--properties--cloud_hosted.md", "provider_name": "bot_defense_app_infrastructure", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloud_hosted"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_defense_app_infrastructure/properties/cloud_hosted/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloud_hosted for xcsh_bot_defense_app_infrastructure.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bot_defense_app_infrastructureCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloud_hosted

Breadcrumbs:

- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference.md)
- cloud_hosted

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: cloud\_hosted, data\_center\_hosted\] F5 Hosted. Infra F5 Hosted.

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

OneOf alternatives in this subsection:

- [cloud_hosted](data-sources--bot_defense_app_infrastructure--properties--cloud_hosted.md#section)
- [data_center_hosted](data-sources--bot_defense_app_infrastructure--properties--data_center_hosted.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [egress](data-sources--bot_defense_app_infrastructure--properties--cloud_hosted--egress.md): complete subsection reference.

<a id="schema-cloud_hosted--infra_host_name"></a>

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

- [ingress](data-sources--bot_defense_app_infrastructure--properties--cloud_hosted--ingress.md): complete subsection reference.

<a id="schema-cloud_hosted--region"></a>

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

- [cloud_hosted.egress](data-sources--bot_defense_app_infrastructure--properties--cloud_hosted--egress.md)
- [cloud_hosted.ingress](data-sources--bot_defense_app_infrastructure--properties--cloud_hosted--ingress.md)
- [Property reference](data-sources--bot_defense_app_infrastructure--reference.md)
- [xcsh_bot_defense_app_infrastructure](../data-sources/bot_defense_app_infrastructure.md)
