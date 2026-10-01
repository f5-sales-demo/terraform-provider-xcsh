---
page_title: "performance_enhancement_mode"
subcategory: ""
description: "performance_enhancement_mode for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1705, "body_sha256": "sha256:c9d6a9796149b5d8bf0712029cf3d2dd907f6ca7a7ccadb26ed052d60307c40a", "canonical_id": "xcsh-docs:resources:fleet:properties:performance_enhancement_mode", "child_ids": ["xcsh-docs:resources:fleet:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:resources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:performance_enhancement_mode", "parent_id": "xcsh-docs:resources:fleet:reference", "path": "docs/guides/resources--fleet--properties--performance_enhancement_mode.md", "provider_name": "fleet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["performance_enhancement_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "performance_enhancement_mode for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode

Breadcrumbs:

- [xcsh_fleet](../resources/fleet.md)
- [Property reference](resources--fleet--reference.md)
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

- [perf_mode_l3_enhanced](resources--fleet--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md): complete subsection reference.

- [perf_mode_l7_enhanced](resources--fleet--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l3_enhanced](resources--fleet--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md)
- [performance_enhancement_mode.perf_mode_l7_enhanced](resources--fleet--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md)
- [Property reference](resources--fleet--reference.md)
- [xcsh_fleet](../resources/fleet.md)
