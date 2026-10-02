---
page_title: "ingress_egress_gw.performance_enhancement_mode"
subcategory: "Infrastructure"
description: "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default."
xcsh_docs: {"aliases": ["ingress egress gw performance enhancement mode"], "body_bytes": 2174, "body_sha256": "sha256:938be3ca794feb6b66a9fae081dfe3ac04f4fe8dcb935e9ce4b704ff45f1c055", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l7_enhanced"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw", "path": "documentation/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/index.md", "product": "distributed-cloud", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3333111201203120-1223122133331331-2131312003210200-0303001232012121-2122230310310231-0010001201223202-3313133220111030-3032010201000230", "registry_path": "docs/guides/data-sources--aws_vpc_site--reference--group-003.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["ingress_egress_gw", "performance_enhancement_mode"], "schema_version": 1, "sections": [{"aliases": ["perf mode l3 enhanced"], "anchor": "section", "description": "L3 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l3_enhanced", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "performance_enhancement_mode", "perf_mode_l3_enhanced"], "syntax": "attribute", "type": "object"}, {"aliases": ["perf mode l7 enhanced"], "anchor": "section", "description": "L7 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:data-sources:aws_vpc_site:properties:ingress_egress_gw:performance_enhancement_mode:perf_mode_l7_enhanced", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["ingress_egress_gw", "performance_enhancement_mode", "perf_mode_l7_enhanced"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# ingress_egress_gw.performance_enhancement_mode

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/)
- ingress_egress_gw.performance_enhancement_mode

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

- [perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/): complete subsection reference.

- [perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/): complete subsection reference.

## Next pages

- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l3_enhanced/)
- [ingress_egress_gw.performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/performance_enhancement_mode/perf_mode_l7_enhanced/)
- [ingress_egress_gw](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/properties/ingress_egress_gw/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
