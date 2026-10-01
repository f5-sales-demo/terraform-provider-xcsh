---
page_title: "tls_parameters.common_params.validation_params.trusted_ca"
subcategory: ""
description: "tls_parameters.common_params.validation_params.trusted_ca for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1737, "body_sha256": "sha256:0847f6d73637d9b0d7876d9f3970c5bde0bbbf39aa0d1874601ef34289868e05", "canonical_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:validation_params:trusted_ca", "child_ids": ["xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:validation_params:trusted_ca:trusted_ca_list"], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:validation_params:trusted_ca", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:validation_params", "path": "docs/guides/resources--virtual_host--properties--tls_parameters--common_params--validation_params--trusted_ca.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "common_params", "validation_params", "trusted_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_parameters/common_params/validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.validation_params.trusted_ca for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.validation_params.trusted_ca

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [tls_parameters](resources--virtual_host--properties--tls_parameters.md)
- [tls_parameters.common_params](resources--virtual_host--properties--tls_parameters--common_params.md)
- [tls_parameters.common_params.validation_params](resources--virtual_host--properties--tls_parameters--common_params--validation_params.md)
- tls_parameters.common_params.validation_params.trusted_ca

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

- [trusted_ca_list](resources--virtual_host--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md): complete subsection reference.

## Next pages

- [tls_parameters.common_params.validation_params.trusted_ca.trusted_ca_list](resources--virtual_host--properties--tls_parameters--common_params--validation_params--trusted_ca--trusted_ca_list.md)
- [tls_parameters.common_params.validation_params](resources--virtual_host--properties--tls_parameters--common_params--validation_params.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
