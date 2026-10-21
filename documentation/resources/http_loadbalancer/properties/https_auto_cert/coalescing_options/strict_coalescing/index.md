---
page_title: "https_auto_cert.coalescing_options.strict_coalescing"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["https auto cert coalescing options strict coalescing"], "body_bytes": 1262, "body_sha256": "sha256:6a23579dfbaec588d0d33654ab58eab50da81e746b2731673e89d680f2e72ba9", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:http_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:coalescing_options:strict_coalescing", "parent_id": "xcsh-docs:resources:http_loadbalancer:properties:https_auto_cert:coalescing_options", "path": "documentation/resources/http_loadbalancer/properties/https_auto_cert/coalescing_options/strict_coalescing/index.md", "product": "distributed-cloud", "provider_name": "http_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1322301231300213-0203330120002322-0333333120221003-0200101211100322-2222032130030022-1030331111032330-2001323330202211-0011220210101033", "registry_path": "docs/guides/resources--http_loadbalancer--reference--group-020.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_auto_cert", "coalescing_options", "strict_coalescing"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/http_loadbalancer/properties/https_auto_cert/coalescing_options/strict_coalescing/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["http_loadbalancerCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_auto_cert.coalescing_options.strict_coalescing

Breadcrumbs:

- [xcsh_http_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/)
- [https_auto_cert](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https_auto_cert/)
- [https_auto_cert.coalescing_options](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/http_loadbalancer/properties/https_auto_cert/coalescing_options/)
- https_auto_cert.coalescing_options.strict_coalescing

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for strict coalescing.

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

Terraform syntax:

```terraform
strict_coalescing = {}
```

This is an empty object or choice marker. It has no direct properties.
