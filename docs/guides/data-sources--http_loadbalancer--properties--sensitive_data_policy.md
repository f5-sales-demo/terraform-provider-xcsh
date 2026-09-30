---
page_title: "sensitive_data_policy"
subcategory: "Load Balancing"
description: "sensitive_data_policy for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1003, "body_sha256": "sha256:b315d2211021d0a79ef850834c5e4cc736ac43bd1e0372a72bec6bf7d5a72d36", "canonical_id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_policy", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_policy:sensitive_data_policy_ref"], "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:sensitive_data_policy", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:reference", "path": "docs/guides/data-sources--http_loadbalancer--properties--sensitive_data_policy.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["sensitive_data_policy"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/sensitive_data_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "sensitive_data_policy for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# sensitive_data_policy

Breadcrumbs:

- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- sensitive_data_policy

<a id="section"></a>

Type: `"single"`. Computed.

Policy configuration for this feature.

Upstream description:

Settings for data type policy.

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

- [sensitive_data_policy_ref](data-sources--http_loadbalancer--properties--sensitive_data_policy--sensitive_data_policy_ref.md): complete subsection reference.

## Next pages

- [sensitive_data_policy.sensitive_data_policy_ref](data-sources--http_loadbalancer--properties--sensitive_data_policy--sensitive_data_policy_ref.md)
- [Property reference](data-sources--http_loadbalancer--reference.md)
- [xcsh_http_loadbalancer](../data-sources/http_loadbalancer.md)
