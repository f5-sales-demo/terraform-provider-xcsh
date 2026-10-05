---
page_title: "https_management.advertise_on_sli_vip.tls_config"
subcategory: ""
description: "This defines various OPTIONS to configure TLS configuration parameters."
xcsh_docs: {"aliases": ["https management advertise on sli vip tls config"], "body_bytes": 3314, "body_sha256": "sha256:5739b0c006b872f9d6da04dd444962691349335ca47c85f56cb10807632c90a8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:custom_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:default_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:low_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:medium_security"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip", "path": "documentation/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0110322210323010-2020033212112220-2221001300203322-1221033012030103-3023103310321232-3310201013003121-3032200300302203-0233131002303201", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_sli_vip", "tls_config"], "schema_version": 1, "sections": [{"aliases": ["https management advertise on sli vip tls config custom security"], "anchor": "section", "description": "This defines TLS protocol config including min/max versions and allowed ciphers.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:custom_security", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip", "tls_config", "custom_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on sli vip tls config default security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:default_security", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip", "tls_config", "default_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on sli vip tls config low security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:low_security", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip", "tls_config", "low_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on sli vip tls config medium security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config:medium_security", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip", "tls_config", "medium_security"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "This defines various OPTIONS to configure TLS configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_sli_vip.tls_config

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/)
- [https_management.advertise_on_sli_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/)
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

- [custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/custom_security/): complete subsection reference.

- [default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/default_security/): complete subsection reference.

- [low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/low_security/): complete subsection reference.

- [medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/medium_security/): complete subsection reference.

## Next pages

- [https_management.advertise_on_sli_vip.tls_config.custom_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/custom_security/)
- [https_management.advertise_on_sli_vip.tls_config.default_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/default_security/)
- [https_management.advertise_on_sli_vip.tls_config.low_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/low_security/)
- [https_management.advertise_on_sli_vip.tls_config.medium_security](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/medium_security/)
- [https_management.advertise_on_sli_vip](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/properties/https_management/advertise_on_sli_vip/)
- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nfv_service/)
