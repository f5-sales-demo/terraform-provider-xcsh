---
page_title: "performance_enhancement_mode"
subcategory: ""
description: "performance_enhancement_mode for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1500, "body_sha256": "sha256:856587a09eb93294eb17c58aceb95dbe12b08041e6100f4d857260252bc2d996", "canonical_id": "xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode", "child_ids": ["xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced"], "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:properties:performance_enhancement_mode", "parent_id": "xcsh-docs:data-sources:securemesh_site:reference", "path": "docs/guides/data-sources--securemesh_site--properties--performance_enhancement_mode.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["performance_enhancement_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/properties/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "performance_enhancement_mode for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
- [Property reference](data-sources--securemesh_site--reference.md)
- performance_enhancement_mode

<a id="section"></a>

Type: `"single"`. Computed.

Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.

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

## Direct properties

- [perf_mode_l3_enhanced](data-sources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--securemesh_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md)
- [Property reference](data-sources--securemesh_site--reference.md)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
