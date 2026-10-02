---
page_title: "api_rate_limit.custom_ip_allowed_list"
subcategory: "Load Balancing"
description: "IP Allowed list using existing ip_prefix_set objects."
xcsh_docs: {"aliases": ["api rate limit custom ip allowed list"], "body_bytes": 1882, "body_sha256": "sha256:300a2575163ae94d440c2b38eb5cbbb0766d7712f1624abba5ee35b375a02f31", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit", "path": "documentation/resources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2333012122103001-0111232101323001-2333210301300122-3302011003223222-3002023122303203-1103001221000201-1310203202022230-2010003132313013", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.custom_ip_allowed_list:RequiredObjectAttributes:rate_limiter_allowed_prefixes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "custom_ip_allowed_list"], "schema_version": 1, "sections": [{"aliases": ["rate limiter allowed prefixes"], "anchor": "section", "description": "References to ip_prefix_set objects. Requests from source IP addresses that are covered by one of the allowed IP Prefixes are not subjected to rate limiting.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:custom_ip_allowed_list:rate_limiter_allowed_prefixes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "custom_ip_allowed_list", "rate_limiter_allowed_prefixes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "IP Allowed list using existing ip_prefix_set objects.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.custom_ip_allowed_list

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/)
- api_rate_limit.custom_ip_allowed_list

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

IP Allowed list using existing ip\_prefix\_set objects.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("rate_limiter_allowed_prefixes")}
```

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
custom_ip_allowed_list {
  # Configure direct properties listed below.
}
```

## Direct properties

- [rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/rate_limiter_allowed_prefixes/): complete subsection reference.

## Next pages

- [api_rate_limit.custom_ip_allowed_list.rate_limiter_allowed_prefixes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/custom_ip_allowed_list/rate_limiter_allowed_prefixes/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
