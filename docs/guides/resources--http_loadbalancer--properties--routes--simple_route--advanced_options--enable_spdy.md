---
page_title: "routes.simple_route.advanced_options.enable_spdy"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.enable_spdy for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1235, "body_sha256": "sha256:5c9a26361a2a939ba80726c9cb96d2778e00705d88af2ff56da2c74bc990c724", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:enable_spdy", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:enable_spdy", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route--advanced_options--enable_spdy.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "enable_spdy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/enable_spdy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.enable_spdy for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.simple_route.advanced_options.enable_spdy

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- routes.simple_route.advanced_options.enable_spdy

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable spdy.

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
enable_spdy = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
