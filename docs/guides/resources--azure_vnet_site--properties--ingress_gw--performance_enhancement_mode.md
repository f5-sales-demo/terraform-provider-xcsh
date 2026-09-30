---
page_title: "ingress_gw.performance_enhancement_mode"
subcategory: "Infrastructure"
description: "ingress_gw.performance_enhancement_mode for xcsh_azure_vnet_site."
xcsh_docs: {"aliases": [], "body_bytes": 2014, "body_sha256": "sha256:241a7f160fd807d8627e3b9eeaec6ee208934877c9e69f75cafc53f23443f1ef", "canonical_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode", "child_ids": ["xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l7_enhanced"], "collection_id": "xcsh-docs:resources:azure_vnet_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw:performance_enhancement_mode", "parent_id": "xcsh-docs:resources:azure_vnet_site:properties:ingress_gw", "path": "docs/guides/resources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode.md", "provider_name": "azure_vnet_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw", "performance_enhancement_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/azure_vnet_site/properties/ingress_gw/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw.performance_enhancement_mode for xcsh_azure_vnet_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["azure_vnet_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# ingress_gw.performance_enhancement_mode

Breadcrumbs:

- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
- [Property reference](resources--azure_vnet_site--reference.md)
- [ingress_gw](resources--azure_vnet_site--properties--ingress_gw.md)
- ingress_gw.performance_enhancement_mode

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default. Server applies
default when omitted.

Upstream description:

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

- [perf_mode_l3_enhanced](resources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced.md): complete subsection reference.

- [perf_mode_l7_enhanced](resources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced.md): complete subsection reference.

## Next pages

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](resources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced.md)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](resources--azure_vnet_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced.md)
- [ingress_gw](resources--azure_vnet_site--properties--ingress_gw.md)
- [xcsh_azure_vnet_site](../resources/azure_vnet_site.md)
