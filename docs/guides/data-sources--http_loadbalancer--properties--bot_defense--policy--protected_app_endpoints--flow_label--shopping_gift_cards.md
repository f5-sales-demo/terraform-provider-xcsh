---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 7259, "body_sha256": "sha256:a94c1238aa85d9f2b2748ee5d9e4d221945f46db3a2d5ecb025f6d1eefd53ec3", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:gift_card_make_purchase_with_gift_card", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:gift_card_validation", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_add_to_cart", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_checkout", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_choose_seat", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_enter_drawing_submission", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_make_payment", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_order", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_price_inquiry", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_promo_code_validation", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_purchase_gift_card", "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_update_quantity"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "docs/guides/data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "shopping_gift_cards"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/shopping_gift_cards/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [bot_defense](data-sources--http_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](data-sources--http_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="section"></a>

Type: `"single"`. Computed.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```

## Direct properties

- [gift_card_make_purchase_with_gift_card](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--gift_card_make_purchase_with_gift_card.md): complete subsection reference.

- [gift_card_validation](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--gift_card_validation.md): complete subsection reference.

- [shop_add_to_cart](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_add_to_cart.md): complete subsection reference.

- [shop_checkout](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_checkout.md): complete subsection reference.

- [shop_choose_seat](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_choose_seat.md): complete subsection reference.

- [shop_enter_drawing_submission](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_enter_drawing_submission.md): complete subsection reference.

- [shop_make_payment](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_make_payment.md): complete subsection reference.

- [shop_order](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_order.md): complete subsection reference.

- [shop_price_inquiry](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_price_inquiry.md): complete subsection reference.

- [shop_promo_code_validation](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_promo_code_validation.md): complete subsection reference.

- [shop_purchase_gift_card](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_purchase_gift_card.md): complete subsection reference.

- [shop_update_quantity](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_update_quantity.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--gift_card_make_purchase_with_gift_card.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--gift_card_validation.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_add_to_cart.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_checkout.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_choose_seat.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_enter_drawing_submission.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_make_payment.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_order.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_price_inquiry.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_promo_code_validation.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_purchase_gift_card.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_update_quantity.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](data-sources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
