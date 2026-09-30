---
page_title: "tls_cert_params.no_client_certificate"
subcategory: ""
description: "tls_cert_params.no_client_certificate for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 923, "body_sha256": "sha256:0a44f3fc4756af55a49fee49ab34e71fcaf7f1e29f6f3fbd4cb1550e277963fe", "canonical_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:no_client_certificate", "child_ids": [], "collection_id": "xcsh-docs:resources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params:no_client_certificate", "parent_id": "xcsh-docs:resources:virtual_host:properties:tls_cert_params", "path": "docs/guides/resources--virtual_host--properties--tls_cert_params--no_client_certificate.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_cert_params", "no_client_certificate"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_host/properties/tls_cert_params/no_client_certificate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_cert_params.no_client_certificate for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_cert_params.no_client_certificate

Breadcrumbs:

- [xcsh_virtual_host](../resources/virtual_host.md)
- [Property reference](resources--virtual_host--reference.md)
- [tls_cert_params](resources--virtual_host--properties--tls_cert_params.md)
- tls_cert_params.no_client_certificate

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
no_client_certificate = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [tls_cert_params](resources--virtual_host--properties--tls_cert_params.md)
- [xcsh_virtual_host](../resources/virtual_host.md)
