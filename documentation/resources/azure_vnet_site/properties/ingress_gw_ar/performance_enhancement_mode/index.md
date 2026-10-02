---
page_title: "ingress_gw_ar.performance_enhancement_mode"
subcategory: "Infrastructure"
description: "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default."
xcsh_docs: {"aliases": ["ingress gw ar performance enhancement mode"], "body_bytes": 2461, "body_sha256": "sha256:b9fbe9136b01d7e623e327bb50be155066c4ae8bcb05c48ea8b4380f84e6dcfb", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar", "path": "documentation/resources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/index.md", "product": "distributed-cloud", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0323133033111032-2020220200033011-0123102123210121-1210333102021330-3303201323321133-1302123212021203-0320101200121331-2101030123233102", "registry_path": "docs/guides/resources--azure_vnet_site--reference--group-008.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_gw_ar", "performance_enhancement_mode"], "schema_version": 1, "sections": [{"aliases": ["perf mode l3 enhanced"], "anchor": "section", "description": "L3 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced:ConflictingObjectAttributes:jumbo,no_jumbo", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:jumbo", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced:ConflictingObjectAttributes:jumbo,no_jumbo", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo", "type": "conflicts"}], "schema_path": ["ingress_gw_ar", "performance_enhancement_mode", "perf_mode_l3_enhanced"], "syntax": "block", "type": "object"}, {"aliases": ["perf mode l7 enhanced"], "anchor": "section", "description": "L7 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced:ConflictingObjectAttributes:jumbo_disabled,jumbo_enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced:ConflictingObjectAttributes:jumbo_disabled,jumbo_enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled", "type": "conflicts"}], "schema_path": ["ingress_gw_ar", "performance_enhancement_mode", "perf_mode_l7_enhanced"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar.performance_enhancement_mode

Breadcrumbs:

- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/)
- ingress_gw_ar.performance_enhancement_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("perf_mode_l3_enhanced",
    "perf_mode_l7_enhanced")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"perf_mode_l3_enhanced\",\"perf_mode_l7_enhanced\"]"
}
```

Terraform syntax:

```terraform
performance_enhancement_mode {
  # Configure direct properties listed below.
}
```

## Direct properties

- [perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/): complete subsection reference.

- [perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/): complete subsection reference.

## Next pages

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l3_enhanced/)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/)
- [ingress_gw_ar](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/properties/ingress_gw_ar/)
- [xcsh_azure_vnet_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/azure_vnet_site/)
