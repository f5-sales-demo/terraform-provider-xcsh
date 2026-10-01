---
page_title: "performance_enhancement_mode.perf_mode_l7_enhanced"
subcategory: ""
description: "performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 2411, "body_sha256": "sha256:b6c98b061248caad4d83b9deb54d35ffa01a51c953c4db4c5bfaad7da022fa7c", "child_ids": ["xcsh-docs:resources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "xcsh-docs:resources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "parent_id": "xcsh-docs:resources:fleet:properties:performance_enhancement_mode", "path": "documentation/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode.perf_mode_l7_enhanced

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/)
- performance_enhancement_mode.perf_mode_l7_enhanced

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

- [jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/): complete subsection reference.

- [jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
