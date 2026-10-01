---
page_title: "tls_cert_params.validation_params.trusted_ca"
subcategory: ""
description: "tls_cert_params.validation_params.trusted_ca for xcsh_virtual_host."
xcsh_docs: {"aliases": [], "body_bytes": 1419, "body_sha256": "sha256:8e374fd002be0f37e6c84f3894f97dd0efa8e7cb700c68d494511c14a8d71347", "canonical_id": "xcsh-docs:data-sources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca", "child_ids": ["xcsh-docs:data-sources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca:trusted_ca_list"], "collection_id": "xcsh-docs:data-sources:virtual_host:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_host:properties:tls_cert_params:validation_params:trusted_ca", "parent_id": "xcsh-docs:data-sources:virtual_host:properties:tls_cert_params:validation_params", "path": "docs/guides/data-sources--virtual_host--properties--tls_cert_params--validation_params--trusted_ca.md", "provider_name": "virtual_host", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["tls_cert_params", "validation_params", "trusted_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_host/properties/tls_cert_params/validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "tls_cert_params.validation_params.trusted_ca for xcsh_virtual_host.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_hostCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
