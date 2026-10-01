---
page_title: "https_management.advertise_on_slo_sli.tls_config"
subcategory: ""
description: "https_management.advertise_on_slo_sli.tls_config for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 3295, "body_sha256": "sha256:237b73b41c48f7f72d06c2a4a6c87dd13277bea1543360716dba21396a992abb", "canonical_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config", "child_ids": ["xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:custom_security", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:default_security", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:low_security", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:medium_security"], "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_slo_sli", "path": "docs/guides/resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["https_management", "advertise_on_slo_sli", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_slo_sli.tls_config for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_sli.tls_config

Breadcrumbs:

- [xcsh_nfv_service](../resources/nfv_service.md)
- [Property reference](resources--nfv_service--reference.md)
- [https_management](resources--nfv_service--properties--https_management.md)
- [https_management.advertise_on_slo_sli](resources--nfv_service--properties--https_management--advertise_on_slo_sli.md)
- https_management.advertise_on_slo_sli.tls_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines various OPTIONS to configure TLS configuration parameters.

Upstream description:

This defines various OPTIONS to configure TLS configuration parameters.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_security",
    "default_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "low_security"),
  validators.ConflictingObjectAttributes("custom_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("default_security",
    "low_security"),
  validators.ConflictingObjectAttributes("default_security",
    "medium_security"),
  validators.ConflictingObjectAttributes("low_security",
    "medium_security")}
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
  "x-ves-oneof-field-choice": "[\"custom_security\",\"default_security\",\"low_security\",\"medium_security\"]"
}
```

Terraform syntax:

```terraform
tls_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--custom_security.md): complete subsection reference.

- [default_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--default_security.md): complete subsection reference.

- [low_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--low_security.md): complete subsection reference.

- [medium_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--medium_security.md): complete subsection reference.

## Next pages

- [https_management.advertise_on_slo_sli.tls_config.custom_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--custom_security.md)
- [https_management.advertise_on_slo_sli.tls_config.default_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--default_security.md)
- [https_management.advertise_on_slo_sli.tls_config.low_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--low_security.md)
- [https_management.advertise_on_slo_sli.tls_config.medium_security](resources--nfv_service--properties--https_management--advertise_on_slo_sli--tls_config--medium_security.md)
- [https_management.advertise_on_slo_sli](resources--nfv_service--properties--https_management--advertise_on_slo_sli.md)
- [xcsh_nfv_service](../resources/nfv_service.md)
