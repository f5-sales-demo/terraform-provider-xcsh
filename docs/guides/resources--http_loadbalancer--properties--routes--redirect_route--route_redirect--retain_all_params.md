---
page_title: "routes.redirect_route.route_redirect.retain_all_params"
subcategory: "Load Balancing"
description: "routes.redirect_route.route_redirect.retain_all_params for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1362, "body_sha256": "sha256:dcc9011ad5b8d86a159e55dfececaf7181592c3a66ee4b8b628c83f055aa1200", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect:retain_all_params", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect:retain_all_params", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:redirect_route:route_redirect", "path": "docs/guides/resources--http_loadbalancer--properties--routes--redirect_route--route_redirect--retain_all_params.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "redirect_route", "route_redirect", "retain_all_params"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/redirect_route/route_redirect/retain_all_params/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.redirect_route.route_redirect.retain_all_params for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.redirect_route.route_redirect.retain_all_params

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- [routes.redirect_route](resources--http_loadbalancer--properties--routes--redirect_route.md)
- [routes.redirect_route.route_redirect](resources--http_loadbalancer--properties--routes--redirect_route--route_redirect.md)
- routes.redirect_route.route_redirect.retain_all_params

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for retain all params.

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
retain_all_params = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes.redirect_route.route_redirect](resources--http_loadbalancer--properties--routes--redirect_route--route_redirect.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
