---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.search"
subcategory: "Load Balancing"
description: "Bot Defense Flow Label Search Category."
xcsh_docs: {"aliases": ["bot defense policy protected app endpoints flow label search"], "body_bytes": 3941, "body_sha256": "sha256:24093e5b63e76946b3f12a8462930e53ceff66d9852df856fe1ad6329f0eb897", "capabilities": ["load-balancing", "security.bot-defense"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:flight_search", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:product_search", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:reservation_search", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:room_search"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "documentation/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3112133023312102-1032003322110323-3131220313222003-3302313012123211-2102303331023320-2003223332333322-2123021130111101-2333021021212123", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-012.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "search"], "schema_version": 1, "sections": [{"aliases": ["bot defense policy protected app endpoints flow label search flight search"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:flight_search", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "search", "flight_search"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy protected app endpoints flow label search product search"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:product_search", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "search", "product_search"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy protected app endpoints flow label search reservation search"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:reservation_search", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "search", "reservation_search"], "syntax": "attribute", "type": "object"}, {"aliases": ["bot defense policy protected app endpoints flow label search room search"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:room_search", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "search", "room_search"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Bot Defense Flow Label Search Category.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.search

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Upstream description:

Bot Defense Flow Label Search Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

## Direct properties

- [flight_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/flight_search/): complete subsection reference.

- [product_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/product_search/): complete subsection reference.

- [reservation_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/reservation_search/): complete subsection reference.

- [room_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/room_search/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/flight_search/)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.product_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/product_search/)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/reservation_search/)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.room_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/room_search/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
