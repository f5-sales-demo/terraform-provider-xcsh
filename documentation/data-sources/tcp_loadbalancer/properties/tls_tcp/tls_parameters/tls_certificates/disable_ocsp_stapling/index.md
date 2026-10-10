---
page_title: "tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling"
subcategory: "Load Balancing"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls tcp tls parameters tls certificates disable ocsp stapling"], "body_bytes": 1374, "body_sha256": "sha256:29f98c7cbcda4db08a9b6ef3c8f22c231b4d16601dbdf91b19bf5da1cc8bcff7", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates:disable_ocsp_stapling", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:properties:tls_tcp:tls_parameters:tls_certificates", "path": "documentation/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/disable_ocsp_stapling/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2211310213312032-2023101020123311-3223220131132300-1013031022211013-2223221012220020-1231010030320333-3300313033101131-1300230323330023", "registry_path": "docs/guides/data-sources--tcp_loadbalancer--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_tcp", "tls_parameters", "tls_certificates", "disable_ocsp_stapling"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/)
- [tls_tcp](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/)
- [tls_tcp.tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/)
- [tls_tcp.tls_parameters.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/properties/tls_tcp/tls_parameters/tls_certificates/)
- tls_tcp.tls_parameters.tls_certificates.disable_ocsp_stapling

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
