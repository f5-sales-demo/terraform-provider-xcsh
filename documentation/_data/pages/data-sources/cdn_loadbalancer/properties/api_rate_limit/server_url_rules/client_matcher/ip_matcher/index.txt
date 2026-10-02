---
page_title: "api_rate_limit.server_url_rules.client_matcher.ip_matcher"
subcategory: "Load Balancing"
description: "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true."
xcsh_docs: {"aliases": ["api rate limit server url rules client matcher ip matcher"], "body_bytes": 2705, "body_sha256": "sha256:8598abf8626de86fe952e4a641b326d9f541b4c0b892b72a023bd0dc49ec3896", "capabilities": ["cdn", "security.rate-limiting"], "category": "cdn", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher:prefix_sets"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher", "parent_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher", "path": "documentation/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/ip_matcher/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2232310130330032-0230113113110122-3311201323331201-1123310322203310-3122013132022210-0203031021101102-0303110220332330-1131012111133130", "registry_path": "docs/guides/data-sources--cdn_loadbalancer--reference--group-005.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher", "ip_matcher"], "schema_version": 1, "sections": [{"aliases": ["invert matcher"], "anchor": "schema-api_rate_limit--server_url_rules--client_matcher--ip_matcher--invert_matcher", "description": "Invert the match result.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher", "ip_matcher", "invert_matcher"], "syntax": "attribute", "type": "bool"}, {"aliases": ["prefix sets"], "anchor": "section", "description": "A list of references to ip_prefix_set objects.", "document_id": "xcsh-docs:data-sources:cdn_loadbalancer:properties:api_rate_limit:server_url_rules:client_matcher:ip_matcher:prefix_sets", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["api_rate_limit", "server_url_rules", "client_matcher", "ip_matcher", "prefix_sets"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/ip_matcher/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Match any IP prefix contained in the list of ip_prefix_sets. The result of the match is inverted if invert_matcher is true.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# api_rate_limit.server_url_rules.client_matcher.ip_matcher

Breadcrumbs:

- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [api_rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/)
- [api_rate_limit.server_url_rules](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/)
- [api_rate_limit.server_url_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/)
- api_rate_limit.server_url_rules.client_matcher.ip_matcher

<a id="section"></a>

Type: `"single"`. Computed.

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

Upstream description:

Match any IP prefix contained in the list of ip\_prefix\_sets. The result of the match is inverted
if invert\_matcher is true.

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

<a id="schema-api_rate_limit--server_url_rules--client_matcher--ip_matcher--invert_matcher"></a>

### invert_matcher property

Type: `"bool"`. Computed.

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

- [prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/ip_matcher/prefix_sets/): complete subsection reference.

## Next pages

- [api_rate_limit.server_url_rules.client_matcher.ip_matcher.prefix_sets](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/ip_matcher/prefix_sets/)
- [api_rate_limit.server_url_rules.client_matcher](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/api_rate_limit/server_url_rules/client_matcher/)
- [xcsh_cdn_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/)
