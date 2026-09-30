---
page_title: "api_rate_limit.bypass_rate_limiting_rules"
subcategory: "Load Balancing"
description: "api_rate_limit.bypass_rate_limiting_rules for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1453, "body_sha256": "sha256:c7b2e1e286b8e0718b1acbe2ece92a0a6efdec276360ef162707fe29b9ea6b0c", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit", "path": "docs/guides/resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "api_rate_limit.bypass_rate_limiting_rules for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# api_rate_limit.bypass_rate_limiting_rules

Breadcrumbs:

- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
- [Property reference](resources--cdn_loadbalancer--reference.md)
- [api_rate_limit](resources--cdn_loadbalancer--properties--api_rate_limit.md)
- api_rate_limit.bypass_rate_limiting_rules

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

Upstream description:

This category defines rules per URL or API group. If request matches any of these rules, skip Rate
Limiting.

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
bypass_rate_limiting_rules {
  # Configure direct properties listed below.
}
```

## Direct properties

- [bypass_rate_limiting_rules](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules.md): complete subsection reference.

## Next pages

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](resources--cdn_loadbalancer--properties--api_rate_limit--bypass_rate_limiting_rules--bypass_rate_limiting_rules.md)
- [api_rate_limit](resources--cdn_loadbalancer--properties--api_rate_limit.md)
- [xcsh_cdn_loadbalancer](../resources/cdn_loadbalancer.md)
