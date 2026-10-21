---
page_title: "tls_parameters.cert_params.volterra_trusted_ca"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls parameters cert params volterra trusted ca"], "body_bytes": 1186, "body_sha256": "sha256:b723a74a8030250fd4b66a64c93514ac60964900b3f6941b04f7aae92c17e205", "capabilities": ["load-balancing.tls"], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params:volterra_trusted_ca", "parent_id": "xcsh-docs:resources:cluster:properties:tls_parameters:cert_params", "path": "documentation/resources/cluster/properties/tls_parameters/cert_params/volterra_trusted_ca/index.md", "product": "distributed-cloud", "provider_name": "cluster", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-0232001101331213-3011133133010213-0313330030310323-1311002133300112-1300011211333120-3323002231113130-0300100332301033-0311100310121120", "registry_path": "docs/guides/resources--cluster--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "cert_params", "volterra_trusted_ca"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cluster/properties/tls_parameters/cert_params/volterra_trusted_ca/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["clusterCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.cert_params.volterra_trusted_ca

Breadcrumbs:

- [xcsh_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/)
- [tls_parameters.cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cluster/properties/tls_parameters/cert_params/)
- tls_parameters.cert_params.volterra_trusted_ca

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for volterra trusted ca.

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
volterra_trusted_ca = {}
```

This is an empty object or choice marker. It has no direct properties.
