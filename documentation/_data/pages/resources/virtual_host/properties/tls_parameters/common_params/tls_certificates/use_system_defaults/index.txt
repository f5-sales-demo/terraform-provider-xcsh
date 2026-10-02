---
page_title: "tls_parameters.common_params.tls_certificates.use_system_defaults"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["cert", "certificate", "existing certificates", "tls certificates", "tls parameters common params tls certificates use system defaults"], "body_bytes": 1763, "body_sha256": "sha256:ae14276d5f82b5cfef8395ec8cfc1e9179c76c080c52d50c09d52c58d91d6f6a", "capabilities": ["load-balancing", "load-balancing.tls"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:use_system_defaults", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates", "path": "documentation/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/index.md", "product": "distributed-cloud", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1321321311010002-0333130122331222-0233010210103222-1020300033133113-0220033020222102-2031100330302111-3223213023203230-1020203123301203", "registry_path": "docs/guides/resources--virtual_host--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "use_system_defaults"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/use_system_defaults/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates.use_system_defaults

Breadcrumbs:

- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/)
- [tls_parameters](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/)
- [tls_parameters.common_params](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/)
- [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/)
- tls_parameters.common_params.tls_certificates.use_system_defaults

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for use system defaults.

Upstream description:

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
use_system_defaults = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_parameters.common_params.tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/)
- [xcsh_virtual_host](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_host/)
