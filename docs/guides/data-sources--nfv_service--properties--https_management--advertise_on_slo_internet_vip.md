---
page_title: "https_management.advertise_on_slo_internet_vip"
subcategory: ""
description: "https_management.advertise_on_slo_internet_vip for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2235, "body_sha256": "sha256:5d448c98231d245eb5f7aca549102af3698311d373f9657919700ea58c19a081", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_internet_vip", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:no_mtls", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_certificates", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:tls_config", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_internet_vip:use_mtls"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_internet_vip", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management", "path": "docs/guides/data-sources--nfv_service--properties--https_management--advertise_on_slo_internet_vip.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_slo_internet_vip"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_slo_internet_vip/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_slo_internet_vip for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_internet_vip

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [https_management](data-sources--nfv_service--properties--https_management.md)
- https_management.advertise_on_slo_internet_vip

<a id="section"></a>

Type: `"single"`. Computed.

Inline TLS Parameters. Inline TLS parameters.

Upstream description:

Inline TLS parameters.

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

## Direct properties

- [no_mtls](data-sources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--no_mtls.md): complete subsection reference.

- [tls_certificates](data-sources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates.md): complete subsection reference.

- [tls_config](data-sources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_config.md): complete subsection reference.

- [use_mtls](data-sources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_slo_internet_vip.no_mtls](data-sources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--no_mtls.md)
- [https_management.advertise_on_slo_internet_vip.tls_certificates](data-sources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_certificates.md)
- [https_management.advertise_on_slo_internet_vip.tls_config](data-sources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--tls_config.md)
- [https_management.advertise_on_slo_internet_vip.use_mtls](data-sources--nfv_service--properties--https_management--advertise_on_slo_internet_vip--use_mtls.md)
- [https_management](data-sources--nfv_service--properties--https_management.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
