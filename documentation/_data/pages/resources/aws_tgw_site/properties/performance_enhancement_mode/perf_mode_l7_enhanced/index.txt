---
page_title: "performance_enhancement_mode.perf_mode_l7_enhanced"
subcategory: ""
description: "L7 enhanced performance mode OPTIONS."
xcsh_docs: {"aliases": ["performance enhancement mode perf mode l7 enhanced"], "body_bytes": 2488, "body_sha256": "sha256:65512fddfe7fce14ee9f52b4e976858827a414e9a4d2e5a984b9d3f683376d4a", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:aws_tgw_site:collection", "completeness": "complete", "id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "parent_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode", "path": "documentation/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3203003101111012-0201100211313210-2100200111130231-0131130201313122-0313010331200103-2030102310003103-0013100312103033-2130200133331330", "registry_path": "docs/guides/resources--aws_tgw_site--reference--group-002.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode.perf_mode_l7_enhanced:ConflictingObjectAttributes:jumbo_disabled,jumbo_enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode.perf_mode_l7_enhanced:ConflictingObjectAttributes:jumbo_disabled,jumbo_enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced"], "schema_version": 1, "sections": [{"aliases": ["jumbo disabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced", "jumbo_disabled"], "syntax": "attribute", "type": "object"}, {"aliases": ["jumbo enabled"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:aws_tgw_site:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced", "jumbo_enabled"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "L7 enhanced performance mode OPTIONS.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode.perf_mode_l7_enhanced

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/)
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

- [jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/): complete subsection reference.

- [jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_disabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_disabled/)
- [performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/properties/performance_enhancement_mode/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/aws_tgw_site/)
