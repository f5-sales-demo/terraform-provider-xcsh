---
page_title: "where"
subcategory: ""
description: "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites"
xcsh_docs: {"aliases": ["where"], "body_bytes": 1525, "body_sha256": "sha256:6912d2b463e69fda89817f9eb8c0dbad2ec40a84416b573735990a029ea409dd", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:where:site", "xcsh-docs:data-sources:discovery:properties:where:virtual_network", "xcsh-docs:data-sources:discovery:properties:where:virtual_site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:where", "parent_id": "xcsh-docs:data-sources:discovery:reference", "path": "documentation/data-sources/discovery/properties/where/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110", "registry_path": "docs/guides/data-sources--discovery--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where"], "schema_version": 1, "sections": [{"aliases": ["where site"], "anchor": "section", "description": "This specifies a direct reference to a site configuration object.", "document_id": "xcsh-docs:data-sources:discovery:properties:where:site", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["where virtual network"], "anchor": "section", "description": "This specifies a direct reference to a network configuration object.", "document_id": "xcsh-docs:data-sources:discovery:properties:where:virtual_network", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "virtual_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["where virtual site"], "anchor": "section", "description": "A reference to virtual_site object.", "document_id": "xcsh-docs:data-sources:discovery:properties:where:virtual_site", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/where/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["discoveryCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- where

<a id="section"></a>

Type: `"single"`. Computed.

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

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/where/site/): complete subsection reference.

- [virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/where/virtual_network/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/where/virtual_site/): complete subsection reference.
