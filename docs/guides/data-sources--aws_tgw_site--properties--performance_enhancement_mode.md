---
page_title: "performance_enhancement_mode"
subcategory: ""
description: "performance_enhancement_mode for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1371, "body_sha256": "sha256:90c8db9c6f75442cb5be640082d8c269dca62f20e01bac40e105cb2f36c7037e", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:properties:performance_enhancement_mode", "child_ids": ["xcsh-docs:data-sources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:data-sources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced"], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_tgw_site:properties:performance_enhancement_mode", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:reference", "path": "docs/guides/data-sources--aws_tgw_site--properties--performance_enhancement_mode.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["performance_enhancement_mode"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/properties/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "performance_enhancement_mode for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# performance_enhancement_mode

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
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

- [perf_mode_l3_enhanced](data-sources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md): complete subsection reference.

- [perf_mode_l7_enhanced](data-sources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l3_enhanced](data-sources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l3_enhanced.md)
- [performance_enhancement_mode.perf_mode_l7_enhanced](data-sources--aws_tgw_site--properties--performance_enhancement_mode--perf_mode_l7_enhanced.md)
- [Property reference](data-sources--aws_tgw_site--reference.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
