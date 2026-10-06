---
page_title: "dns_ntp_config"
subcategory: ""
description: "Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site."
xcsh_docs: {"aliases": ["dns ntp config"], "body_bytes": 1989, "body_sha256": "sha256:48c49017e4f955cf0e726c0a98040fd36c3a5e5dc49e5d54ed8d2ef40606b841", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:f5_dns_default", "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:f5_ntp_default"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/dns_ntp_config/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0212020030133010-3223023110132101-2003232322020333-0130023111013232-0023300023001223-0131110322121120-0330310333311013-0300201302020000", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-006.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "dns_ntp_config:ConflictingObjectAttributes:custom_dns,f5_dns_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dns_ntp_config:ConflictingObjectAttributes:custom_ntp,f5_ntp_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dns_ntp_config:ConflictingObjectAttributes:custom_dns,f5_dns_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:f5_dns_default", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "dns_ntp_config:ConflictingObjectAttributes:custom_ntp,f5_ntp_default", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:f5_ntp_default", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["dns_ntp_config"], "schema_version": 1, "sections": [{"aliases": ["dns ntp config custom dns"], "anchor": "section", "description": "DNS Servers.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:custom_dns", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dns_ntp_config", "custom_dns"], "syntax": "block", "type": "object"}, {"aliases": ["dns ntp config custom ntp"], "anchor": "section", "description": "NTP Servers.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:custom_ntp", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["dns_ntp_config", "custom_ntp"], "syntax": "block", "type": "object"}, {"aliases": ["dns ntp config f5 dns default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:f5_dns_default", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_ntp_config", "f5_dns_default"], "syntax": "attribute", "type": "object"}, {"aliases": ["dns ntp config f5 ntp default"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:dns_ntp_config:f5_ntp_default", "enum_extraction_complete": false, "enum_validators": [], "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["dns_ntp_config", "f5_ntp_default"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/dns_ntp_config/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# dns_ntp_config

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- dns_ntp_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify DNS and NTP servers that will be used by the nodes in this Customer Edge site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("custom_dns",
    "f5_dns_default"),
  validators.ConflictingObjectAttributes("custom_ntp",
    "f5_ntp_default")}
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
  "x-ves-oneof-field-dns_server_choice": "[\"custom_dns\",\"f5_dns_default\"]",
  "x-ves-oneof-field-ntp_server_choice": "[\"custom_ntp\",\"f5_ntp_default\"]"
}
```

Terraform syntax:

```terraform
dns_ntp_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [custom_dns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/custom_dns/): complete subsection reference.

- [custom_ntp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/custom_ntp/): complete subsection reference.

- [f5_dns_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/f5_dns_default/): complete subsection reference.

- [f5_ntp_default](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/dns_ntp_config/f5_ntp_default/): complete subsection reference.
