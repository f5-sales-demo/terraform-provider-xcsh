---
page_title: "https_management.advertise_on_slo_vip.tls_config"
subcategory: ""
description: "This defines various OPTIONS to configure TLS configuration parameters."
xcsh_docs: {"aliases": ["https management advertise on slo vip tls config"], "body_bytes": 3314, "body_sha256": "sha256:329ff6aa6c19a74de09addc09c06fe1da5eca25cfe1266170c8cb44d9e40b408", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_config:custom_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_config:default_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_config:low_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_config:medium_security"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_config", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip", "path": "documentation/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_config/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3130313110212202-3233203131203031-1300221202131002-3112121030130200-3311320033012120-0333301110000020-2023133110202300-1230200211303322", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_slo_vip", "tls_config"], "schema_version": 1, "sections": [{"aliases": ["https management advertise on slo vip tls config custom security"], "anchor": "section", "description": "This defines TLS protocol config including min/max versions and allowed ciphers.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_config:custom_security", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_slo_vip", "tls_config", "custom_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on slo vip tls config default security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_config:default_security", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_vip", "tls_config", "default_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on slo vip tls config low security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_config:low_security", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_vip", "tls_config", "low_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on slo vip tls config medium security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_vip:tls_config:medium_security", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_vip", "tls_config", "medium_security"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines various OPTIONS to configure TLS configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_slo_vip.tls_config

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/)
- [https_management.advertise_on_slo_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/)
- https_management.advertise_on_slo_vip.tls_config

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

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_config/medium_security/): complete subsection reference.

## Next pages

- [https_management.advertise_on_slo_vip.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_config/custom_security/)
- [https_management.advertise_on_slo_vip.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_config/default_security/)
- [https_management.advertise_on_slo_vip.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_config/low_security/)
- [https_management.advertise_on_slo_vip.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/tls_config/medium_security/)
- [https_management.advertise_on_slo_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_slo_vip/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
