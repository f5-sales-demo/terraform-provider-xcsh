---
page_title: "performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["performance enhancement mode perf mode l3 enhanced no jumbo"], "body_bytes": 1272, "body_sha256": "sha256:9c9671cf50edca854a1e5567cf8f3e384ff317af3ed9ed5328753ff264488af7", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "path": "documentation/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-1010332012232311-3200012011213203-0100120310321002-0302223012302120-0330213001013212-3001211012002031-2230202000313123-0230321103313311", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["performance_enhancement_mode", "perf_mode_l3_enhanced", "no_jumbo"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/no_jumbo/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/)
- [performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l3_enhanced/)
- performance_enhancement_mode.perf_mode_l3_enhanced.no_jumbo

<a id="section"></a>

Type: `["object", {}]`. Optional.

Enable this option

Additional upstream details:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
no_jumbo = {}
```

This is an empty object or choice marker. It has no direct properties.
