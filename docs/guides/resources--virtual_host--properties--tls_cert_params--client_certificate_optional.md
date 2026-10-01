---
page_title: "tls_cert_params.client_certificate_optional"
subcategory: ""
description: "tls_cert_params.client_certificate_optional for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1040, "body_sha256": "sha256:4dc725f18cc12d4a012d8047b0136959f78d7f7b8290aa3a06f7a9bcba1b6262", "canonical_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:client_certificate_optional", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:client_certificate_optional", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params", "path": "docs/guides/resources--virtual_host--properties--tls_cert_params--client_certificate_optional.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_cert_params", "client_certificate_optional"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_cert_params/client_certificate_optional/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_cert_params.client_certificate_optional for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_cert_params.client_certificate_optional

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [tls_cert_params](resources--virtual_host--properties--tls_cert_params.md)
- tls_cert_params.client_certificate_optional

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

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
client_certificate_optional = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_cert_params](resources--virtual_host--properties--tls_cert_params.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
