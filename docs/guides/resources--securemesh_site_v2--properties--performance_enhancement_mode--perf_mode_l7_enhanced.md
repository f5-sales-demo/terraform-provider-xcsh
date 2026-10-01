---
page_title: "performance_enhancement_mode.perf_mode_l7_enhanced"
subcategory: ""
description: "performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2101, "body_sha256": "sha256:3801b9ae94be12e083c08847a7f79c782134d76ec695c855568636e7ad05ca1f", "canonical_id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled"], "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "parent_id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode", "path": "docs/guides/resources--securemesh_site_v2--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/performance_enhancement_mode/perf_mode_l7_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode.perf_mode_l7_enhanced

Breadcrumbs:

- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
- [Property reference](resources--securemesh_site_v2--reference.md)
- [performance_enhancement_mode](resources--securemesh_site_v2--properties--performance_enhancement_mode.md)
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

- [jumbo_disabled](resources--securemesh_site_v2--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md): complete subsection reference.

- [jumbo_enabled](resources--securemesh_site_v2--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](resources--securemesh_site_v2--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](resources--securemesh_site_v2--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md)
- [performance_enhancement_mode](resources--securemesh_site_v2--properties--performance_enhancement_mode.md)
- [xcsh_securemesh_site_v2](../resources/securemesh_site_v2.md)
