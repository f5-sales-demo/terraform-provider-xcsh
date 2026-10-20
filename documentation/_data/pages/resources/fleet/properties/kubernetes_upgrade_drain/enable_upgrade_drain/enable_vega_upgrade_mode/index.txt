---
page_title: "kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["kubernetes upgrade drain enable upgrade drain enable vega upgrade mode"], "body_bytes": 1292, "body_sha256": "sha256:8b39c9fc62137bd5dfbf54d6255c51c242d094007eb193102c8ef34dc799f2be", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:kubernetes_upgrade_drain:enable_upgrade_drain:enable_vega_upgrade_mode", "parent_id": "xcsh-docs:resources:fleet:properties:kubernetes_upgrade_drain:enable_upgrade_drain", "path": "documentation/resources/fleet/properties/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3200231012222202-2330200310232322-0020012321110323-1102302030333121-2003110201110220-1131310233010112-0020202112123212-0302331102210222", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["kubernetes_upgrade_drain", "enable_upgrade_drain", "enable_vega_upgrade_mode"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/kubernetes_upgrade_drain/enable_upgrade_drain/enable_vega_upgrade_mode/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["fleetCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/kubernetes_upgrade_drain/)
- [kubernetes_upgrade_drain.enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/kubernetes_upgrade_drain/enable_upgrade_drain/)
- kubernetes_upgrade_drain.enable_upgrade_drain.enable_vega_upgrade_mode

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for enable vega upgrade mode.

Additional upstream details:

This can be used for messages where no values are needed.

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
enable_vega_upgrade_mode = {}
```

This is an empty object or choice marker. It has no direct properties.
