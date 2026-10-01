---
page_title: "jwt_validation.action.block"
subcategory: "Load Balancing"
description: "jwt_validation.action.block for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1127, "body_sha256": "sha256:7d0c392283fcec0491168115a34bc8b9f174c13d6aa384df4979f3ab55129154", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action:block", "child_ids": [], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action:block", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:jwt_validation:action", "path": "docs/guides/resources--http_loadbalancer--properties--jwt_validation--action--block.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "action", "block"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/jwt_validation/action/block/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.action.block for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.action.block

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [jwt_validation](resources--http_loadbalancer--properties--jwt_validation.md)
- [jwt_validation.action](resources--http_loadbalancer--properties--jwt_validation--action.md)
- jwt_validation.action.block

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
block = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [jwt_validation.action](resources--http_loadbalancer--properties--jwt_validation--action.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
