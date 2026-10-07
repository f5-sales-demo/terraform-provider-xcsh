---
page_title: "where"
subcategory: ""
description: "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites"
xcsh_docs: {"aliases": ["where"], "body_bytes": 2016, "body_sha256": "sha256:d06b80d3cfe831c96e9bf8c778c336bdccd7faa13f032eddc8131588c3f97cd9", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:where:site", "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "xcsh-docs:resources:advertise_policy:properties:where:virtual_site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:where", "parent_id": "xcsh-docs:resources:advertise_policy:reference", "path": "documentation/resources/advertise_policy/properties/where/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101", "registry_path": "docs/guides/resources--advertise_policy--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_network", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:virtual_network,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:virtual_network,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_site", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["where"], "schema_version": 1, "sections": [{"aliases": ["where site"], "anchor": "section", "description": "This specifies a direct reference to a site configuration object.", "document_id": "xcsh-docs:resources:advertise_policy:properties:where:site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where.site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:site:disable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:site:enable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.site:RequiredObjectAttributes:ref", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:site:ref", "type": "requires"}], "schema_path": ["where", "site"], "syntax": "block", "type": "object"}, {"aliases": ["where virtual network"], "anchor": "section", "description": "This specifies a direct reference to a network configuration object.", "document_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_network:RequiredObjectAttributes:ref", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network:ref", "type": "requires"}], "schema_path": ["where", "virtual_network"], "syntax": "block", "type": "object"}, {"aliases": ["where virtual site"], "anchor": "section", "description": "A reference to virtual_site object.", "document_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_site:disable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_site:enable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:RequiredObjectAttributes:ref", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_site:ref", "type": "requires"}], "schema_path": ["where", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/where/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/)
- where

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

NetworkSiteRefSelector defines a union of reference to site or reference to virtual\_network or
reference to virtual\_site It is used to determine virtual network using following rules \* Direct
reference to virtual\_network object \* Site local network when referring to site object \* All site
local networks for sites selected by referring to virtual\_site object.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
    "virtual_network"),
  validators.ConflictingObjectAttributes("site",
    "virtual_site"),
  validators.ConflictingObjectAttributes("virtual_network",
    "virtual_site")}
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_network\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/site/): complete subsection reference.

- [virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/virtual_network/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/advertise_policy/properties/where/virtual_site/): complete subsection reference.
