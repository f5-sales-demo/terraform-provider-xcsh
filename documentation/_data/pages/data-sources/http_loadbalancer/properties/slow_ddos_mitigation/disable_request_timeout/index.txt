---
page_title: "slow_ddos_mitigation.disable_request_timeout"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["duration", "slow ddos mitigation disable request timeout"], "body_bytes": 1040, "body_sha256": "sha256:ffc6e9ad5bddf02a30cc1e0805f35cb3fffb287051eb024b4f9cf635c2a2a972", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:http_loadbalancer:properties:slow_ddos_mitigation:disable_request_timeout", "parent_id": "xcsh-docs:data-sources:http_loadbalancer:properties:slow_ddos_mitigation", "path": "documentation/data-sources/http_loadbalancer/properties/slow_ddos_mitigation/disable_request_timeout/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1233000100013121-3020120210323111-0003301133330201-1100110211111231-2120230111010023-0120130000033030-0200232130311223-0100102033111221", "registry_path": "docs/guides/data-sources--http_loadbalancer--reference--group-026.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["slow_ddos_mitigation", "disable_request_timeout"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/http_loadbalancer/properties/slow_ddos_mitigation/disable_request_timeout/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# slow_ddos_mitigation.disable_request_timeout

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/)
- [slow_ddos_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/http_loadbalancer/properties/slow_ddos_mitigation/)
- slow_ddos_mitigation.disable_request_timeout

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable request timeout.

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
