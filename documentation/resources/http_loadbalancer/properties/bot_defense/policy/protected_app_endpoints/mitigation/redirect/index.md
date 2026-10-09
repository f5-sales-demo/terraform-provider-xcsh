---
page_title: "bot_defense.policy.protected_app_endpoints.mitigation.redirect"
subcategory: "Load Balancing"
description: "Redirect request to a custom URI."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints mitigation redirect"], "body_bytes": 2625, "body_sha256": "sha256:aa1580b19faa0248709e6227abf1015e908f5d30fbb3f1e2e5ec51a22891b062", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/redirect/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3202002210033013-2200001201003233-3212102232033313-2123032120111221-1100023002110111-3301331223001203-2223102221022213-0232013123220211", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-011.md", "relationships": [{"anchor": "schema-bot_defense--policy--protected_app_endpoints--mitigation--redirect--uri", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.redirect:RequiredObjectAttributes:uri", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "redirect"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy protected app endpoints mitigation redirect uri"], "anchor": "schema-bot_defense--policy--protected_app_endpoints--mitigation--redirect--uri", "description": "URI location for redirect may be relative or absolute.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "redirect", "uri"], "syntax": "attribute", "type": "string"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/redirect/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Redirect request to a custom URI.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.mitigation.redirect

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/)
- bot_defense.policy.protected_app_endpoints.mitigation.redirect

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Redirect bot mitigation. Redirect request to a custom URI.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.RequiredObjectAttributes("uri")}
```

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

Terraform syntax:

```terraform
redirect {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-bot_defense--policy--protected_app_endpoints--mitigation--redirect--uri"></a>

### uri property

Type: `"string"`. Optional.

URI location for redirect may be relative or absolute.

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
      "validatedAt": "2026-10-08T03:45:36+00:00"
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
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.url_or_uri_ref": "true"
  }
}
```
