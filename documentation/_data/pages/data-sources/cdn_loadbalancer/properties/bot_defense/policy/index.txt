---
page_title: "bot_defense.policy"
subcategory: "Load Balancing"
description: "This defines various configuration OPTIONS for Bot Defense policy."
xcsh_docs: {"aliases": ["bot defense policy"], "body_bytes": 4680, "body_sha256": "sha256:c83e6878423b1b81a6d2194ccba46044d0a1093a05eb6088a994cd27d037402c", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:disable_js_insert", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:disable_mobile_sdk", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense", "path": "documentation/data-sources/cdn_loadbalancer/properties/bot_defense/policy/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2001112201103000-1012012101030300-1031303312100103-1132002211130320-1231133203111310-2302021103110210-2021131010220131-0120120110311111", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy disable js insert"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:disable_js_insert", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "disable_js_insert"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy disable mobile sdk"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:disable_mobile_sdk", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "disable_mobile_sdk"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy javascript mode"], "anchor": "schema-bot_defense--policy--javascript_mode", "description": "Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested synchronously, and it is", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "javascript_mode"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot defense policy js download path"], "anchor": "schema-bot_defense--policy--js_download_path", "description": "Customize Bot Defense Client JavaScript path. If not specified, default `/common.js`", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "js_download_path"], "syntax": "attribute", "type": "string"}, {"aliases": ["bot defense policy js insert all pages"], "anchor": "section", "description": "Insert Bot Defense JavaScript in all pages.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "js_insert_all_pages"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy js insert all pages except"], "anchor": "section", "description": "Insert Bot Defense JavaScript in all pages with the exceptions.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insert_all_pages_except", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "js_insert_all_pages_except"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy js insertion rules"], "anchor": "section", "description": "This defines custom JavaScript insertion rules for Bot Defense Policy.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:js_insertion_rules", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "js_insertion_rules"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy mobile sdk config"], "anchor": "section", "description": "Mobile SDK configuration.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "mobile_sdk_config"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy protected app endpoints"], "anchor": "section", "description": "List of protected endpoints. Limit: Approx '128 endpoints per Load Balancer (LB)' upto 4 LBs, '32 endpoints per LB' after 4 LBs.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This defines various configuration OPTIONS for Bot Defense policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/)
- bot_defense.policy

<a id="section"></a>

Type: `"single"`. Computed.

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

- [disable_js_insert](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/disable_js_insert/): complete subsection reference.

- [disable_mobile_sdk](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/disable_mobile_sdk/): complete subsection reference.

<a id="schema-bot_defense--policy--javascript_mode"></a>

### javascript_mode property

Type: `"string"`. Computed.

\[Enum: ASYNC\_JS\_NO\_CACHING|ASYNC\_JS\_CACHING|SYNC\_JS\_NO\_CACHING|SYNC\_JS\_CACHING\] Web
Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested asynchronously,
and it is cacheable Bot Defense JavaScript for telemetry collection is requested.. Possible values
are \`ASYNC\_JS\_NO\_CACHING\`, \`ASYNC\_JS\_CACHING\`, \`SYNC\_JS\_NO\_CACHING\`,
\`SYNC\_JS\_CACHING\`. Defaults to \`ASYNC\_JS\_NO\_CACHING\`.

Additional upstream details:

Web Client JavaScript Mode. Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
asynchronously, and it is cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is non-cacheable Bot Defense JavaScript for telemetry collection is requested
synchronously, and it is cacheable.

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
    "ves.io.schema.rules.string.http_path": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.string.http_path": "true"
  }
}
```

- [js_insert_all_pages](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages/): complete subsection reference.

- [js_insert_all_pages_except](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insert_all_pages_except/): complete subsection reference.

- [js_insertion_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/js_insertion_rules/): complete subsection reference.

- [mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/): complete subsection reference.

- [protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/): complete subsection reference.
