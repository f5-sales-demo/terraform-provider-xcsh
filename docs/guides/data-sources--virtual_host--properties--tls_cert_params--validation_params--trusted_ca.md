---
page_title: "tls_cert_params.validation_params.trusted_ca"
subcategory: ""
description: "tls_cert_params.validation_params.trusted_ca for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1320, "body_sha256": "sha256:8e59ece38622244a62db839f2a7875b95cbff668fe705af6c56068df53b965c4", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca:trusted_ca_list"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:tls_cert_params:validation_params", "path": "docs/guides/data-sources--virtual_host--properties--tls_cert_params--validation_params--trusted_ca.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_cert_params", "validation_params", "trusted_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_cert_params.validation_params.trusted_ca for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# tls_cert_params.validation_params.trusted_ca

Breadcrumbs:

- [xcsh_virtual_host](../data-sources/virtual_host.md)
- [Property reference](data-sources--virtual_host--reference.md)
- [tls_cert_params](data-sources--virtual_host--properties--tls_cert_params.md)
- [tls_cert_params.validation_params](data-sources--virtual_host--properties--tls_cert_params--validation_params.md)
- tls_cert_params.validation_params.trusted_ca

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

- [trusted_ca_list](data-sources--virtual_host--properties--tls_cert_params--validation_params--trusted_ca--trusted_ca_list.md): complete subsection reference.

## Next pages

- [tls_cert_params.validation_params.trusted_ca.trusted_ca_list](data-sources--virtual_host--properties--tls_cert_params--validation_params--trusted_ca--trusted_ca_list.md)
- [tls_cert_params.validation_params](data-sources--virtual_host--properties--tls_cert_params--validation_params.md)
- [xcsh_virtual_host](../data-sources/virtual_host.md)
