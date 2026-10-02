---
page_title: "custom_network_config.slo_config"
subcategory: ""
description: "Site local network configuration."
xcsh_docs: {"aliases": ["custom network config slo config"], "body_bytes": 4690, "body_sha256": "sha256:b6f98575340b1fcad1e7dd420e34a72b4669746157fc939600f302b050431190", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:dc_cluster_group", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:labels", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_dc_cluster_group", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_static_routes", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_static_v6_routes", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config", "path": "documentation/resources/voltstack_site/properties/custom_network_config/slo_config/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2132012231020221-2131132132130320-1002200210132332-0310023302113001-1011130332312013-1311323332102120-3023012223323311-3302301022121322", "registry_path": "docs/guides/resources--voltstack_site--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config:ConflictingObjectAttributes:dc_cluster_group,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:dc_cluster_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config:ConflictingObjectAttributes:dc_cluster_group,no_dc_cluster_group", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_dc_cluster_group", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config:ConflictingObjectAttributes:no_static_v6_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_static_v6_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config:ConflictingObjectAttributes:no_static_v6_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "slo_config"], "schema_version": 1, "sections": [{"aliases": ["dc cluster group"], "anchor": "section", "description": "This type establishes a direct reference from one object(the referrer) to another(the referred). Such a reference is in form of tenant/namespace/name.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:dc_cluster_group", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-custom_network_config--slo_config--dc_cluster_group--name", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.dc_cluster_group:RequiredObjectAttributes:name", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:dc_cluster_group", "type": "requires"}], "schema_path": ["custom_network_config", "slo_config", "dc_cluster_group"], "syntax": "block", "type": "object"}, {"aliases": ["labels"], "anchor": "section", "description": "Add Labels for this network, these labels can be used in firewall policy.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:labels", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["custom_network_config", "slo_config", "labels"], "syntax": "block", "type": "object"}, {"aliases": ["no dc cluster group"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_dc_cluster_group", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "slo_config", "no_dc_cluster_group"], "syntax": "attribute", "type": "object"}, {"aliases": ["no static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "slo_config", "no_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["no static v6 routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:no_static_v6_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "slo_config", "no_static_v6_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_routes:static_routes", "type": "requires"}], "schema_path": ["custom_network_config", "slo_config", "static_routes"], "syntax": "block", "type": "object"}, {"aliases": ["static v6 routes"], "anchor": "section", "description": "List of IPv6 static routes.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.slo_config.static_v6_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:slo_config:static_v6_routes:static_routes", "type": "requires"}], "schema_path": ["custom_network_config", "slo_config", "static_v6_routes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/slo_config/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Site local network configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.slo_config

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- custom_network_config.slo_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Site Local Network Configuration. Site local network configuration.

Upstream description:

Site local network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("dc_cluster_group",
    "no_dc_cluster_group"),
  validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_static_v6_routes",
    "static_v6_routes")}
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
  "x-ves-oneof-field-dc_cluster_group_choice": "[\"dc_cluster_group\",\"no_dc_cluster_group\"]",
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_static_v6_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
slo_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/dc_cluster_group/): complete subsection reference.

- [labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/labels/): complete subsection reference.

- [no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/no_dc_cluster_group/): complete subsection reference.

- [no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/no_static_routes/): complete subsection reference.

- [no_static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/no_static_v6_routes/): complete subsection reference.

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/): complete subsection reference.

- [static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/): complete subsection reference.

## Next pages

- [custom_network_config.slo_config.dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/dc_cluster_group/)
- [custom_network_config.slo_config.labels](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/labels/)
- [custom_network_config.slo_config.no_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/no_dc_cluster_group/)
- [custom_network_config.slo_config.no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/no_static_routes/)
- [custom_network_config.slo_config.no_static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/no_static_v6_routes/)
- [custom_network_config.slo_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/static_routes/)
- [custom_network_config.slo_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/slo_config/static_v6_routes/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
