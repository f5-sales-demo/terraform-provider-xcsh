---
page_title: "bot_defense.policy.protected_app_endpoints.mitigation"
subcategory: "Load Balancing"
description: "Modify Bot Defense behavior for a matching request."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints mitigation"], "body_bytes": 3266, "body_sha256": "sha256:51119610d16d4aa6ea5397478a0a8478bb58af7e3c45086b323ffd835f00aa5c", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:block", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3301120112032220-1332132001233133-2231302202321031-2323002231200320-3011221220112331-1233011023211002-2003000312300020-3021303022120101", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-009.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation:ConflictingObjectAttributes:block,flag", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation:ConflictingObjectAttributes:block,redirect", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:block", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation:ConflictingObjectAttributes:block,flag", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation:ConflictingObjectAttributes:flag,redirect", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation:ConflictingObjectAttributes:block,redirect", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation:ConflictingObjectAttributes:flag,redirect", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation"], "schema_version": 1, "sections": [{"aliases": ["block"], "anchor": "section", "description": "Block request and respond with custom content.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:block", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "block"], "syntax": "block", "type": "object"}, {"aliases": ["flag"], "anchor": "section", "description": "Flag mitigation action.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.flag:ConflictingObjectAttributes:append_headers,no_headers", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:append_headers", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.flag:ConflictingObjectAttributes:append_headers,no_headers", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:flag:no_headers", "type": "conflicts"}], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "flag"], "syntax": "block", "type": "object"}, {"aliases": ["redirect"], "anchor": "section", "description": "Redirect request to a custom URI.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-bot_defense--policy--protected_app_endpoints--mitigation--redirect--uri", "enforcement": "provider-schema", "group": "bot_defense.policy.protected_app_endpoints.mitigation.redirect:RequiredObjectAttributes:uri", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:mitigation:redirect", "type": "requires"}], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "mitigation", "redirect"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Modify Bot Defense behavior for a matching request.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.mitigation

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- bot_defense.policy.protected_app_endpoints.mitigation

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Modify Bot Defense behavior for a matching request.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("block",
    "flag"),
  validators.ConflictingObjectAttributes("block",
    "redirect"),
  validators.ConflictingObjectAttributes("flag",
    "redirect")}
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
  "x-ves-oneof-field-action_type": "[\"block\",\"flag\",\"redirect\"]"
}
```

Terraform syntax:

```terraform
mitigation {
  # Configure direct properties listed below.
}
```

## Direct properties

- [block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/block/): complete subsection reference.

- [flag](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/): complete subsection reference.

- [redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/redirect/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.mitigation.block](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/block/)
- [bot_defense.policy.protected_app_endpoints.mitigation.flag](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/flag/)
- [bot_defense.policy.protected_app_endpoints.mitigation.redirect](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/mitigation/redirect/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
