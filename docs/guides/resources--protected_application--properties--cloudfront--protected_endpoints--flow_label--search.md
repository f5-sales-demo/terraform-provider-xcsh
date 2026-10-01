---
page_title: "cloudfront.protected_endpoints.flow_label.search"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.search for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 3553, "body_sha256": "sha256:68deee4ccf73894e0ffaf69f2dd7a6bd111af5ad5ec0f8e996544fe152cd1f87", "canonical_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search", "child_ids": ["xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search:flight_search", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search:product_search", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search:reservation_search", "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search:room_search"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search", "parent_id": "xcsh-docs:resources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "docs/guides/resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "search"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/properties/cloudfront/protected_endpoints/flow_label/search/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.search for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.search

Breadcrumbs:

- [xcsh_protected_application](../resources/protected_application.md)
- [Property reference](resources--protected_application--reference.md)
- [cloudfront](resources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](resources--protected_application--properties--cloudfront--protected_endpoints.md)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- cloudfront.protected_endpoints.flow_label.search

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

- [flight_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--flight_search.md): complete subsection reference.

- [product_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--product_search.md): complete subsection reference.

- [reservation_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--reservation_search.md): complete subsection reference.

- [room_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--room_search.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.search.flight_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--flight_search.md)
- [cloudfront.protected_endpoints.flow_label.search.product_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--product_search.md)
- [cloudfront.protected_endpoints.flow_label.search.reservation_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--reservation_search.md)
- [cloudfront.protected_endpoints.flow_label.search.room_search](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--room_search.md)
- [cloudfront.protected_endpoints.flow_label](resources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- [xcsh_protected_application](../resources/protected_application.md)
