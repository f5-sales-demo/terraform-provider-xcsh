---
page_title: "bot_defense.policy.mobile_sdk_config.mobile_identifier"
subcategory: "Load Balancing"
description: "Mobile traffic identifier type."
xcsh_docs: {"aliases": ["bot defense policy mobile sdk config mobile identifier"], "body_bytes": 1532, "body_sha256": "sha256:21c483df007da1d1246222b4b92cfff212cdcf038e84aed8fd04de83e606d184", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "path": "documentation/resources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3202003010100121-0220322013111031-3221310112201330-1201313332330031-1010323111303122-3113200303311102-0101302332313101-0001011323222120", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "mobile_sdk_config", "mobile_identifier"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy mobile sdk config mobile identifier headers"], "anchor": "section", "description": "Headers that can be used to identify mobile traffic.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_not_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_not_present,check_present", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:check_present", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_not_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:item", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:ConflictingListObjectAttributes:check_present,item", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers:item", "type": "conflicts"}, {"anchor": "schema-bot_defense--policy--mobile_sdk_config--mobile_identifier--headers--name", "enforcement": "provider-schema", "group": "bot_defense.policy.mobile_sdk_config.mobile_identifier.headers:RequiredListObjectAttributes:name", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier:headers", "type": "requires"}], "schema_path": ["bot_defense", "policy", "mobile_sdk_config", "mobile_identifier", "headers"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Mobile traffic identifier type.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.mobile_sdk_config.mobile_identifier

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.mobile_sdk_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/)
- bot_defense.policy.mobile_sdk_config.mobile_identifier

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Mobile Traffic Identifier. Mobile traffic identifier type.

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

- [headers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/headers/): complete subsection reference.
