---
page_title: "ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced"
subcategory: "Infrastructure"
description: "ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1893, "body_sha256": "sha256:e73acc839bc23105c466ca501a07fcb2a9ffc6368515a6da3fc19dd4e01a3110", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l3_enhanced:jumbo", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:performance_enhancement_mode", "path": "docs/guides/data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw", "performance_enhancement_mode", "perf_mode_l3_enhanced"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [ingress_gw](data-sources--gcp_vpc_site--properties--ingress_gw.md)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode.md)
- ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for perf mode l3 enhanced.

Upstream description:

L3 enhanced performance mode OPTIONS.

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

## Direct properties

- [jumbo](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md): complete subsection reference.

- [no_jumbo](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md): complete subsection reference.

## Next pages

- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.jumbo](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--jumbo.md)
- [ingress_gw.performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l3_enhanced--no_jumbo.md)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
