---
page_title: "tls_parameters.common_params.validation_params.trusted_ca"
subcategory: ""
description: "Reference to Root CA Certificate."
xcsh_docs: {"aliases": ["tls parameters common params validation params trusted ca"], "body_bytes": 2093, "body_sha256": "sha256:471eb8237e3b7211bb8c8ecfda67edef27f26607fbb3e6931e3a5010d9c02a8e", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:validation_params:trusted_ca:trusted_ca_list"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:validation_params:trusted_ca", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:validation_params", "path": "documentation/data-sources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3333331333311130-0131012003032211-0001030100130121-3211202103121133-2033233301023120-2201133131321102-3012333212020221-1322102322122012", "registry_path": "docs/guides/data-sources--virtual_host--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "common_params", "validation_params", "trusted_ca"], "schema_version": 1, "sections": [{"aliases": ["tls parameters common params validation params trusted ca trusted ca list"], "anchor": "section", "description": "Reference to Root CA Certificate.", "document_id": "xcsh-docs:data-sources:virtual_host:properties:tls_parameters:common_params:validation_params:trusted_ca:trusted_ca_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["tls_parameters", "common_params", "validation_params", "trusted_ca", "trusted_ca_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reference to Root CA Certificate.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.validation_params.trusted_ca

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/)
- [tls_parameters.common_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/validation_params/)
- tls_parameters.common_params.validation_params.trusted_ca

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/): complete subsection reference.

## Next pages

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/trusted_ca_list/)
- [tls_parameters.common_params.validation_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/properties/tls_parameters/common_params/validation_params/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_host/)
