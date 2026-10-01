---
page_title: "routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1590, "body_sha256": "sha256:9e2e9ce7e6d0a108c2e65caba4d9670604632e1698800a5cc534e6c27cfec0a9", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:all_load_balancer_domains", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy:all_load_balancer_domains", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:advanced_options:csrf_policy", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy--all_load_balancer_domains.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "csrf_policy", "all_load_balancer_domains"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/advanced_options/csrf_policy/all_load_balancer_domains/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](resources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy.md)
- routes.simple_route.advanced_options.csrf_policy.all_load_balancer_domains

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for all load balancer domains.

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
all_load_balancer_domains = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.simple_route.advanced_options.csrf_policy](resources--http_loadbalancer--properties--routes--simple_route--advanced_options--csrf_policy.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
