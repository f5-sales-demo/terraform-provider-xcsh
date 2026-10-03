---
page_title: "bot_defense.policy.mobile_sdk_config"
subcategory: "Load Balancing"
description: "Mobile SDK configuration."
xcsh_docs: {"aliases": ["bot defense policy mobile sdk config"], "body_bytes": 1740, "body_sha256": "sha256:5c442c2892ae36bcc58fb21402a0123e6e0168fdf35535b32ede39e4529b42c6", "capabilities": ["cdn", "security.bot-defense"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy", "path": "documentation/data-sources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2212311100013330-3302031011003100-3121130332112301-2222033303133101-0101302002131221-3213321232023301-3022211202010011-1010112330103303", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-007.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "mobile_sdk_config"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy mobile sdk config mobile identifier"], "anchor": "section", "description": "Mobile traffic identifier type.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:bot_defense:policy:mobile_sdk_config:mobile_identifier", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["bot_defense", "policy", "mobile_sdk_config", "mobile_identifier"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Mobile SDK configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.mobile_sdk_config

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/)
- bot_defense.policy.mobile_sdk_config

<a id="section"></a>

Type: `"single"`. Computed.

Mobile SDK Configuration. Mobile SDK configuration.

Upstream description:

Mobile SDK configuration.

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

## Direct properties

- [mobile_identifier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/): complete subsection reference.

## Next pages

- [bot_defense.policy.mobile_sdk_config.mobile_identifier](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/mobile_sdk_config/mobile_identifier/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/bot_defense/policy/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
