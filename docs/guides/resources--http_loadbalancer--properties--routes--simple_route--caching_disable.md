---
page_title: "routes.simple_route.caching_disable"
subcategory: "Load Balancing"
description: "routes.simple_route.caching_disable for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1155, "body_sha256": "sha256:4bf25737581ead591d260e02c17b87b6ac415a4c6cac23d773e9d522be2df6f3", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_disable", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:caching_disable", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route--caching_disable.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "caching_disable"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/caching_disable/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.caching_disable for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.simple_route.caching_disable

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- routes.simple_route.caching_disable

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching disable.

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
caching_disable = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
