---
page_title: "tls_parameters.common_params.tls_certificates.disable_ocsp_stapling"
subcategory: ""
description: "tls_parameters.common_params.tls_certificates.disable_ocsp_stapling for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1417, "body_sha256": "sha256:33f0746e55833a4815ad98fd3b9475e88c9c93fe716d7ddb47aff59b5ef2b43e", "canonical_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:disable_ocsp_stapling", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates:disable_ocsp_stapling", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_parameters:common_params:tls_certificates", "path": "docs/guides/resources--virtual_host--properties--tls_parameters--common_params--tls_certificates--disable_ocsp_stapling.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_parameters", "common_params", "tls_certificates", "disable_ocsp_stapling"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_parameters/common_params/tls_certificates/disable_ocsp_stapling/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_parameters.common_params.tls_certificates.disable_ocsp_stapling for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [tls_parameters](resources--virtual_host--properties--tls_parameters.md)
- [tls_parameters.common_params](resources--virtual_host--properties--tls_parameters--common_params.md)
- [tls_parameters.common_params.tls_certificates](resources--virtual_host--properties--tls_parameters--common_params--tls_certificates.md)
- tls_parameters.common_params.tls_certificates.disable_ocsp_stapling

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable ocsp stapling.

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
disable_ocsp_stapling = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_parameters.common_params.tls_certificates](resources--virtual_host--properties--tls_parameters--common_params--tls_certificates.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
