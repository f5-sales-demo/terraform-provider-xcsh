---
page_title: "site_virtual_sites"
subcategory: ""
description: "This defines a way to advertise a VIP on specific sites."
xcsh_docs: {"aliases": ["site virtual sites"], "body_bytes": 1561, "body_sha256": "sha256:e65b45a70739483acbc9920ec8e42c4c23e4e30ba8c51b1eefa456287b9a9900", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:proxy:collection", "completeness": "complete", "id": "xcsh-docs:resources:proxy:properties:site_virtual_sites", "parent_id": "xcsh-docs:resources:proxy:reference", "path": "documentation/resources/proxy/properties/site_virtual_sites/index.md", "product": "distributed-cloud", "provider_name": "proxy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2231031213302012-1012133303113311-3002123232030330-2111023130330233-3022330222131101-1133112003113022-3202002122210211-1311102012032302", "registry_path": "docs/guides/resources--proxy--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "site_virtual_sites:RequiredObjectAttributes:advertise_where", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["site_virtual_sites"], "schema_version": 1, "sections": [{"aliases": ["advertise where"], "anchor": "section", "description": "Where should this load balancer be available.", "document_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-site_virtual_sites--advertise_where--port", "enforcement": "provider-schema", "group": "site_virtual_sites.advertise_where:ConflictingListObjectAttributes:port,use_default_port", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_virtual_sites.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:site", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_virtual_sites.advertise_where:ConflictingListObjectAttributes:port,use_default_port", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:use_default_port", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "site_virtual_sites.advertise_where:ConflictingListObjectAttributes:site,virtual_site", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:proxy:properties:site_virtual_sites:advertise_where:virtual_site", "type": "conflicts"}], "schema_path": ["site_virtual_sites", "advertise_where"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/proxy/properties/site_virtual_sites/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This defines a way to advertise a VIP on specific sites.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["proxyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# site_virtual_sites

Breadcrumbs:

- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- site_virtual_sites

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Defines a way to advertise a VIP on specific sites.

Upstream description:

This defines a way to advertise a VIP on specific sites.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("advertise_where")}
```

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
site_virtual_sites {
  # Configure direct properties listed below.
}
```

## Direct properties

- [advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_virtual_sites/advertise_where/): complete subsection reference.

## Next pages

- [site_virtual_sites.advertise_where](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/site_virtual_sites/advertise_where/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/properties/)
- [xcsh_proxy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/proxy/)
