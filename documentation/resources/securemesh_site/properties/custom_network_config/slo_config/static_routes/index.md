---
page_title: "custom_network_config.slo_config.static_routes"
subcategory: ""
description: "List of static routes."
xcsh_docs: {"aliases": ["custom network config slo config static routes"], "body_bytes": 2078, "body_sha256": "sha256:f9b8a8724748eb41c07da8380e6cef071825d77e59e1a929edc240186f757b56", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes", "parent_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config", "path": "documentation/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3010113203231221-3123220110230202-2113112333002310-2031202012033131-0200322211112112-1130010030313220-0013120321130120-0122012002320022", "registry_path": "docs/guides/resources--securemesh_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "slo_config", "static_routes"], "schema_version": 1, "sections": [{"aliases": ["static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "list", "relationships": [{"anchor": "schema-custom_network_config--slo_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "choice"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:default_gateway", "type": "choice"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:RequiredOneOfListObjectAttributes:default_gateway,ip_address,node_interface", "source": "ast-validator:RequiredOneOfListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface", "type": "choice"}, {"anchor": "schema-custom_network_config--slo_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "conflicts"}, {"anchor": "schema-custom_network_config--slo_config--static_routes--static_routes--ip_address", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:ip_address,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,ip_address", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:default_gateway", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:default_gateway,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:ConflictingListObjectAttributes:ip_address,node_interface", "source": "ast-validator:ConflictingListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes:node_interface", "type": "conflicts"}, {"anchor": "schema-custom_network_config--slo_config--static_routes--static_routes--ip_prefixes", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes.static_routes:RequiredListObjectAttributes:ip_prefixes", "source": "ast-validator:RequiredListObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "requires"}], "schema_path": ["custom_network_config", "slo_config", "static_routes", "static_routes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "List of static routes.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config.static_routes

Breadcrumbs:

- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/)
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

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/): complete subsection reference.

## Next pages

- [custom_network_config.slo_config.static_routes.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/static_routes/static_routes/)
- [custom_network_config.slo_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/properties/custom_network_config/slo_config/)
- [xcsh_securemesh_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site/)
