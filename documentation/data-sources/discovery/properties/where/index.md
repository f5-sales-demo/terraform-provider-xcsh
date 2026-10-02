---
page_title: "where"
subcategory: ""
description: "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites"
xcsh_docs: {"aliases": ["where"], "body_bytes": 2476, "body_sha256": "sha256:bc73fb849fe9ac4faa61a395c2071229cad6daca15343f687d9da2e08fce1464", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:where:site", "xcsh-docs:data-sources:discovery:properties:where:virtual_network", "xcsh-docs:data-sources:discovery:properties:where:virtual_site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:where", "parent_id": "xcsh-docs:data-sources:discovery:reference", "path": "documentation/data-sources/discovery/properties/where/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2012322022221322-1311331021011113-0130102221212012-0201002002131232-0010230001123303-2302100002002110-0033013201231311-0223310320002110", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "This specifies a direct reference to a site configuration object.", "document_id": "xcsh-docs:data-sources:discovery:properties:where:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual network"], "anchor": "section", "description": "This specifies a direct reference to a network configuration object.", "document_id": "xcsh-docs:data-sources:discovery:properties:where:virtual_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "virtual_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "A reference to virtual_site object.", "document_id": "xcsh-docs:data-sources:discovery:properties:where:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/where/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
local..

Upstream description:

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

## Next pages

- [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/where/site/)
- [where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/where/virtual_network/)
- [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/where/virtual_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
