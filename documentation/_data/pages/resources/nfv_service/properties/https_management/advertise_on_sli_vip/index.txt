---
page_title: "https_management.advertise_on_sli_vip"
subcategory: ""
description: "Inline TLS parameters."
xcsh_docs: {"aliases": ["https management advertise on sli vip"], "body_bytes": 1781, "body_sha256": "sha256:dafca0b6bc591738af76c70752a92760b009893f9f4c9c53e3fcfef8b2607f08", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:no_mtls", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config", "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:use_mtls"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nfv_service:collection", "completeness": "complete", "id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip", "parent_id": "xcsh-docs:resources:nfv_service:properties:https_management", "path": "documentation/resources/nfv_service/properties/https_management/advertise_on_sli_vip/index.md", "product": "distributed-cloud", "provider_name": "nfv_service", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0331000010312010-2210033023202103-2113002300230301-3022110023310301-2232203113103003-1020202001202210-3321212230202210-0102133323033302", "registry_path": "docs/guides/resources--nfv_service--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["https_management", "advertise_on_sli_vip"], "schema_version": 1, "sections": [{"aliases": ["https management advertise on sli vip no mtls"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:no_mtls", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip", "no_mtls"], "syntax": "attribute", "type": "object"}, {"aliases": ["cert", "certificate", "existing certificates", "https management advertise on sli vip tls certificates", "tls certificates"], "anchor": "section", "description": "Users can add one or more certificates that share the same set of domains. For example, domain.com and *.domain.com - but use different signature algorithms.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_certificates", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip", "tls_certificates"], "syntax": "block", "type": "object"}, {"aliases": ["https management advertise on sli vip tls config"], "anchor": "section", "description": "This defines various OPTIONS to configure TLS configuration parameters.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:tls_config", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip", "tls_config"], "syntax": "block", "type": "object"}, {"aliases": ["https management advertise on sli vip use mtls"], "anchor": "section", "description": "Validation context for downstream client TLS connections.", "document_id": "xcsh-docs:resources:nfv_service:properties:https_management:advertise_on_sli_vip:use_mtls", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["https_management", "advertise_on_sli_vip", "use_mtls"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nfv_service/properties/https_management/advertise_on_sli_vip/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Inline TLS parameters.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["nfv_serviceCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# https_management.advertise_on_sli_vip

Breadcrumbs:

- [xcsh_nfv_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/)
- [https_management](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/)
- https_management.advertise_on_sli_vip

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Inline TLS Parameters. Inline TLS parameters.

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

- [no_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_sli_vip/no_mtls/): complete subsection reference.

- [tls_certificates](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_certificates/): complete subsection reference.

- [tls_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_sli_vip/tls_config/): complete subsection reference.

- [use_mtls](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nfv_service/properties/https_management/advertise_on_sli_vip/use_mtls/): complete subsection reference.
