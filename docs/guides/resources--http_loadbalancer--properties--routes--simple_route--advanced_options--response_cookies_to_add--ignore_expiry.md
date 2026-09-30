---
page_title: "routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1515, "body_sha256": "sha256:7dc00759428b60a353a7c7ff940320e919291ff19c51441de7e76b48bd35c2d6", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_expiry", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_expiry", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_expiry.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "response_cookies_to_add", "ignore_expiry"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/ignore_expiry/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_expiry

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore expiry.

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
ignore_expiry = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
