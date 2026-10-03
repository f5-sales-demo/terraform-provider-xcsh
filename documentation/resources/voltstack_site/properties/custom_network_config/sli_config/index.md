---
page_title: "custom_network_config.sli_config"
subcategory: ""
description: "Site local inside network configuration."
xcsh_docs: {"aliases": ["custom network config sli config"], "body_bytes": 3259, "body_sha256": "sha256:829435d1c23fd5c60d163428fdbb93ae650ed7d8556e6f3dd1666f7ff2b56907", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:no_static_routes", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:no_v6_static_routes", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_routes", "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_v6_routes"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:voltstack_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config", "parent_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config", "path": "documentation/resources/voltstack_site/properties/custom_network_config/sli_config/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1102301023303321-1010303333232102-2100110003031111-0013100113133212-2221231112132310-3212231331133310-1311003003203321-3102111010332313", "registry_path": "docs/guides/resources--voltstack_site--reference--group-005.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:no_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:no_v6_static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config:ConflictingObjectAttributes:no_static_routes,static_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_routes", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config:ConflictingObjectAttributes:no_v6_static_routes,static_v6_routes", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_v6_routes", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["custom_network_config", "sli_config"], "schema_version": 1, "sections": [{"aliases": ["custom network config sli config no static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:no_static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sli_config", "no_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config sli config no v6 static routes"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:no_v6_static_routes", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["custom_network_config", "sli_config", "no_v6_static_routes"], "syntax": "attribute", "type": "object"}, {"aliases": ["custom network config sli config static routes"], "anchor": "section", "description": "List of static routes.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config.static_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_routes:static_routes", "type": "requires"}], "schema_path": ["custom_network_config", "sli_config", "static_routes"], "syntax": "block", "type": "object"}, {"aliases": ["custom network config sli config static v6 routes"], "anchor": "section", "description": "List of IPv6 static routes.", "document_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_v6_routes", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "custom_network_config.sli_config.static_v6_routes:RequiredObjectAttributes:static_routes", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:voltstack_site:properties:custom_network_config:sli_config:static_v6_routes:static_routes", "type": "requires"}], "schema_path": ["custom_network_config", "sli_config", "static_v6_routes"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/voltstack_site/properties/custom_network_config/sli_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Site local inside network configuration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# custom_network_config.sli_config

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- custom_network_config.sli_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Site local inside network configuration.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("no_static_routes",
    "static_routes"),
  validators.ConflictingObjectAttributes("no_v6_static_routes",
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
  "x-ves-oneof-field-static_route_choice": "[\"no_static_routes\",\"static_routes\"]",
  "x-ves-oneof-field-static_v6_route_choice": "[\"no_v6_static_routes\",\"static_v6_routes\"]"
}
```

Terraform syntax:

```terraform
sli_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/sli_config/no_static_routes/): complete subsection reference.

- [no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/sli_config/no_v6_static_routes/): complete subsection reference.

- [static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/sli_config/static_routes/): complete subsection reference.

- [static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/sli_config/static_v6_routes/): complete subsection reference.

## Next pages

- [custom_network_config.sli_config.no_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/sli_config/no_static_routes/)
- [custom_network_config.sli_config.no_v6_static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/sli_config/no_v6_static_routes/)
- [custom_network_config.sli_config.static_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/sli_config/static_routes/)
- [custom_network_config.sli_config.static_v6_routes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/sli_config/static_v6_routes/)
- [custom_network_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/properties/custom_network_config/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/voltstack_site/)
