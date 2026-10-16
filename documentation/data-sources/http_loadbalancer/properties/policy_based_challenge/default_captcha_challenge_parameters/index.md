---
page_title: "policy_based_challenge.default_captcha_challenge_parameters"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["policy based challenge default captcha challenge parameters"], "body_bytes": 1087, "body_sha256": "sha256:ad2c00fa2fd968c589394a59bee7243450b3ee7eaa9d5f531738f4b44ca11ad4", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge:default_captcha_challenge_parameters", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:policy_based_challenge", "path": "documentation/data-sources/http_loadbalancer/properties/policy_based_challenge/default_captcha_challenge_parameters/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-3033122122132230-3012313221123333-2130013103000000-2130320102210233-2202310122202323-0320010320033221-3010302311012312-1221131320330202", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-022.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["policy_based_challenge", "default_captcha_challenge_parameters"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/policy_based_challenge/default_captcha_challenge_parameters/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# policy_based_challenge.default_captcha_challenge_parameters

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [policy_based_challenge](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/policy_based_challenge/)
- policy_based_challenge.default_captcha_challenge_parameters

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for default captcha challenge parameters.

Additional upstream details:

This can be used for messages where no values are needed.

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

This is an empty object or choice marker. It has no direct properties.
