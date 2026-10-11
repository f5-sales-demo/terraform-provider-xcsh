---
page_title: "where"
subcategory: ""
description: "VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual_site It used to refer site or a group of sites indicated by virtual site."
xcsh_docs: {"aliases": ["where"], "body_bytes": 1401, "body_sha256": "sha256:55166ead1dc7192cca7fb6c8ec33e30263f577cd580deba8a8970933c5ea316f", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:where:site", "xcsh-docs:resources:bgp:properties:where:virtual_site"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:where", "parent_id": "xcsh-docs:resources:bgp:reference", "path": "documentation/resources/bgp/properties/where/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231", "registry_path": "docs/guides/resources--bgp--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["where"], "schema_version": 1, "sections": [{"aliases": ["where site"], "anchor": "section", "description": "This specifies a direct reference to a site configuration object.", "document_id": "xcsh-docs:resources:bgp:properties:where:site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where.site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:site:disable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:site:enable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.site:RequiredObjectAttributes:ref", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:site:ref", "type": "requires"}], "schema_path": ["where", "site"], "syntax": "block", "type": "object"}, {"aliases": ["where virtual site"], "anchor": "section", "description": "A reference to virtual_site object.", "document_id": "xcsh-docs:resources:bgp:properties:where:virtual_site", "enum_extraction_complete": false, "enum_validators": [], "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:disable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:enable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:RequiredObjectAttributes:ref", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:ref", "type": "requires"}], "schema_path": ["where", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/where/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual_site It used to refer site or a group of sites indicated by virtual site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["bgpCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# where

Breadcrumbs:

- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- where

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Provider validators and defaults (from schema source):

```go
EnumExtractionComplete: false
Validators: []validator.Object{validators.ConflictingObjectAttributes("site",
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
  "x-ves-oneof-field-ref_or_selector": "[\"site\",\"virtual_site\"]"
}
```

Terraform syntax:

```terraform
where {
  # Configure direct properties listed below.
}
```

## Direct properties

- [site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/site/): complete subsection reference.

- [virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/virtual_site/): complete subsection reference.
