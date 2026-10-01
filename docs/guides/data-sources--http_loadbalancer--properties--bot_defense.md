---
page_title: "bot_defense"
subcategory: "Load Balancing"
description: "bot_defense for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 3947, "body_sha256": "sha256:bcee60209637d09008ed16de68ee6ae6a056227b69b11d7204f7d97cb6a822d4", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:disable_cors_support", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:enable_cors_support", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- bot_defense

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: bot\_defense, bot\_defense\_advanced\_protection, disable\_bot\_defense; Default:
disable\_bot\_defense\] Defines various configuration OPTIONS for Bot Defense Policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense Policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-cors_support_choice": "[\"disable_cors_support\",\"enable_cors_support\"]"
}
```

OneOf alternatives in this subsection:

- [bot_defense](data-sources--http_loadbalancer--properties--bot_defense.md#section)
- [bot_defense_advanced_protection](data-sources--http_loadbalancer--properties--bot_defense_advanced_protection.md#section)
- [disable_bot_defense](data-sources--http_loadbalancer--properties--disable_bot_defense.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [disable_cors_support](data-sources--http_loadbalancer--properties--bot_defense--disable_cors_support.md): complete subsection reference.

- [enable_cors_support](data-sources--http_loadbalancer--properties--bot_defense--enable_cors_support.md): complete subsection reference.

- [policy](data-sources--http_loadbalancer--properties--bot_defense--policy.md): complete subsection reference.

<a id="schema-bot_defense--regional_endpoint"></a>

### regional_endpoint property

Type: `"string"`. Computed.

\[Enum: AUTO|US|EU|ASIA\] Defines a selection for Bot Defense region - AUTO: AUTO Automatic
selection based on client IP address - US: US US region - EU: EU European Union region - ASIA: ASIA
Asia region. Possible values are \`AUTO\`, \`US\`, \`EU\`, \`ASIA\`. Defaults to \`AUTO\`.

Upstream description:

Defines a selection for Bot Defense region

&#8203;- AUTO: AUTO

Automatic selection based on client IP address &#8203;- US: US

US region &#8203;- EU: EU

European Union region &#8203;- ASIA: ASIA

Asia region.

Receipt-pinned upstream constraints:

```json
{
  "default": "AUTO",
  "enum": [
    "AUTO",
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

<a id="schema-bot_defense--timeout"></a>

### timeout property

Type: `"number"`. Computed.

The timeout for the inference check, in milliseconds.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 60000,
    "metadata": {
      "confidence": 0.99,
      "source": "api-probed",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minimum": 0,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.gte": "0",
    "ves.io.schema.rules.uint32.lte": "60000"
  }
}
```

## Next pages

- [bot_defense.disable_cors_support](data-sources--http_loadbalancer--properties--bot_defense--disable_cors_support.md)
- [bot_defense.enable_cors_support](data-sources--http_loadbalancer--properties--bot_defense--enable_cors_support.md)
- [bot_defense.policy](data-sources--http_loadbalancer--properties--bot_defense--policy.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
