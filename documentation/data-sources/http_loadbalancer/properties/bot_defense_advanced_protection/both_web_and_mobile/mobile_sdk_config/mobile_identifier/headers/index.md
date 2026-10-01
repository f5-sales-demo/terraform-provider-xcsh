---
page_title: "bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers"
subcategory: "Load Balancing"
description: "bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 6240, "body_sha256": "sha256:50b4d170ff03b2eaba109c39af04d3602289987dd59f7e6c175b9e9c32e3ef29", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile_sdk_config:mobile_identifier:headers:check_not_present", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile_sdk_config:mobile_identifier:headers:check_present", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile_sdk_config:mobile_identifier:headers:item"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile_sdk_config:mobile_identifier:headers", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense_advanced_protection:both_web_and_mobile:mobile_sdk_config:mobile_identifier", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/mobile_identifier/headers/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["bot_defense_advanced_protection", "both_web_and_mobile", "mobile_sdk_config", "mobile_identifier", "headers"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/mobile_identifier/headers/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [bot_defense_advanced_protection](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/)
- [bot_defense_advanced_protection.both_web_and_mobile](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/mobile_identifier/)
- bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers

<a id="section"></a>

Type: `"list"`. Computed.

Headers that can be used to identify mobile traffic.

Receipt-pinned upstream constraints:

```json
{
  "maxItems": 32,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "array",
    "deterministic": true,
    "maxItems": 32,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minItems": 0,
    "uniqueItems": true
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.repeated.max_items": "32",
    "ves.io.schema.rules.repeated.unique": "true"
  }
}
```

## Direct properties

- [check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/mobile_identifier/headers/check_not_present/): complete subsection reference.

- [check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/mobile_identifier/headers/check_present/): complete subsection reference.

- [item](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/mobile_identifier/headers/item/): complete subsection reference.

<a id="schema-bot_defense_advanced_protection--both_web_and_mobile--mobile_sdk_config--mobile_identifier--headers--name"></a>

### name property

Type: `"string"`. Computed.

Header Name. A case-insensitive HTTP header name.

Upstream description:

A case-insensitive HTTP header name.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 256,
  "x-f5xc-constraints": {
    "byteLength": {
      "max": 256
    },
    "category": "discovery",
    "characterSet": {
      "allowed": "[a-z0-9-]",
      "description": "Lowercase letter start, alphanumeric with hyphens, alphanumeric end",
      "required": "[a-z0-9]",
      "restricted": "[^a-z0-9-]"
    },
    "constraintType": "string",
    "deterministic": true,
    "format": "dns-label",
    "formatDescription": "DNS-1035 label: must start with a lowercase letter, may contain lowercase alphanumeric and hyphens, must end with alphanumeric",
    "maxLength": 63,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    },
    "minLength": 1,
    "pattern": "^[a-z]([-a-z0-9]*[a-z0-9])?$",
    "validation": {
      "rfc": "RFC 1035",
      "standard": "DNS-1035 label (alpha-first)"
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
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.http_header_field": "true",
    "ves.io.schema.rules.string.max_bytes": "256"
  }
}
```

## Next pages

- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_not_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/mobile_identifier/headers/check_not_present/)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.check_present](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/mobile_identifier/headers/check_present/)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier.headers.item](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/mobile_identifier/headers/item/)
- [bot_defense_advanced_protection.both_web_and_mobile.mobile_sdk_config.mobile_identifier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense_advanced_protection/both_web_and_mobile/mobile_sdk_config/mobile_identifier/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
