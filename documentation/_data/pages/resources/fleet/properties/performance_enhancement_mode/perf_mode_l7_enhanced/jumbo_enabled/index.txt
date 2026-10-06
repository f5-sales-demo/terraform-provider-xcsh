---
page_title: "performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled"
subcategory: ""
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["performance enhancement mode perf mode l7 enhanced jumbo enabled"], "body_bytes": 1252, "body_sha256": "sha256:209871de48f2b54f5e8f35b3bbb7b0e12068342b59920defa10a410970fcde63", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced:jumbo_enabled", "parent_id": "xcsh-docs:resources:fleet:properties:performance_enhancement_mode:perf_mode_l7_enhanced", "path": "documentation/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-1302330021030223-0330123321303210-1322121120211302-2100021230020332-0311232102333332-2132001121311132-1133203232213212-0103021132211220", "registry_path": "docs/guides/resources--fleet--reference--group-002.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["performance_enhancement_mode", "perf_mode_l7_enhanced", "jumbo_enabled"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/jumbo_enabled/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [performance_enhancement_mode](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/)
- [performance_enhancement_mode.perf_mode_l7_enhanced](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/performance_enhancement_mode/perf_mode_l7_enhanced/)
- performance_enhancement_mode.perf_mode_l7_enhanced.jumbo_enabled

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
jumbo_enabled = {}
```

This is an empty object or choice marker. It has no direct properties.
