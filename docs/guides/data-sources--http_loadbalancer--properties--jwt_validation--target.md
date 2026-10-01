---
page_title: "jwt_validation.target"
subcategory: "Load Balancing"
description: "jwt_validation.target for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1709, "body_sha256": "sha256:95a20e0e8980ec0a221d056e514118c8b858f058167af1e11a6946ebf2d9b905", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:all_endpoint", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:api_groups", "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target:base_paths"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:target", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation", "path": "docs/guides/data-sources--http_loadbalancer--properties--jwt_validation--target.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["jwt_validation", "target"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/target/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "jwt_validation.target for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.target

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [jwt_validation](data-sources--http_loadbalancer--properties--jwt_validation.md)
- jwt_validation.target

<a id="section"></a>

Type: `"single"`. Computed.

Define endpoints for which JWT token validation will be performed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-target": "[\"all_endpoint\",\"api_groups\",\"base_paths\"]"
}
```

## Direct properties

- [all_endpoint](data-sources--http_loadbalancer--properties--jwt_validation--target--all_endpoint.md): complete subsection reference.

- [api_groups](data-sources--http_loadbalancer--properties--jwt_validation--target--api_groups.md): complete subsection reference.

- [base_paths](data-sources--http_loadbalancer--properties--jwt_validation--target--base_paths.md): complete subsection reference.

## Next pages

- [jwt_validation.target.all_endpoint](data-sources--http_loadbalancer--properties--jwt_validation--target--all_endpoint.md)
- [jwt_validation.target.api_groups](data-sources--http_loadbalancer--properties--jwt_validation--target--api_groups.md)
- [jwt_validation.target.base_paths](data-sources--http_loadbalancer--properties--jwt_validation--target--base_paths.md)
- [jwt_validation](data-sources--http_loadbalancer--properties--jwt_validation.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
