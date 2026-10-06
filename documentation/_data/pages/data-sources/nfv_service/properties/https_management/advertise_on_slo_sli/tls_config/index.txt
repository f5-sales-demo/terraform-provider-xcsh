---
page_title: "https_management.advertise_on_slo_sli.tls_config"
subcategory: ""
description: "This defines various OPTIONS to configure TLS configuration parameters."
xcsh_docs: {"aliases": ["https management advertise on slo sli tls config"], "body_bytes": 2021, "body_sha256": "sha256:b8eaf4780898c1f70cdf6017d3ea895c99d5b56157a579daa788c39d1a7b51fc", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:custom_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:default_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:low_security", "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:medium_security"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config", "parent_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli", "path": "documentation/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-0013213311020323-3322133321210311-0133311201331113-3030131213302302-3223013203033222-1301222311322010-3202011111310311-3111212023222102", "registry_path": "docs/guides/data-sources--nfv_service--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_slo_sli", "tls_config"], "schema_version": 1, "sections": [{"aliases": ["https management advertise on slo sli tls config custom security"], "anchor": "section", "description": "This defines TLS protocol config including min/max versions and allowed ciphers.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:custom_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_slo_sli", "tls_config", "custom_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on slo sli tls config default security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:default_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_sli", "tls_config", "default_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on slo sli tls config low security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:low_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_sli", "tls_config", "low_security"], "syntax": "attribute", "type": "object"}, {"aliases": ["https management advertise on slo sli tls config medium security"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:nfv_service:properties:https_management:advertise_on_slo_sli:tls_config:medium_security", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_slo_sli", "tls_config", "medium_security"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nfv_service/properties/https_management/advertise_on_slo_sli/tls_config/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This defines various OPTIONS to configure TLS configuration parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
