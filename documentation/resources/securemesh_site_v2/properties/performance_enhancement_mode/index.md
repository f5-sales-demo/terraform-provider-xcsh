---
page_title: "performance_enhancement_mode"
subcategory: ""
description: "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default."
xcsh_docs: {"aliases": ["performance enhancement mode"], "body_bytes": 2243, "body_sha256": "sha256:a754bf6b271d6ed129fe4a8274f6583cb5eb3ccebe13905fc518b1993e526c83", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": ["xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode", "parent_id": "xcsh-docs:resources:securemesh_site_v2:reference", "path": "documentation/resources/securemesh_site_v2/properties/performance_enhancement_mode/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2130013131033231-3232223210311301-1121020200130212-1011012213303102-1120012133331003-0202331032111120-2302012123333131-0300312001101100", "registry_path": "docs/guides/resources--securemesh_site_v2--reference--group-017.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode:ConflictingObjectAttributes:perf_mode_l3_enhanced,perf_mode_l7_enhanced", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["performance_enhancement_mode"], "schema_version": 1, "sections": [{"aliases": ["performance enhancement mode perf mode l3 enhanced"], "anchor": "section", "description": "L3 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l3_enhanced", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode.perf_mode_l3_enhanced:ConflictingObjectAttributes:jumbo,no_jumbo", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l3_enhanced:jumbo", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode.perf_mode_l3_enhanced:ConflictingObjectAttributes:jumbo,no_jumbo", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l3_enhanced:no_jumbo", "type": "conflicts"}], "schema_path": ["performance_enhancement_mode", "perf_mode_l3_enhanced"], "syntax": "block", "type": "object"}, {"aliases": ["performance enhancement mode perf mode l7 enhanced"], "anchor": "section", "description": "L7 enhanced performance mode OPTIONS.", "document_id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode.perf_mode_l7_enhanced:ConflictingObjectAttributes:jumbo_disabled,jumbo_enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_disabled", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "performance_enhancement_mode.perf_mode_l7_enhanced:ConflictingObjectAttributes:jumbo_disabled,jumbo_enabled", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:securemesh_site_v2:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled", "type": "conflicts"}], "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/securemesh_site_v2/properties/performance_enhancement_mode/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Optimize the site for L3 or L7 traffic processing. L7 optimized is the default.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
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

- [perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/performance_enhancement_mode/perf_mode_l3_enhanced/): complete subsection reference.

- [perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/performance_enhancement_mode/perf_mode_l7_enhanced/): complete subsection reference.

## Next pages

- [performance_enhancement_mode.perf_mode_l3_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/performance_enhancement_mode/perf_mode_l3_enhanced/)
- [performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/performance_enhancement_mode/perf_mode_l7_enhanced/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/properties/)
- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/securemesh_site_v2/)
