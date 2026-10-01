---
page_title: "cloudfront.protected_endpoints.flow_label.search"
subcategory: ""
description: "cloudfront.protected_endpoints.flow_label.search for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 2895, "body_sha256": "sha256:46f9947de6fd879f6a1dabf8506ff16f115ef8f2656f5dad52d7bf7f6d99fac5", "canonical_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search", "child_ids": ["xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search:flight_search", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search:product_search", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search:reservation_search", "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search:room_search"], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label:search", "parent_id": "xcsh-docs:data-sources:protected_application:properties:cloudfront:protected_endpoints:flow_label", "path": "docs/guides/data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["cloudfront", "protected_endpoints", "flow_label", "search"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/properties/cloudfront/protected_endpoints/flow_label/search/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "cloudfront.protected_endpoints.flow_label.search for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# cloudfront.protected_endpoints.flow_label.search

Breadcrumbs:

- [xcsh_protected_application](../data-sources/protected_application.md)
- [Property reference](data-sources--protected_application--reference.md)
- [cloudfront](data-sources--protected_application--properties--cloudfront.md)
- [cloudfront.protected_endpoints](data-sources--protected_application--properties--cloudfront--protected_endpoints.md)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- cloudfront.protected_endpoints.flow_label.search

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

- [flight_search](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--flight_search.md): complete subsection reference.

- [product_search](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--product_search.md): complete subsection reference.

- [reservation_search](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--reservation_search.md): complete subsection reference.

- [room_search](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--room_search.md): complete subsection reference.

## Next pages

- [cloudfront.protected_endpoints.flow_label.search.flight_search](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--flight_search.md)
- [cloudfront.protected_endpoints.flow_label.search.product_search](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--product_search.md)
- [cloudfront.protected_endpoints.flow_label.search.reservation_search](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--reservation_search.md)
- [cloudfront.protected_endpoints.flow_label.search.room_search](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label--search--room_search.md)
- [cloudfront.protected_endpoints.flow_label](data-sources--protected_application--properties--cloudfront--protected_endpoints--flow_label.md)
- [xcsh_protected_application](../data-sources/protected_application.md)
