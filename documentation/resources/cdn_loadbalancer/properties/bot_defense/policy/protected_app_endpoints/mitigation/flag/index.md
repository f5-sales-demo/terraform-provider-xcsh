---
page_title: "bot_defense.policy.protected_app_endpoints.mitigation.flag"
subcategory: "Load Balancing"
description: "Flag mitigation action."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints mitigation flag"], "body_bytes": 3106, "body_sha256": "sha256:bcb036b1c609a3f7d066f18410e27180492f16d53ab83b3aabfef7162d63a6e2", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:no_headers"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3000200131002020-3221332213012332-2210213230212022-3133311113302123-3310211110332232-2230130301130221-0020030110210203-3233023122232110", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.flag:ConflictingObjectAttributes:append_headers,no_headers", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.flag:ConflictingObjectAttributes:append_headers,no_headers", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:no_headers", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "flag"], "schema_version": 1, "sections": [{"aliases": ["append headers"], "anchor": "section", "description": "Append flag mitigation headers to forwarded request.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--protected_app_endpoints--mitigation--flag--append_headers--auto_type_header_name", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers:RequiredObjectAttributes:auto_type_header_name,inference_header_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "type": "requires"}, {"anchor": "schema-bot_defense--policy--protected_app_endpoints--mitigation--flag--append_headers--inference_header_name", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers:RequiredObjectAttributes:auto_type_header_name,inference_header_name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "type": "requires"}], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "flag", "append_headers"], "syntax": "block", "type": "object"}, {"aliases": ["no headers"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:no_headers", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "flag", "no_headers"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Flag mitigation action.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.mitigation.flag

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/)
- bot_defense.policy.protected_app_endpoints.mitigation.flag

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Select Flag Bot Mitigation Action. Flag mitigation action.

Upstream description:

Flag mitigation action.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("append_headers",
    "no_headers")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-send_headers_choice": "[\"append_headers\",\"no_headers\"]"
}
```

Terraform syntax:

```terraform
flag {
  # Configure direct properties listed below.
}
```

## Direct properties

- [append_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/append_headers/): complete subsection reference.

- [no_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/no_headers/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.mitigation.flag.append_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/append_headers/)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag.no_headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/no_headers/)
- [bot_defense.policy.protected_app_endpoints.mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
