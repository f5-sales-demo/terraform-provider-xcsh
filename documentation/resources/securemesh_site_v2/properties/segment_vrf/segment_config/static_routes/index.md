---
page_title: "segment_vrf.segment_config.static_routes"
subcategory: ""
description: "Configuration parameter for static routes."
xcsh_docs: {"aliases": ["segment vrf segment config static routes"], "body_bytes": 1987, "body_sha256": "sha256:c99d642a8cefa3ebe8a3623c3aabbe779b35db1500c0877e7998281c40a484ff", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config", "path": "documentation/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3120301202101232-2230103100103301-2121131023233002-0120102301001221-2223230202130320-3222131230221031-1220300123020010-0002320201312113", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["segment_vrf", "segment_config", "static_routes"], "schema_version": 1, "sections": [{"aliases": ["static routes"], "anchor": "section", "description": "Configuration parameter for static routes", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-segment_vrf--segment_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes", "type": "choice"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:default_gateway", "type": "choice"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface", "type": "choice"}, {"anchor": "schema-segment_vrf--segment_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes", "type": "conflicts"}, {"anchor": "schema-segment_vrf--segment_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes.static_routes:ConflictingListObjectAttributes:ip_address,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes.static_routes:ConflictingListObjectAttributes:ip_address,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes:node_interface", "type": "conflicts"}, {"anchor": "schema-segment_vrf--segment_config--static_routes--static_routes--ip_prefixes", "enforcement": "provider-schema", "group": "segment_vrf.segment_config.static_routes.static_routes:RequiredListObjectAttributes:ip_prefixes", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:segment_vrf:segment_config:static_routes:static_routes", "type": "requires"}], "schema_path": ["segment_vrf", "segment_config", "static_routes", "static_routes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# segment_vrf.segment_config.static_routes

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [segment_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/)
- [segment_vrf.segment_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/)
- segment_vrf.segment_config.static_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("static_routes")}
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
static_routes {
  # Configure direct properties listed below.
}
```

## Direct properties

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/): complete subsection reference.

## Next pages

- [segment_vrf.segment_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/static_routes/static_routes/)
- [segment_vrf.segment_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/segment_vrf/segment_config/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
