---
page_title: "https_management.advertise_on_sli_vip.tls_config"
subcategory: ""
description: "https_management.advertise_on_sli_vip.tls_config for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2624, "body_sha256": "sha256:d7246cad87ccf1841b94b24da9f6be6f3d2c46f0c73c8413710f4c682cbc77a5", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:custom_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:default_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:low_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip", "path": "docs/guides/data-sources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_sli_vip", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_sli_vip.tls_config for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_sli_vip.tls_config

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [https_management](data-sources--nfv_service--properties--https_management.md)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--properties--https_management--advertise_on_sli_vip.md)
- https_management.advertise_on_sli_vip.tls_config

<a id="section"></a>

Type: `"single"`. Computed.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

## Direct properties

- [custom_security](data-sources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--custom_security.md): complete subsection reference.

- [default_security](data-sources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--default_security.md): complete subsection reference.

- [low_security](data-sources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--low_security.md): complete subsection reference.

- [medium_security](data-sources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_sli_vip.tls_config.custom_security](data-sources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--custom_security.md)
- [https_management.advertise_on_sli_vip.tls_config.default_security](data-sources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--default_security.md)
- [https_management.advertise_on_sli_vip.tls_config.low_security](data-sources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--low_security.md)
- [https_management.advertise_on_sli_vip.tls_config.medium_security](data-sources--nfv_service--properties--https_management--advertise_on_sli_vip--tls_config--medium_security.md)
- [https_management.advertise_on_sli_vip](data-sources--nfv_service--properties--https_management--advertise_on_sli_vip.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
