---
page_title: "api_rate_limit.server_url_rules.client_matcher.ip_matcher"
subcategory: "Load Balancing"
description: "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["api rate limit server url rules client matcher ip matcher"], "body_bytes": 2958, "body_sha256": "sha256:d44bab4f8229c5d9593ea537332d15e290304f7ab1356a6c0ed9d3aed9fcf6e6", "capabilities": ["load-balancing", "security.rate-limiting"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher:prefix_sets"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher", "path": "documentation/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/ip_matcher/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1201312112023220-3110102133111300-3200210011301312-1102221323231012-3201330200312020-0221113222213323-0212313312200010-1101310120100232", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-008.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "api_rate_limit.server_url_rules.client_matcher.ip_matcher:RequiredObjectAttributes:prefix_sets", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher:prefix_sets", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher", "ip_matcher"], "schema_version": 1, "sections": [{"aliases": ["api rate limit server url rules client matcher ip matcher invert matcher"], "anchor": "schema-api_rate_limit--server_url_rules--client_matcher--ip_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher", "ip_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["api rate limit server url rules client matcher ip matcher prefix sets"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:resources:http_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher:prefix_sets", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher", "ip_matcher", "prefix_sets"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/ip_matcher/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules.client_matcher.ip_matcher

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/)
- [api_rate_limit.server_url_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("prefix_sets")}
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
ip_matcher {
  # Configure direct properties listed below.
}
```

## Direct properties

<a id="schema-api_rate_limit--server_url_rules--client_matcher--ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Optional.

Invert IP Matcher. Invert the match result.

Upstream description:

Invert the match result.

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

- [prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/ip_matcher/prefix_sets/): complete subsection reference.

## Next pages

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/ip_matcher/prefix_sets/)
- [api_rate_limit.server_url_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/)
- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
