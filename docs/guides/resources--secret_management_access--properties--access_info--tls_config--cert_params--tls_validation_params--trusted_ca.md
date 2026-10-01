---
page_title: "access_info.tls_config.cert_params.tls_validation_params.trusted_ca"
subcategory: ""
description: "access_info.tls_config.cert_params.tls_validation_params.trusted_ca for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 2074, "body_sha256": "sha256:6b673bc9f8064fb743d05e8a29169fb34495f99a3be74eae6ab6cb159d6471e8", "canonical_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca", "child_ids": ["xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca:trusted_ca_list"], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params:trusted_ca", "parent_id": "xcsh-docs:resources:secret_management_access:properties:access_info:tls_config:cert_params:tls_validation_params", "path": "docs/guides/resources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params--trusted_ca.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["access_info", "tls_config", "cert_params", "tls_validation_params", "trusted_ca"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/properties/access_info/tls_config/cert_params/tls_validation_params/trusted_ca/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "access_info.tls_config.cert_params.tls_validation_params.trusted_ca for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# access_info.tls_config.cert_params.tls_validation_params.trusted_ca

Breadcrumbs:

- [xcsh_secret_management_access](../resources/secret_management_access.md)
- [Property reference](resources--secret_management_access--reference.md)
- [access_info](resources--secret_management_access--properties--access_info.md)
- [access_info.tls_config](resources--secret_management_access--properties--access_info--tls_config.md)
- [access_info.tls_config.cert_params](resources--secret_management_access--properties--access_info--tls_config--cert_params.md)
- [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params.md)
- access_info.tls_config.cert_params.tls_validation_params.trusted_ca

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

- [trusted_ca_list](resources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md): complete subsection reference.

## Next pages

- [access_info.tls_config.cert_params.tls_validation_params.trusted_ca.trusted_ca_list](resources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params--trusted_ca--trusted_ca_list.md)
- [access_info.tls_config.cert_params.tls_validation_params](resources--secret_management_access--properties--access_info--tls_config--cert_params--tls_validation_params.md)
- [xcsh_secret_management_access](../resources/secret_management_access.md)
