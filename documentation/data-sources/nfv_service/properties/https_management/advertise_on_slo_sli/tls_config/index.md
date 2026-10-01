---
page_title: "https_management.advertise_on_slo_sli.tls_config"
subcategory: ""
description: "https_management.advertise_on_slo_sli.tls_config for xcsh_nfv_service."
xcsh_docs: {"aliases": [], "body_bytes": 3314, "body_sha256": "sha256:a82412805a0ff0a92a3ba523e63ca2c4d207b3452f2e7db5218a9dad4b5992e0", "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:custom_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:default_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:low_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:medium_security"], "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli", "path": "documentation/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/index.md", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "properties", "schema_path": ["https_management", "advertise_on_slo_sli", "tls_config"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "https_management.advertise_on_slo_sli.tls_config for xcsh_nfv_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_sli.tls_config

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/)
- [https_management.advertise_on_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/)
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

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/medium_security/): complete subsection reference.

## Next pages

- [https_management.advertise_on_slo_sli.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/custom_security/)
- [https_management.advertise_on_slo_sli.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/default_security/)
- [https_management.advertise_on_slo_sli.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/low_security/)
- [https_management.advertise_on_slo_sli.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/medium_security/)
- [https_management.advertise_on_slo_sli](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
