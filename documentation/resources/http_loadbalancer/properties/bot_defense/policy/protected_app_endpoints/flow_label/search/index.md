---
page_title: "bot_defense.policy.protected_app_endpoints.flow_label.search"
subcategory: "Load Balancing"
description: "bot_defense.policy.protected_app_endpoints.flow_label.search for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 4497, "body_sha256": "sha256:f28d7d5654ee614c107d9a3161b8372970c79b5e6826e9d7db444611f700e234", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:flight_search", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:product_search", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:reservation_search", "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search:room_search"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label:search", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:bot_defense:policy:protected_app_endpoints:flow_label", "path": "documentation/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/index.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["bot_defense", "policy", "protected_app_endpoints", "flow_label", "search"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "bot_defense.policy.protected_app_endpoints.flow_label.search for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# bot_defense.policy.protected_app_endpoints.flow_label.search

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [bot_defense](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/)
- [bot_defense.policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/)
- [bot_defense.policy.protected_app_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- bot_defense.policy.protected_app_endpoints.flow_label.search

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Bot Defense Flow Label Search Category. Bot Defense Flow Label Search Category.

Upstream description:

Bot Defense Flow Label Search Category.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("flight_search",
    "product_search"),
  validators.ConflictingObjectAttributes("flight_search",
    "reservation_search"),
  validators.ConflictingObjectAttributes("flight_search",
    "room_search"),
  validators.ConflictingObjectAttributes("product_search",
    "reservation_search"),
  validators.ConflictingObjectAttributes("product_search",
    "room_search"),
  validators.ConflictingObjectAttributes("reservation_search",
    "room_search")}
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
  "x-ves-oneof-field-label_choice": "[\"flight_search\",\"product_search\",\"reservation_search\",\"room_search\"]"
}
```

Terraform syntax:

```terraform
search {
  # Configure direct properties listed below.
}
```

## Direct properties

- [flight_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/flight_search/): complete subsection reference.

- [product_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/product_search/): complete subsection reference.

- [reservation_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/reservation_search/): complete subsection reference.

- [room_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/room_search/): complete subsection reference.

## Next pages

- [bot_defense.policy.protected_app_endpoints.flow_label.search.flight_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/flight_search/)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.product_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/product_search/)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.reservation_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/reservation_search/)
- [bot_defense.policy.protected_app_endpoints.flow_label.search.room_search](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/search/room_search/)
- [bot_defense.policy.protected_app_endpoints.flow_label](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/bot_defense/policy/protected_app_endpoints/flow_label/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
