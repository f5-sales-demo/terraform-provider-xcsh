---
page_title: "rate_limit.rate_limiter.action_block"
subcategory: "Load Balancing"
description: "Action where a user is blocked from making further requests after exceeding rate limit threshold."
xcsh_docs: {"aliases": ["rate limit rate limiter action block"], "body_bytes": 1706, "body_sha256": "sha256:535cef90913ee8a98350a390447c8e087548db18d01403c9d4e0f34574aeac0f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter", "path": "documentation/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1200211123003010-1233100203111302-0210123300133003-2122300031101021-1201313133330132-2231222222013231-2311202021020223-0023311303233311", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-024.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["rate_limit", "rate_limiter", "action_block"], "schema_version": 1, "sections": [{"aliases": ["rate limit rate limiter action block hours"], "anchor": "section", "description": "Input Duration Hours.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:hours", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "action_block", "hours"], "syntax": "attribute", "type": "object"}, {"aliases": ["rate limit rate limiter action block minutes"], "anchor": "section", "description": "Input Duration Minutes.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:minutes", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "action_block", "minutes"], "syntax": "attribute", "type": "object"}, {"aliases": ["rate limit rate limiter action block seconds"], "anchor": "section", "description": "Input Duration Seconds.", "document_id": "xcsh-docs:data-sources:http_loadbalancer:properties:rate_limit:rate_limiter:action_block:seconds", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["rate_limit", "rate_limiter", "action_block", "seconds"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Action where a user is blocked from making further requests after exceeding rate limit threshold.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# rate_limit.rate_limiter.action_block

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [rate_limit](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/)
- [rate_limit.rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/)
- rate_limit.rate_limiter.action_block

<a id="section"></a>

Type: `"single"`. Computed.

Action where a user is blocked from making further requests after exceeding rate limit threshold.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-block_duration_choice": "[\"hours\",\"minutes\",\"seconds\"]"
}
```

## Direct properties

- [hours](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/hours/): complete subsection reference.

- [minutes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/minutes/): complete subsection reference.

- [seconds](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/rate_limit/rate_limiter/action_block/seconds/): complete subsection reference.
