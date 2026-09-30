---
page_title: "performance_enhancement_mode.perf_mode_l7_enhanced"
subcategory: ""
description: "performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1722, "body_sha256": "sha256:fbb1b1823debabc124db9bc19d9e8caf900a4c57a3234f8ef9451b896622d56b", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "xcsh-docs:data-sources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:performance_enhancement_mode", "path": "docs/guides/data-sources--securemesh_site_v2--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/performance_enhancement_mode/perf_mode_l7_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# performance_enhancement_mode.perf_mode_l7_enhanced

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [performance_enhancement_mode](data-sources--securemesh_site_v2--properties--performance_enhancement_mode.md)
- performance_enhancement_mode.perf_mode_l7_enhanced

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l7 enhanced.

Upstream description:

L7 enhanced performance mode OPTIONS.

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

## Direct properties

- [jumbo_disabled](data-sources--securemesh_site_v2--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md): complete subsection reference.

- [jumbo_enabled](data-sources--securemesh_site_v2--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--securemesh_site_v2--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--securemesh_site_v2--properties--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md)
- [performance_enhancement_mode](data-sources--securemesh_site_v2--properties--performance_enhancement_mode.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
