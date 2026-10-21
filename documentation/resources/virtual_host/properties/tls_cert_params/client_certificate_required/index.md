---
page_title: "tls_cert_params.client_certificate_required"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["tls cert params client certificate required"], "body_bytes": 1035, "body_sha256": "sha256:ee93ed16ce50abc4119beaef60911c5c24a787301205cfdd723faff854466f84", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:client_certificate_required", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params", "path": "documentation/resources/virtual_host/properties/tls_cert_params/client_certificate_required/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1003023201132001-2031223012013002-0033031311323313-1000012312300110-3220100130333033-2010311200120011-0223020021323130-0212233213122231", "registry_path": "docs/guides/resources--virtual_host--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_cert_params", "client_certificate_required"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_cert_params/client_certificate_required/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_cert_params.client_certificate_required

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/)
- tls_cert_params.client_certificate_required

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
client_certificate_required = {}
```

This is an empty object or choice marker. It has no direct properties.
