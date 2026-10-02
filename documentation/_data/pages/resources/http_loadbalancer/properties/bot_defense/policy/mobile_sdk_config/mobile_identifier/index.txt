---
page_title: "bot_defense.policy.mobile_sdk_config.mobile_identifier"
subcategory: "Load Balancing"
description: "Mobile traffic identifier type."
xcsh_docs: {"aliases": ["bot defense policy mobile sdk config mobile identifier"], "body_bytes": 2126, "body_sha256": "sha256:6a2a3aa409db3733cfa8d20f897ce7971f999e0234f451a503b110fa8bc9f875", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2020001123330232-1331033320103030-2100212231130230-2211020212323023-2231323333310011-1130221220222202-1100231030232103-0020022130322213", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-011.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "mobile_sdk_config", "mobile_identifier"], "schema_version": 1, "sections": [{"aliases": ["headers"], "anchor": "section", "description": "Headers that can be used to identify mobile traffic.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:item", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--mobile_sdk_config--mobile_identifier--headers--name", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers", "type": "requires"}], "schema_path": ["bot_defense", "policy", "mobile_sdk_config", "mobile_identifier", "headers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Mobile traffic identifier type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.mobile_sdk_config.mobile_identifier

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

Upstream description:

Mobile traffic identifier type.

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
mobile_identifier {
  # Configure direct properties listed below.
}
```

## Direct properties

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/headers/): complete subsection reference.

## Next pages

- [bot_defense.policy.mobile_sdk_config.mobile_identifier.headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/headers/)
- [bot_defense.policy.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
