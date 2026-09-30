---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 13558, "body_sha256": "sha256:1cee155554c9efab44f46b73a0cfffd644fc6db8b4eafba7dd2492c836ab96cc", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:gift_card_make_purchase_with_gift_card", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:gift_card_validation", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_add_to_cart", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_checkout", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_choose_seat", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_enter_drawing_submission", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_make_payment", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_order", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_price_inquiry", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_promo_code_validation", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_purchase_gift_card", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards:shop_update_quantity"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:shopping_gift_cards", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "docs/guides/resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "shopping_gift_cards"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/shopping_gift_cards/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [bot_defense](resources--http_loadbalancer--properties--bot_defense.md)
- [bot_defense.policy](resources--http_loadbalancer--properties--bot_defense--policy.md)
- [bot_defense.policy.protected_app_endpoints](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Shopping &amp; Gift Cards Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "gift_card_validation"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_add_to_cart"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_order"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("gift_card_make_purchase_with_gift_card",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_add_to_cart"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_order"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("gift_card_validation",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_checkout"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_add_to_cart",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_choose_seat"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_checkout",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_enter_drawing_submission"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_choose_seat",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_make_payment"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_enter_drawing_submission",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_order"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_make_payment",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_price_inquiry"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_order",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_promo_code_validation"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_price_inquiry",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_promo_code_validation",
    "shop_purchase_gift_card"),
  validators.ConflictingObjectAttributes("shop_promo_code_validation",
    "shop_update_quantity"),
  validators.ConflictingObjectAttributes("shop_purchase_gift_card",
    "shop_update_quantity")}
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
  "x-ves-oneof-field-label_choice": "[\"gift_card_make_purchase_with_gift_card\",\"gift_card_validation\",\"shop_add_to_cart\",\"shop_checkout\",\"shop_choose_seat\",\"shop_enter_drawing_submission\",\"shop_make_payment\",\"shop_order\",\"shop_price_inquiry\",\"shop_promo_code_validation\",\"shop_purchase_gift_card\",\"shop_update_quantity\"]"
}
```

Terraform syntax:

```terraform
shopping_gift_cards {
  # Configure direct properties listed below.
}
```

## Direct properties

- [gift_card_make_purchase_with_gift_card](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--gift_card_make_purchase_with_gift_card.md): complete subsection reference.

- [gift_card_validation](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--gift_card_validation.md): complete subsection reference.

- [shop_add_to_cart](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_add_to_cart.md): complete subsection reference.

- [shop_checkout](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_checkout.md): complete subsection reference.

- [shop_choose_seat](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_choose_seat.md): complete subsection reference.

- [shop_enter_drawing_submission](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_enter_drawing_submission.md): complete subsection reference.

- [shop_make_payment](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_make_payment.md): complete subsection reference.

- [shop_order](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_order.md): complete subsection reference.

- [shop_price_inquiry](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_price_inquiry.md): complete subsection reference.

- [shop_promo_code_validation](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_promo_code_validation.md): complete subsection reference.

- [shop_purchase_gift_card](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_purchase_gift_card.md): complete subsection reference.

- [shop_update_quantity](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_update_quantity.md): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--gift_card_make_purchase_with_gift_card.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.gift_card_validation](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--gift_card_validation.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_add_to_cart.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_checkout](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_checkout.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_choose_seat.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_enter_drawing_submission.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_make_payment](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_make_payment.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_order](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_order.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_price_inquiry.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_promo_code_validation.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_purchase_gift_card.md)
- [bot_defense.policy.protected_app_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label--shopping_gift_cards--shop_update_quantity.md)
- [bot_defense.policy.protected_app_endpoints.flow_label](resources--http_loadbalancer--properties--bot_defense--policy--protected_app_endpoints--flow_label.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
