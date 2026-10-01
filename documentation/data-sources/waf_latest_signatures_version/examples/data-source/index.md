---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_latest_signatures_version."
xcsh_docs: {"aliases": [], "body_bytes": 1362, "body_sha256": "sha256:1deb6a8f0ea30c15bd94fc4d6d99a497994a049e19f7575b767aeb559bddce71", "child_ids": [], "collection_id": "xcsh-docs:data-sources:waf_latest_signatures_version:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6a87350aef0ac879fd92c4cae37ff6595f45e1a414503ebb0313bae760a73e29", "source_path": "examples/data-sources/xcsh_waf_latest_signatures_version/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_latest_signatures_version:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_latest_signatures_version:examples", "path": "documentation/data-sources/waf_latest_signatures_version/examples/data-source/index.md", "provider_name": "waf_latest_signatures_version", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_latest_signatures_version/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_waf_latest_signatures_version.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_waf_latest_signatures_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_latest_signatures_version/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_latest_signatures_version/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_latest_signatures_version/data-source.tf`; digest `sha256:6a87350aef0ac879fd92c4cae37ff6595f45e1a414503ebb0313bae760a73e29`.

```terraform
# WAFLatestSignaturesVersion DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_latest_signatures_version" "example" {
}

output "waf_latest_signatures_version_result" {
  value = data.xcsh_waf_latest_signatures_version.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_latest_signatures_version/examples/)
- [xcsh_waf_latest_signatures_version](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_latest_signatures_version/)
