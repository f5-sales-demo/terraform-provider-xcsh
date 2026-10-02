---
page_title: "where"
subcategory: ""
description: "VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual_site It used to refer site or a group of sites indicated by virtual site."
xcsh_docs: {"aliases": ["where"], "body_bytes": 1999, "body_sha256": "sha256:69c2bacbc3c972e19b473c1afce010fb0d2328745bc410c2484bb7b6d6c8eb7c", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:properties:where:site", "xcsh-docs:resources:bgp:properties:where:virtual_site"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:properties:where", "parent_id": "xcsh-docs:resources:bgp:reference", "path": "documentation/resources/bgp/properties/where/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3310113311010321-2211233202102322-3001232112032331-0322032021203310-0031203331133113-2121220003210231-0303000101131332-1233013013233231", "registry_path": "docs/guides/resources--bgp--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where:ConflictingObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["where"], "schema_version": 1, "sections": [{"aliases": ["site"], "anchor": "section", "description": "This specifies a direct reference to a site configuration object.", "document_id": "xcsh-docs:resources:bgp:properties:where:site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where.site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:site:disable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:site:enable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.site:RequiredObjectAttributes:ref", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:site:ref", "type": "requires"}], "schema_path": ["where", "site"], "syntax": "block", "type": "object"}, {"aliases": ["virtual site"], "anchor": "section", "description": "A reference to virtual_site object.", "document_id": "xcsh-docs:resources:bgp:properties:where:virtual_site", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:disable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:ConflictingObjectAttributes:disable_internet_vip,enable_internet_vip", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:enable_internet_vip", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "where.virtual_site:RequiredObjectAttributes:ref", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:bgp:properties:where:virtual_site:ref", "type": "requires"}], "schema_path": ["where", "virtual_site"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/properties/where/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual_site It used to refer site or a group of sites indicated by virtual site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["bgpCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Upstream description:

VirtualSiteSiteRefSelector defines a union of reference to site or reference to virtual\_site It
used to refer site or a group of sites indicated by virtual site.

Provider validators and defaults (from schema source):

```go
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

## Next pages

- [where.site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/site/)
- [where.virtual_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/where/virtual_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [xcsh_bgp](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/)
