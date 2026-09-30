---
page_title: "routes.simple_route.query_params.remove_all_params"
subcategory: "Load Balancing"
description: "routes.simple_route.query_params.remove_all_params for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1235, "body_sha256": "sha256:1ee03bfc24068d17959aba925529e3d8cc58aa1c727287384202bb2d24600ca1", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params:remove_all_params", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params:remove_all_params", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:simple_route:query_params", "path": "docs/guides/resources--http_loadbalancer--properties--routes--simple_route--query_params--remove_all_params.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "simple_route", "query_params", "remove_all_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/simple_route/query_params/remove_all_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.simple_route.query_params.remove_all_params for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.simple_route.query_params.remove_all_params

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.simple_route](resources--http_loadbalancer--properties--routes--simple_route.md)
- [routes.simple_route.query_params](resources--http_loadbalancer--properties--routes--simple_route--query_params.md)
- routes.simple_route.query_params.remove_all_params

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for remove all params.

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
remove_all_params = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.simple_route.query_params](resources--http_loadbalancer--properties--routes--simple_route--query_params.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
