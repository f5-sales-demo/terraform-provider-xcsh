---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_attack_signatures."
xcsh_docs: {"aliases": [], "body_bytes": 1176, "body_sha256": "sha256:11404abab402e02c82593252ec5067aa3a7af3640cd564e4a4df7c66c3a210af", "child_ids": [], "collection_id": "xcsh-docs:data-sources:waf_attack_signatures:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:5857e356bd3dea3a26e0c6c17fd28eb9757681e0e4bfa5915066871a8e38e467", "source_path": "examples/data-sources/xcsh_waf_attack_signatures/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_attack_signatures:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_attack_signatures:examples", "path": "documentation/data-sources/waf_attack_signatures/examples/data-source/index.md", "provider_name": "waf_attack_signatures", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_attack_signatures/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_waf_attack_signatures.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_waf_attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_attack_signatures/data-source.tf`; digest `sha256:5857e356bd3dea3a26e0c6c17fd28eb9757681e0e4bfa5915066871a8e38e467`.

```terraform
# WAFAttackSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_attack_signatures" "example" {
}

output "waf_attack_signatures_result" {
  value = data.xcsh_waf_attack_signatures.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/examples/)
- [xcsh_waf_attack_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_attack_signatures/)
