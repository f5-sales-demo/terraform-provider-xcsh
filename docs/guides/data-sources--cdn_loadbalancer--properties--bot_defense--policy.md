---
page_title: "bot_defense.policy"
subcategory: "Load Balancing"
description: "bot_defense.policy for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 5447, "body_sha256": "sha256:6f278a526c5c513e88a79f707dc5026e3c7d0284a4ff416657c64a57ca62611c", "canonical_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:disable_js_insert", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:disable_mobile_sdk", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense", "path": "docs/guides/data-sources--cdn_loadbalancer--properties--bot_defense--policy.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
- [Property reference](data-sources--cdn_loadbalancer--reference.md)
- [bot_defense](data-sources--cdn_loadbalancer--properties--bot_defense.md)
- bot_defense.policy

<a id="section"></a>

Type: `"single"`. Computed.

Defines various configuration OPTIONS for Bot Defense policy.

Upstream description:

This defines various configuration OPTIONS for Bot Defense policy.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-java_script_choice": "[\"disable_js_insert\",\"js_insert_all_pages\",\"js_insert_all_pages_except\",\"js_insertion_rules\"]",
  "x-ves-oneof-field-mobile_sdk_choice": "[\"disable_mobile_sdk\",\"mobile_sdk_config\"]"
}
```

## Direct properties

- [disable_js_insert](data-sources--cdn_loadbalancer--properties--bot_defense--policy--disable_js_insert.md): complete subsection reference.

- [disable_mobile_sdk](data-sources--cdn_loadbalancer--properties--bot_defense--policy--disable_mobile_sdk.md): complete subsection reference.

<a id="schema-bot_defense--policy--javascript_mode"></a>

### javascript_mode property

Type: `"string"`. Computed.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Upstream description:

Web Client JavaScript Mode.

Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable
Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is non-cacheable Bot
Defense JavaScript for telemetry collection is requested synchronously, and it is cacheable.

Receipt-pinned upstream constraints:

```json
{
  "default": "ASYNC_JS_NO_CACHING",
  "enum": [
    "ASYNC_JS_NO_CACHING",
    "ASYNC_JS_CACHING",
    "SYNC_JS_NO_CACHING",
    "SYNC_JS_CACHING"
  ],
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

<a id="schema-bot_defense--policy--js_download_path"></a>

### js_download_path property

Type: `"string"`. Computed.

Customize Bot Defense Client JavaScript path. If not specified, default

Upstream description:

Customize Bot Defense Client JavaScript path. If not specified, default \`/common.js\`

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "general",
    "constraintType": "string",
    "maxLength": 1024,
    "metadata": {
      "confidence": 0.85,
      "source": "inferred",
      "validatedAt": "2026-09-29T03:20:54+00:00"
    }
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insert_all_pages.md): complete subsection reference.

- [js_insert_all_pages_except](data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insert_all_pages_except.md): complete subsection reference.

- [js_insertion_rules](data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules.md): complete subsection reference.

- [mobile_sdk_config](data-sources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config.md): complete subsection reference.

- [protected_app_endpoints](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md): complete subsection reference.

## Next pages

- [bot_defense.policy.disable_js_insert](data-sources--cdn_loadbalancer--properties--bot_defense--policy--disable_js_insert.md)
- [bot_defense.policy.disable_mobile_sdk](data-sources--cdn_loadbalancer--properties--bot_defense--policy--disable_mobile_sdk.md)
- [bot_defense.policy.js_insert_all_pages](data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insert_all_pages.md)
- [bot_defense.policy.js_insert_all_pages_except](data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insert_all_pages_except.md)
- [bot_defense.policy.js_insertion_rules](data-sources--cdn_loadbalancer--properties--bot_defense--policy--js_insertion_rules.md)
- [bot_defense.policy.mobile_sdk_config](data-sources--cdn_loadbalancer--properties--bot_defense--policy--mobile_sdk_config.md)
- [bot_defense.policy.protected_app_endpoints](data-sources--cdn_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense](data-sources--cdn_loadbalancer--properties--bot_defense.md)
- [xcsh_cdn_loadbalancer](../data-sources/cdn_loadbalancer.md)
