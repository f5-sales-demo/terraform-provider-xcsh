---
page_title: "local_vrf.slo_config.static_routes"
subcategory: ""
description: "Configuration parameter for static routes."
xcsh_docs: {"aliases": ["local vrf slo config static routes"], "body_bytes": 1929, "body_sha256": "sha256:734af729c26047c7a396763ba917421584bbbf30bc3971ac9203ea93abc7e9c6", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config", "path": "documentation/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_routes/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0023202222232300-1110220133212031-3222010110202222-0012232022210323-2213302012231302-1111020202332120-1233301230212203-2012210333111210", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-012.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_vrf", "slo_config", "static_routes"], "schema_version": 1, "sections": [{"aliases": ["static routes"], "anchor": "section", "description": "Configuration parameter for static routes", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-local_vrf--slo_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes", "type": "choice"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes:default_gateway", "type": "choice"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes:node_interface", "type": "choice"}, {"anchor": "schema-local_vrf--slo_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes", "type": "conflicts"}, {"anchor": "schema-local_vrf--slo_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:ip_address,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:ip_address,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes:node_interface", "type": "conflicts"}, {"anchor": "schema-local_vrf--slo_config--static_routes--static_routes--ip_prefixes", "enforcement": "provider-schema", "group": "local_vrf.slo_config.static_routes.static_routes:RequiredListObjectAttributes:ip_prefixes", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:local_vrf:slo_config:static_routes:static_routes", "type": "requires"}], "schema_path": ["local_vrf", "slo_config", "static_routes", "static_routes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_routes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Configuration parameter for static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_vrf.slo_config.static_routes

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [local_vrf](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/)
- [local_vrf.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/)
- local_vrf.slo_config.static_routes

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

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_routes/static_routes/): complete subsection reference.

## Next pages

- [local_vrf.slo_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/static_routes/static_routes/)
- [local_vrf.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/local_vrf/slo_config/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
