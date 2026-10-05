---
page_title: "tls_cert_params.validation_params.trusted_ca"
subcategory: ""
description: "Reference to Root CA Certificate."
xcsh_docs: {"aliases": ["tls cert params validation params trusted ca"], "body_bytes": 1916, "body_sha256": "sha256:02caf6b80b68d14486290f9264f0c7d0dd7629c18909e4e0958c29816f9d3ad5", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca:trusted_ca_list"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params", "path": "documentation/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3102322203102032-1010001013023210-2012213032232022-3012220300112201-1212033212213030-3021001301210011-1330200222200322-0000203010231131", "registry_path": "docs/guides/resources--virtual_host--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_cert_params", "validation_params", "trusted_ca"], "schema_version": 1, "sections": [{"aliases": ["tls cert params validation params trusted ca trusted ca list"], "anchor": "section", "description": "Reference to Root CA Certificate.", "document_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca:trusted_ca_list", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["tls_cert_params", "validation_params", "trusted_ca", "trusted_ca_list"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Reference to Root CA Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_cert_params.validation_params.trusted_ca

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [tls_cert_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/)
- [tls_cert_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/)
- tls_cert_params.validation_params.trusted_ca

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Root CA Certificate Reference. Reference to Root CA Certificate.

Upstream description:

Reference to Root CA Certificate.

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
trusted_ca {
  # Configure direct properties listed below.
}
```

## Direct properties

- [trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/trusted_ca_list/): complete subsection reference.

## Next pages

- [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/trusted_ca_list/)
- [tls_cert_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_cert_params/validation_params/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
