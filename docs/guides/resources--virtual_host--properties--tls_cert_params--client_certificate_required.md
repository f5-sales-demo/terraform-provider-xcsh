---
page_title: "tls_cert_params.client_certificate_required"
subcategory: ""
description: "tls_cert_params.client_certificate_required for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1040, "body_sha256": "sha256:6a8f904e9c4f50a0a33d12398296e387fe53f40a1e90ee299223c302a23c0350", "canonical_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:client_certificate_required", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:client_certificate_required", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params", "path": "docs/guides/resources--virtual_host--properties--tls_cert_params--client_certificate_required.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_cert_params", "client_certificate_required"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_cert_params/client_certificate_required/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_cert_params.client_certificate_required for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# tls_cert_params.client_certificate_required

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [tls_cert_params](resources--virtual_host--properties--tls_cert_params.md)
- tls_cert_params.client_certificate_required

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
client_certificate_required = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_cert_params](resources--virtual_host--properties--tls_cert_params.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
