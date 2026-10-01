---
page_title: "xcsh_waf_latest_signatures_version"
subcategory: ""
description: "xcsh_waf_latest_signatures_version for xcsh_waf_latest_signatures_version."
xcsh_docs: {"aliases": [], "body_bytes": 1293, "body_sha256": "sha256:35513843912a5b5b847c6a187222277e3d18d9cbdadf8c45b1db2c7b5f6d0130", "child_ids": ["xcsh-docs:data-sources:waf_latest_signatures_version:reference", "xcsh-docs:data-sources:waf_latest_signatures_version:examples"], "collection_id": "xcsh-docs:data-sources:waf_latest_signatures_version:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_latest_signatures_version:fundamentals", "parent_id": null, "path": "documentation/data-sources/waf_latest_signatures_version/index.md", "provider_name": "waf_latest_signatures_version", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_latest_signatures_version/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_waf_latest_signatures_version for xcsh_waf_latest_signatures_version.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_waf_latest_signatures_version

Breadcrumbs:

- xcsh_waf_latest_signatures_version

Resource retrieval operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_latest_signatures_version/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_latest_signatures_version/examples/)
