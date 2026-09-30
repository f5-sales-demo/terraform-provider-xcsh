---
page_title: "ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced"
subcategory: "Infrastructure"
description: "ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2191, "body_sha256": "sha256:289c3a3aba7c42fc64a0929506d005841383f4c7b1743a24800898a7424968b8", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l3_enhanced:jumbo", "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_egress_gw:performance_enhancement_mode", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_egress_gw", "performance_enhancement_mode", "perf_mode_l3_enhanced"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_egress_gw](resources--azure_vnet_site--properties--ingress_egress_gw.md)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode.md)
- ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("jumbo",
    "no_jumbo")}
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
  "x-ves-oneof-field-perf_mode_choice": "[\"jumbo\",\"no_jumbo\"]"
}
```

Terraform syntax:

```terraform
perf_mode_l3_enhanced {
  # Configure direct properties listed below.
}
```

## Direct properties

- [jumbo](resources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md): complete subsection reference.

- [no_jumbo](resources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md): complete subsection reference.

## Next pages

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](resources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](resources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md)
- [ingress_egress_gw.performance_enhancement_mode](resources--azure_vnet_site--properties--ingress_egress_gw--performance_enhancement_mode.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
