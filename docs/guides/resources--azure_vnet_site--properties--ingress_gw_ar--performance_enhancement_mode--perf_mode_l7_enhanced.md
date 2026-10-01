---
page_title: "ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced"
subcategory: "Infrastructure"
description: "ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2318, "body_sha256": "sha256:f8dda395a234be96d1f914dfd7adcb1b51533bdc129c13fdacb9d8ceff0b1cec", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode:perf_mode_l7_enhanced", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw_ar:performance_enhancement_mode", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw_ar", "performance_enhancement_mode", "perf_mode_l7_enhanced"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_gw_ar/performance_enhancement_mode/perf_mode_l7_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_gw_ar](resources--azure_vnet_site--properties--ingress_gw_ar.md)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode.md)
- ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo_disabled",
    "jumbo_enabled")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo_disabled\",\"jumbo_enabled\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l7_enhanced {
  # Configure direct properties listed below.
}
```

## Direct properties

- [jumbo_disabled](resources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md): complete subsection reference.

- [jumbo_enabled](resources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md): complete subsection reference.

## Next pages

- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md)
- [ingress_gw_ar.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md)
- [ingress_gw_ar.performance_enhancement_mode](resources--azure_vnet_site--properties--ingress_gw_ar--performance_enhancement_mode.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
