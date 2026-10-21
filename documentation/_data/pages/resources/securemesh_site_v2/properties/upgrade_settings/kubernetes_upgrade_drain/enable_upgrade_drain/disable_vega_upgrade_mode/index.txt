---
page_title: "upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["upgrade settings kubernetes upgrade drain enable upgrade drain disable vega upgrade mode"], "body_bytes": 1599, "body_sha256": "sha256:23152a4a114159795a4927059e198cf62b27761b2f53d0d81b6201b3e202699a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain:disable_vega_upgrade_mode", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:upgrade_settings:kubernetes_upgrade_drain:enable_upgrade_drain", "path": "documentation/resources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3112020123021302-3203002223103220-2301221002322203-3123211132101102-1030100001330230-2033323301101002-3032133031210311-0220323010030202", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["upgrade_settings", "kubernetes_upgrade_drain", "enable_upgrade_drain", "disable_vega_upgrade_mode"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/disable_vega_upgrade_mode/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [upgrade_settings](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/upgrade_settings/)
- [upgrade_settings.kubernetes_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/)
- [upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/upgrade_settings/kubernetes_upgrade_drain/enable_upgrade_drain/)
- upgrade_settings.kubernetes_upgrade_drain.enable_upgrade_drain.disable_vega_upgrade_mode

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for disable vega upgrade mode.

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
disable_vega_upgrade_mode = {}
```

This is an empty object or choice marker. It has no direct properties.
