---
page_title: "routes.simple_route.advanced_options.mirror_policy"
subcategory: "Load Balancing"
description: "routes.simple_route.advanced_options.mirror_policy for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 2469, "body_sha256": "sha256:bfd8c749e8b5d1359047233f32ffd24a9c022c651d2bd514be901c8e1929373c", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:origin_pool", "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy:percent"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options:mirror_policy", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:routes:simple_route:advanced_options", "path": "docs/guides/data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--mirror_policy.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "advanced_options", "mirror_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/routes/simple_route/advanced_options/mirror_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.advanced_options.mirror_policy for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.advanced_options.mirror_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [routes](data-sources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](data-sources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- routes.simple_route.advanced_options.mirror_policy

<a id="section"></a>

Type: `"single"`. Computed.

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
'fire and forget', meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow
origin..

Upstream description:

MirrorPolicy is used for shadowing traffic from one origin pool to another. The approach used is
"fire and forget", meaning it will not wait for the shadow origin pool to respond before returning
the response from the primary origin pool. All normal statistics are collected for the shadow origin
pool making this feature useful for testing and troubleshooting.

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

## Direct properties

- [origin_pool](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--mirror_policy--origin_pool.md): complete subsection reference.

- [percent](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--mirror_policy--percent.md): complete subsection reference.

## Next pages

- [routes.simple_route.advanced_options.mirror_policy.origin_pool](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--mirror_policy--origin_pool.md)
- [routes.simple_route.advanced_options.mirror_policy.percent](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options--mirror_policy--percent.md)
- [routes.simple_route.advanced_options](data-sources--http_loadbalancer--properties--routes--simple_route--advanced_options.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
