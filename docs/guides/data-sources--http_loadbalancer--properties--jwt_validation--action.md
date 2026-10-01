---
page_title: "jwt_validation.action"
subcategory: "Load Balancing"
description: "jwt_validation.action for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1333, "body_sha256": "sha256:5c1899d3175d27e7cec7cd260056393341982c54c0b41bbecb0b86a0ba894b1a", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:action", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:action:block", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:action:report"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:action", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation", "path": "docs/guides/data-sources--http_loadbalancer--properties--jwt_validation--action.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.action for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.action

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [jwt_validation](data-sources--http_loadbalancer--properties--jwt_validation.md)
- jwt_validation.action

<a id="section"></a>

Type: `"single"`. Computed.

Action

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-action_choice": "[\"block\",\"report\"]"
}
```

## Direct properties

- [block](data-sources--http_loadbalancer--properties--jwt_validation--action--block.md): complete subsection reference.

- [report](data-sources--http_loadbalancer--properties--jwt_validation--action--report.md): complete subsection reference.

## Next pages

- [jwt_validation.action.block](data-sources--http_loadbalancer--properties--jwt_validation--action--block.md)
- [jwt_validation.action.report](data-sources--http_loadbalancer--properties--jwt_validation--action--report.md)
- [jwt_validation](data-sources--http_loadbalancer--properties--jwt_validation.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
