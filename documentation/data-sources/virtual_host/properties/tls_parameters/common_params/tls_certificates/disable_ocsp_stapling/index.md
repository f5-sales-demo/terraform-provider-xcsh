---
page_title: "tls_parameters.common_params.tls_certificates.disable_ocsp_stapling"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls parameters common params tls certificates disable ocsp stapling"], "body_bytes": 1400, "body_sha256": "sha256:e0b3e0cde757626bb0f422cae4941bdd1620ac717d06f50812724fcd425329b0", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates:disable_ocsp_stapling", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:tls_certificates", "path": "documentation/data-sources/virtual_host/properties/tls_parameters/common_params/tls_certificates/disable_ocsp_stapling/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-1210223002300101-1110011013032331-1312303120101220-2322030232030110-3011330011000030-3022000300311013-3330122120300020-0310201103231212", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "disable_ocsp_stapling"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/tls_parameters/common_params/tls_certificates/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/)
- [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/tls_certificates/)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="section"></a>

Type: `["object", {}]`. Computed.

Configuration parameter for disable ocsp stapling.

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
