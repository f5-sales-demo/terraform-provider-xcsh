---
page_title: "upgrade_settings"
subcategory: ""
description: "Specify how a site will be upgraded."
xcsh_docs: {"aliases": ["upgrade settings"], "body_bytes": 1078, "body_sha256": "sha256:81405ee9341f9d14b3a51ded8e8fc03b0166448ecdfde76f6246242c5a133970", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:upgrade_settings", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/upgrade_settings/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2122031130011121-2120222320030031-2032103312220112-3121002222230322-2332133023022103-2321103112123110-2013202113333333-2203321000331223", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["upgrade_settings"], "schema_version": 1, "sections": [{"aliases": ["upgrade settings kubernetes upgrade drain"], "anchor": "section", "description": "Specify how worker nodes within a site will be upgraded.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "upgrade_settings.kubernetes_upgrade_drain:ConflictingObjectAttributes:disable_upgrade_drain,enable_upgrade_drain", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:disable_upgrade_drain", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "upgrade_settings.kubernetes_upgrade_drain:ConflictingObjectAttributes:disable_upgrade_drain,enable_upgrade_drain", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain", "type": "conflicts"}], "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/upgrade_settings/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Specify how a site will be upgraded.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upgrade_settings

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- upgrade_settings

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for upgrade settings.

Additional upstream details:

Specify how a site will be upgraded.

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
upgrade_settings {
  # Configure direct properties listed below.
}
```

## Direct properties

- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/): complete subsection reference.
