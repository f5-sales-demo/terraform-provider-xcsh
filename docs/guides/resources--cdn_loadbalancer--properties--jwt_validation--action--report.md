---
page_title: "jwt_validation.action.report"
subcategory: "Load Balancing"
description: "jwt_validation.action.report for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1122, "body_sha256": "sha256:bc84c6aad5743562fbfc88b03a0a20388b1f2bbad9c98a5f5b53a70966eb22a2", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action:report", "child_ids": [], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action:report", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:jwt_validation:action", "path": "docs/guides/resources--cdn_loadbalancer--properties--jwt_validation--action--report.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "action", "report"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/jwt_validation/action/report/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.action.report for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.action.report

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [jwt_validation](resources--cdn_loadbalancer--properties--jwt_validation.md)
- [jwt_validation.action](resources--cdn_loadbalancer--properties--jwt_validation--action.md)
- jwt_validation.action.report

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
report = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [jwt_validation.action](resources--cdn_loadbalancer--properties--jwt_validation--action.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
