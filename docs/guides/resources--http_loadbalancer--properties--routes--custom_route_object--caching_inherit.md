---
page_title: "routes.custom_route_object.caching_inherit"
subcategory: "Load Balancing"
description: "routes.custom_route_object.caching_inherit for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1098, "body_sha256": "sha256:a891f0d4c13054207ac4800341ab139adaed201298f47253fcf7a1a319064ef1", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_inherit", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object:caching_inherit", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:custom_route_object", "path": "docs/guides/resources--http_loadbalancer--properties--routes--custom_route_object--caching_inherit.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "custom_route_object", "caching_inherit"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/custom_route_object/caching_inherit/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.custom_route_object.caching_inherit for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# routes.custom_route_object.caching_inherit

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.custom_route_object](resources--http_loadbalancer--properties--routes--custom_route_object.md)
- routes.custom_route_object.caching_inherit

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for caching inherit.

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
caching_inherit = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.custom_route_object](resources--http_loadbalancer--properties--routes--custom_route_object.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
