---
page_title: "routes.route_state_disabled"
subcategory: "Load Balancing"
description: "routes.route_state_disabled for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1000, "body_sha256": "sha256:38d0b82540a4fbe6ee30f0fe7bdaac9a6b0273adb10328a7768f4b331bb316af", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_disabled", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:routes:route_state_disabled", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:routes", "path": "docs/guides/resources--http_loadbalancer--properties--routes--route_state_disabled.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["routes", "route_state_disabled"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/routes/route_state_disabled/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "routes.route_state_disabled for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# routes.route_state_disabled

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [routes](resources--http_loadbalancer--properties--routes.md)
- routes.route_state_disabled

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
route_state_disabled = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [routes](resources--http_loadbalancer--properties--routes.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
