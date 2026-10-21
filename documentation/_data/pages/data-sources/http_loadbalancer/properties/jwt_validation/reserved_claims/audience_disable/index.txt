---
page_title: "jwt_validation.reserved_claims.audience_disable"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["jwt validation reserved claims audience disable"], "body_bytes": 1193, "body_sha256": "sha256:421aab3bc479c598b190e09dae8ae70ba6223580b86093a71d599f665498a3db", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:reserved_claims:audience_disable", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:jwt_validation:reserved_claims", "path": "documentation/data-sources/http_loadbalancer/properties/jwt_validation/reserved_claims/audience_disable/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3031023031030233-2100312032300111-2330311112001331-2202321232311020-0311322330211103-1103300201201030-1320213302302223-3232300330013301", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-019.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["jwt_validation", "reserved_claims", "audience_disable"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/jwt_validation/reserved_claims/audience_disable/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# jwt_validation.reserved_claims.audience_disable

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [jwt_validation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/)
- [jwt_validation.reserved_claims](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/jwt_validation/reserved_claims/)
- jwt_validation.reserved_claims.audience_disable

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for audience disable.

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
