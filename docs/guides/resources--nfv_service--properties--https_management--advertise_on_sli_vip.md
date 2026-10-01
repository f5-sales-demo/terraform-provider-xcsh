---
page_title: "https_management.advertise_on_sli_vip"
subcategory: ""
description: "https_management.advertise_on_sli_vip for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2423, "body_sha256": "sha256:15509921a444c251828f5312297275554e0a1e243d8080b0cc547c951eb448bf", "canonical_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip", "child_ids": ["xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:no_mtls", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:use_mtls"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management", "path": "docs/guides/resources--nfv_service--properties--https_management--advertise_on_sli_vip.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_sli_vip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_sli_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_sli_vip for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_sli_vip

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [https_management](resources--nfv_service--properties--https_management.md)
- https_management.advertise_on_sli_vip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("tls_certificates"),
  validators.ConflictingObjectAttributes("no_mtls",
    "use_mtls")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-mtls_choice": "[\"no_mtls\",\"use_mtls\"]"
}
```

Terraform syntax:

```terraform
advertise_on_sli_vip {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_mtls](resources--nfv_service--properties--https_management--advertise_on_sli_vip--no_mtls.md): complete subsection reference.

- [tls_certificates](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates.md): complete subsection reference.

- [tls_config](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config.md): complete subsection reference.

- [use_mtls](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_sli_vip.no_mtls](resources--nfv_service--properties--https_management--advertise_on_sli_vip--no_mtls.md)
- [https_management.advertise_on_sli_vip.tls_certificates](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_certificates.md)
- [https_management.advertise_on_sli_vip.tls_config](resources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config.md)
- [https_management.advertise_on_sli_vip.use_mtls](resources--nfv_service--properties--https_management--advertise_on_sli_vip--use_mtls.md)
- [https_management](resources--nfv_service--properties--https_management.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
