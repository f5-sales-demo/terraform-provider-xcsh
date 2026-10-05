---
page_title: "where"
subcategory: ""
description: "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites"
xcsh_docs: {"aliases": ["where"], "body_bytes": 2560, "body_sha256": "sha256:dd420fe479ab8cc952f7530409f2fbdbeb38468435e5624f8e4c6fecd0cfba20", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:advertise_policy:properties:where:site", "xcsh-docs:data-sources:advertise_policy:properties:where:virtual_network", "xcsh-docs:data-sources:advertise_policy:properties:where:virtual_site"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:properties:where", "parent_id": "xcsh-docs:data-sources:advertise_policy:reference", "path": "documentation/data-sources/advertise_policy/properties/where/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3031120312202103-3312322201132331-3211232330311113-3303330100012123-2023223003011103-2303121112322110-3133101002313302-0322223323111230", "registry_path": "docs/guides/data-sources--advertise_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["where"], "schema_version": 1, "sections": [{"aliases": ["where site"], "anchor": "section", "description": "This specifies a direct reference to a site configuration object.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:where:site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "site"], "syntax": "attribute", "type": "object"}, {"aliases": ["where virtual network"], "anchor": "section", "description": "This specifies a direct reference to a network configuration object.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:where:virtual_network", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "virtual_network"], "syntax": "attribute", "type": "object"}, {"aliases": ["where virtual site"], "anchor": "section", "description": "A reference to virtual_site object.", "document_id": "xcsh-docs:data-sources:advertise_policy:properties:where:virtual_site", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["where", "virtual_site"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/properties/where/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "NetworkSiteRefSelector defines a union of reference to site or reference to virtual_network or reference to virtual_site It is used to determine virtual network using following rules * Direct reference to virtual_network object * Site local network when referring to site object * All site local networks for sites", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/)
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

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/): complete subsection reference.

- [virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/): complete subsection reference.

## Next pages

- [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/site/)
- [where.virtual_network](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_network/)
- [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/where/virtual_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/)
- [xcsh_advertise_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/)
