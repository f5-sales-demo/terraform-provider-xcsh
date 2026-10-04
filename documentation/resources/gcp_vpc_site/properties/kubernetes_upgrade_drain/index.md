---
page_title: "kubernetes_upgrade_drain"
subcategory: "Infrastructure"
description: "Specify how worker nodes within a site will be upgraded."
xcsh_docs: {"aliases": ["kubernetes upgrade drain"], "body_bytes": 2140, "body_sha256": "sha256:3971874c6f9677c9be8565735dc758123a69b0df4e1ab25520998607383d61c2", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain", "parent_id": "xcsh-docs:resources:gcp_vpc_site:reference", "path": "documentation/resources/gcp_vpc_site/properties/kubernetes_upgrade_drain/index.md", "product": "distributed-cloud", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2101212102330031-0010112330122012-1102313232223312-0222112303201021-3012113311101013-3130313120212310-1030123002232113-0300133312323323", "registry_path": "docs/guides/resources--gcp_vpc_site--reference--group-003.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain:ConflictingObjectAttributes:disable_upgrade_drain,enable_upgrade_drain", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain:ConflictingObjectAttributes:disable_upgrade_drain,enable_upgrade_drain", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["kubernetes_upgrade_drain"], "schema_version": 1, "sections": [{"aliases": ["kubernetes upgrade drain disable upgrade drain"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:disable_upgrade_drain", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["kubernetes_upgrade_drain", "disable_upgrade_drain"], "syntax": "attribute", "type": "object"}, {"aliases": ["kubernetes upgrade drain enable upgrade drain"], "anchor": "section", "description": "Specify batch upgrade settings for worker nodes within a site.", "document_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_count", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain.enable_upgrade_drain:ConflictingObjectAttributes:drain_max_unavailable_node_count,drain_max_unavailable_node_percentage", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "type": "conflicts"}, {"anchor": "schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_max_unavailable_node_percentage", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain.enable_upgrade_drain:ConflictingObjectAttributes:drain_max_unavailable_node_count,drain_max_unavailable_node_percentage", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain.enable_upgrade_drain:ConflictingObjectAttributes:disable_vega_upgrade_mode,enable_vega_upgrade_mode", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain.enable_upgrade_drain:ConflictingObjectAttributes:disable_vega_upgrade_mode,enable_vega_upgrade_mode", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode", "type": "conflicts"}, {"anchor": "schema-kubernetes_upgrade_drain--enable_upgrade_drain--drain_node_timeout", "enforcement": "provider-schema", "group": "kubernetes_upgrade_drain.enable_upgrade_drain:RequiredObjectAttributes:drain_node_timeout", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:gcp_vpc_site:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "type": "requires"}], "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/gcp_vpc_site/properties/kubernetes_upgrade_drain/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Specify how worker nodes within a site will be upgraded.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain

Breadcrumbs:

- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- kubernetes_upgrade_drain

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Specify how worker nodes within a site will be upgraded.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_upgrade_drain",
    "enable_upgrade_drain")}
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
  "x-ves-oneof-field-kubernetes_upgrade_drain_enable_choice": "[\"disable_upgrade_drain\",\"enable_upgrade_drain\"]"
}
```

Terraform syntax:

```terraform
kubernetes_upgrade_drain {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/): complete subsection reference.

- [enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/): complete subsection reference.

## Next pages

- [kubernetes_upgrade_drain.disable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/kubernetes_upgrade_drain/disable_upgrade_drain/)
- [kubernetes_upgrade_drain.enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/kubernetes_upgrade_drain/enable_upgrade_drain/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/properties/)
- [xcsh_gcp_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/gcp_vpc_site/)
