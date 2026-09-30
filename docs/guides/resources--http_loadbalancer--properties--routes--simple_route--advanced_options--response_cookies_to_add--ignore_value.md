---
page_title: "routes.simple_route.advanced_options.response_cookies_to_add.ignore_value"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.response_cookies_to_add.ignore_value for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1511, "body_sha256": "sha256:cea6d2a0a7e4f7f140d4fa4fb494ed2e3014cc74584bf6f83644a1fd434c08cb", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_value", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add:ignore_value", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:response_cookies_to_add", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add--ignore_value.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "response_cookies_to_add", "ignore_value"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/response_cookies_to_add/ignore_value/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.response_cookies_to_add.ignore_value for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.simple_route.advanced_options.response_cookies_to_add.ignore_value

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md)
- routes.simple_route.advanced_options.response_cookies_to_add.ignore_value

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for ignore value.

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
ignore_value = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.simple_route.advanced_options.response_cookies_to_add](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--response_cookies_to_add.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
