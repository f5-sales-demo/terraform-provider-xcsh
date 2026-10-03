---
page_title: "api_rate_limit.server_url_rules.client_matcher.asn_matcher"
subcategory: "Load Balancing"
description: "Match any AS number contained in the list of bgp_asn_sets."
xcsh_docs: {"aliases": ["api rate limit server url rules client matcher asn matcher"], "body_bytes": 2298, "body_sha256": "sha256:e777213dab414f79bd43430ffc9d21f1389de9e86869019c438365e1a6cf10ba", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher:asn_sets"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher", "parent_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher", "path": "documentation/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/asn_matcher/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3123212302233332-0002110301331331-1012311122210102-0301230130022331-3210120210310121-3333320313133003-1131312131002201-1023000203303332", "registry_path": "docs/guides/resources--cdn_loadbalancer--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher.asn_matcher:RequiredObjectAttributes:asn_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher:asn_sets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher", "asn_matcher"], "schema_version": 1, "sections": [{"aliases": ["api rate limit server url rules client matcher asn matcher asn sets"], "anchor": "section", "description": "A list of references to bgp_asn_set objects.", "document_id": "xcsh-docs:resources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:asn_matcher:asn_sets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher", "asn_matcher", "asn_sets"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/asn_matcher/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Match any AS number contained in the list of bgp_asn_sets.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules.client_matcher.asn_matcher

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/)
- [api_rate_limit.server_url_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/)
- api_rate_limit.server_url_rules.client_matcher.asn_matcher

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

- [asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/asn_matcher/asn_sets/): complete subsection reference.

## Next pages

- [api_rate_limit.server_url_rules.client_matcher.asn_matcher.asn_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/asn_matcher/asn_sets/)
- [api_rate_limit.server_url_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/)
