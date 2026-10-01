---
page_title: "routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1596, "body_sha256": "sha256:33c41b813ccb589c1ac15b514fdaef2883b11cd15c78f3905b79cf186dcaf4a1", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:samesite_strict", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:samesite_strict", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--samesite_strict.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "response_cookies_to_add", "samesite_strict"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/samesite_strict/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md)
- routes.simple_route.advanced_options.response_cookies_to_add.samesite_strict

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Upstream description:

This can be used for messages where no values are needed.

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
samesite_strict = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
