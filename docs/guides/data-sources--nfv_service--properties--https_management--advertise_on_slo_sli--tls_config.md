---
page_title: "https_management.advertise_on_slo_sli.tls_config"
subcategory: ""
description: "https_management.advertise_on_slo_sli.tls_config for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 2525, "body_sha256": "sha256:3e64f1dd3fe428485607e0b6b7eda50cb2fd8bbc2771b2c1450fc20cc9f3744d", "canonical_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:custom_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:default_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:low_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli", "path": "docs/guides/data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_slo_sli", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_slo_sli.tls_config for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# https_management.advertise_on_slo_sli.tls_config

Breadcrumbs:

- [xcsh_nfv_service](../data-sources/nfv_service.md)
- [Property reference](data-sources--nfv_service--reference.md)
- [https_management](data-sources--nfv_service--properties--https_management.md)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli.md)
- https_management.advertise_on_slo_sli.tls_config

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

- [custom_security](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--custom_security.md): complete subsection reference.

- [default_security](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--default_security.md): complete subsection reference.

- [low_security](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--low_security.md): complete subsection reference.

- [medium_security](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_slo_sli.tls_config.custom_security](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--custom_security.md)
- [https_management.advertise_on_slo_sli.tls_config.default_security](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--default_security.md)
- [https_management.advertise_on_slo_sli.tls_config.low_security](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--low_security.md)
- [https_management.advertise_on_slo_sli.tls_config.medium_security](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--medium_security.md)
- [https_management.advertise_on_slo_sli](data-sources--nfv_service--properties--https_management--advertise_on_slo_sli.md)
- [xcsh_nfv_service](../data-sources/nfv_service.md)
