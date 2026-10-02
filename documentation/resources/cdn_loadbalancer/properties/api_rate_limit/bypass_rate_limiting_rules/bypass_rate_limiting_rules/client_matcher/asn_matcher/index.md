---
page_title: "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher"
subcategory: "Load Balancing"
description: "Match any AS number contained in the list of bgp_asn_sets."
xcsh_docs: {"aliases": ["api rate limit bypass rate limiting rules bypass rate limiting rules client matcher asn matcher"], "body_bytes": 2889, "body_sha256": "sha256:0807166175b1bacda8385bb9cff0123e5e5dc491eefd0c8d6eec7bcd96a5c7ba", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher:asn_matcher:asn_sets"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher:asn_matcher", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher", "path": "documentation/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/asn_matcher/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3032300131130131-0102330220303332-1020322132022203-0113232203130012-1222313231100203-2032203212320100-1320232230113231-1320333230103021", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-004.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher:RequiredObjectAttributes:asn_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher:asn_matcher:asn_sets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules", "client_matcher", "asn_matcher"], "schema_version": 1, "sections": [{"aliases": ["asn sets"], "anchor": "section", "description": "A list of references to bgp_asn_set objects.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:bypass_rate_limiting_rules:bypass_rate_limiting_rules:client_matcher:asn_matcher:asn_sets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "bypass_rate_limiting_rules", "bypass_rate_limiting_rules", "client_matcher", "asn_matcher", "asn_sets"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/asn_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Match any AS number contained in the list of bgp_asn_sets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/)
- api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any AS number contained in the list of bgp\_asn\_sets.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("asn_sets")}
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
asn_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

- [asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/asn_matcher/asn_sets/): complete subsection reference.

## Next pages

- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/asn_matcher/asn_sets/)
- [api_rate_limit.bypass_rate_limiting_rules.bypass_rate_limiting_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/bypass_rate_limiting_rules/bypass_rate_limiting_rules/client_matcher/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
