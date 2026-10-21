---
page_title: "tls_parameters.client_certificate_optional"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls parameters client certificate optional"], "body_bytes": 1047, "body_sha256": "sha256:beb90bd601526ab956eadcd13ca76c5ff700d9d4e957f9312fed468e6f5c3907", "capabilities": ["load-balancing.tls", "networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters:client_certificate_optional", "parent_id": "xcsh-docs:resources:advertise_policy:properties:tls_parameters", "path": "documentation/resources/advertise_policy/properties/tls_parameters/client_certificate_optional/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3011211000321121-1203211002310022-1301212313121113-1320112310012002-2011003113031221-3102303113123230-1103303232223311-2132302212012103", "registry_path": "docs/guides/resources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "client_certificate_optional"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/tls_parameters/client_certificate_optional/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.client_certificate_optional

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/tls_parameters/)
- tls_parameters.client_certificate_optional

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
client_certificate_optional = {}
```

This is an empty object or choice marker. It has no direct properties.
