---
page_title: "where"
subcategory: ""
description: "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites"
xcsh_docs: {"aliases": ["where"], "body_bytes": 1666, "body_sha256": "sha256:c13568ea9436d69878c38a5c7a447e15a7c1db509c1e4912a98ec1359aadb33c", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:advertise_policy:properties:where:site", "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "xcsh-docs:resources:advertise_policy:properties:where:virtual_site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:advertise_policy:properties:where", "parent_id": "xcsh-docs:resources:advertise_policy:reference", "path": "documentation/resources/advertise_policy/properties/where/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-1132021011013021-1001033122102311-2311002131222013-0101112131230102-3330100301313310-2331313131211020-1323310311022023-1012213313100101", "registry_path": "docs/guides/resources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where"], "schema_version": 1, "sections": [{"aliases": ["where site"], "anchor": "section", "description": "This specifies a direct reference to a site configuration object.", "document_id": "xcsh-docs:resources:advertise_policy:properties:where:site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "site"], "syntax": "block", "type": "object"}, {"aliases": ["where virtual network"], "anchor": "section", "description": "This specifies a direct reference to a network configuration object.", "document_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_network", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "virtual_network"], "syntax": "block", "type": "object"}, {"aliases": ["where virtual site"], "anchor": "section", "description": "A reference to virtual_site object.", "document_id": "xcsh-docs:resources:advertise_policy:properties:where:virtual_site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/advertise_policy/properties/where/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
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
