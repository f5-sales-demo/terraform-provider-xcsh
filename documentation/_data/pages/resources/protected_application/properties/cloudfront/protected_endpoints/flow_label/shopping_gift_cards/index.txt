---
page_title: "cloudfront.protected_endpoints.flow_label.shopping_gift_cards"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.shopping_gift_cards for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 14625, "body_sha256": "sha256:9aeb3faa1293ef5b48e023960fdbe556247f0c9be44e27625179831d2b7b65ae", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:gift_card_make_purchase_with_gift_card", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:gift_card_validation", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:shop_add_to_cart", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:shop_checkout", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:shop_choose_seat", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:shop_enter_drawing_submission", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:shop_make_payment", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:shop_order", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:shop_price_inquiry", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:shop_promo_code_validation", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:shop_purchase_gift_card", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards:shop_update_quantity"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:shopping_gift_cards", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "documentation/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/index.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "shopping_gift_cards"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.shopping_gift_cards for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.shopping_gift_cards

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [cloudfront](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/)
- [cloudfront.protected_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- cloudfront.protected_endpoints.flow_label.shopping_gift_cards

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

- [gift_card_make_purchase_with_gift_card](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/gift_card_make_purchase_with_gift_card/): complete subsection reference.

- [gift_card_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/gift_card_validation/): complete subsection reference.

- [shop_add_to_cart](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_add_to_cart/): complete subsection reference.

- [shop_checkout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_checkout/): complete subsection reference.

- [shop_choose_seat](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_choose_seat/): complete subsection reference.

- [shop_enter_drawing_submission](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_enter_drawing_submission/): complete subsection reference.

- [shop_make_payment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_make_payment/): complete subsection reference.

- [shop_order](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_order/): complete subsection reference.

- [shop_price_inquiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_price_inquiry/): complete subsection reference.

- [shop_promo_code_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_promo_code_validation/): complete subsection reference.

- [shop_purchase_gift_card](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_purchase_gift_card/): complete subsection reference.

- [shop_update_quantity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_update_quantity/): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_make_purchase_with_gift_card](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/gift_card_make_purchase_with_gift_card/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.gift_card_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/gift_card_validation/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_add_to_cart](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_add_to_cart/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_checkout](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_checkout/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_choose_seat](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_choose_seat/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_enter_drawing_submission](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_enter_drawing_submission/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_make_payment](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_make_payment/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_order](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_order/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_price_inquiry](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_price_inquiry/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_promo_code_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_promo_code_validation/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_purchase_gift_card](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_purchase_gift_card/)
- [cloudfront.protected_endpoints.flow_label.shopping_gift_cards.shop_update_quantity](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/shopping_gift_cards/shop_update_quantity/)
- [cloudfront.protected_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
