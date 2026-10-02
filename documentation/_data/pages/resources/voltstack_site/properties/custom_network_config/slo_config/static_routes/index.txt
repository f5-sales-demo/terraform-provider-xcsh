---
page_title: "custom_network_config.slo_config.static_routes"
subcategory: ""
description: "List of static routes."
xcsh_docs: {"aliases": ["custom network config slo config static routes"], "body_bytes": 2068, "body_sha256": "sha256:de0f7205c3f5ec3983f787c3564ab190d3a52974a6b474810ece57ddeacab3fa", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config", "path": "documentation/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1222233330322332-1311022211003101-0230120133122311-3030023301210101-0021032123033010-0203312312123023-3130210113323311-3100333122023020", "registry_path": "docs/guides/resources--voltstack_site--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "slo_config", "static_routes"], "schema_version": 1, "sections": [{"aliases": ["static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-custom_network_config--slo_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "choice"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes:default_gateway", "type": "choice"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface", "type": "choice"}, {"anchor": "schema-custom_network_config--slo_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "conflicts"}, {"anchor": "schema-custom_network_config--slo_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:ip_address,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:ip_address,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface", "type": "conflicts"}, {"anchor": "schema-custom_network_config--slo_config--static_routes--static_routes--ip_prefixes", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:RequiredListObjectAttributes:ip_prefixes", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "requires"}], "schema_path": ["custom_network_config", "slo_config", "static_routes", "static_routes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "List of static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config.static_routes

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/)
- custom_network_config.slo_config.static_routes

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for static routes.

Upstream description:

List of static routes.

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

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/static_routes/): complete subsection reference.

## Next pages

- [custom_network_config.slo_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/static_routes/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
