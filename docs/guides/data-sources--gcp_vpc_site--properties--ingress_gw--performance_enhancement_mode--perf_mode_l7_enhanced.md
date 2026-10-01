---
page_title: "ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced"
subcategory: "Infrastructure"
description: "ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_gcp_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1963, "body_sha256": "sha256:e69d28737330eea7fbd77308562e41e4438b6157a2836d501e2e1285e934342e", "canonical_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l7_enhanced", "child_ids": ["xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled"], "collection_id": "xcsh-docs:data-sources:gcp_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:performance_enhancement_mode:perf_mode_l7_enhanced", "parent_id": "xcsh-docs:data-sources:gcp_vpc_site:properties:ingress_gw:performance_enhancement_mode", "path": "docs/guides/data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced.md", "provider_name": "gcp_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["ingress_gw", "performance_enhancement_mode", "perf_mode_l7_enhanced"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/gcp_vpc_site/properties/ingress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced for xcsh_gcp_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["gcp_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

Breadcrumbs:

- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
- [Property reference](data-sources--gcp_vpc_site--reference.md)
- [ingress_gw](data-sources--gcp_vpc_site--properties--ingress_gw.md)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode.md)
- ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced

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

- [jumbo_disabled](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md): complete subsection reference.

- [jumbo_enabled](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md): complete subsection reference.

## Next pages

- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_disabled.md)
- [ingress_gw.performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode--perf_mode_l7_enhanced--jumbo_enabled.md)
- [ingress_gw.performance_enhancement_mode](data-sources--gcp_vpc_site--properties--ingress_gw--performance_enhancement_mode.md)
- [xcsh_gcp_vpc_site](../data-sources/gcp_vpc_site.md)
