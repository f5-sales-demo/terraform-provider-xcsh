---
page_title: "caching_policy.custom_cache_rule"
subcategory: "Load Balancing"
description: "caching_policy.custom_cache_rule for xcsh_http_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1223, "body_sha256": "sha256:17993519904142c04e8bb8803d9edf776be387218b1c10d5722090e0ab2e83f5", "canonical_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule:cdn_cache_rules"], "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy:custom_cache_rule", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:caching_policy", "path": "docs/guides/resources--http_loadbalancer--properties--caching_policy--custom_cache_rule.md", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["caching_policy", "custom_cache_rule"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/caching_policy/custom_cache_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "caching_policy.custom_cache_rule for xcsh_http_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# caching_policy.custom_cache_rule

Breadcrumbs:

- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
- [Property reference](resources--http_loadbalancer--reference.md)
- [caching_policy](resources--http_loadbalancer--properties--caching_policy.md)
- caching_policy.custom_cache_rule

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Custom Cache Rules. Caching policies for CDN.

Upstream description:

Caching policies for CDN.

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
custom_cache_rule {
  # Configure direct properties listed below.
}
```

## Direct properties

- [cdn_cache_rules](resources--http_loadbalancer--properties--caching_policy--custom_cache_rule--cdn_cache_rules.md): complete subsection reference.

## Next pages

- [caching_policy.custom_cache_rule.cdn_cache_rules](resources--http_loadbalancer--properties--caching_policy--custom_cache_rule--cdn_cache_rules.md)
- [caching_policy](resources--http_loadbalancer--properties--caching_policy.md)
- [xcsh_http_loadbalancer](../resources/http_loadbalancer.md)
