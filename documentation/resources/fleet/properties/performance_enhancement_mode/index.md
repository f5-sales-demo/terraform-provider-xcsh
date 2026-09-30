---
page_title: "performance_enhancement_mode"
subcategory: ""
description: "performance_enhancement_mode for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2014, "body_sha256": "sha256:dae91b00b9367d698f8ca54ab75fb248607eef3b4cedc0bc6c8a2b6fc97aae05", "child_ids": ["xcsh-docs:resources:fleet:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:resources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:performance_enhancement_mode", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "documentation/resources/fleet/properties/performance_enhancement_mode/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["performance_enhancement_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "performance_enhancement_mode for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# performance_enhancement_mode

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- performance_enhancement_mode

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

- [perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/perf_mode_l3_enhanced/): complete subsection reference.

- [perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/perf_mode_l3_enhanced/)
- [performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
