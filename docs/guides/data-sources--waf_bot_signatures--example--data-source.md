---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_bot_signatures."
xcsh_docs: {"aliases": [], "body_bytes": 1036, "body_sha256": "sha256:7fc42ff965f652b7295ca675c702b7c7b30d2be2872d8a9bbb472741ba2c4674", "canonical_id": "xcsh-docs:data-sources:waf_bot_signatures:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:waf_bot_signatures:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0cea795e7e9cddaecc9f49e84c9c13efa8ac69b7e21b4ee8402220f9f8ec4b3f", "source_path": "examples/data-sources/xcsh_waf_bot_signatures/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_bot_signatures:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_bot_signatures:examples", "path": "docs/guides/data-sources--waf_bot_signatures--example--data-source.md", "provider_name": "waf_bot_signatures", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_bot_signatures/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_waf_bot_signatures.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_waf_bot_signatures](../data-sources/waf_bot_signatures.md)
- [Examples](data-sources--waf_bot_signatures--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_waf_bot_signatures/data-source.tf`; digest `sha256:0cea795e7e9cddaecc9f49e84c9c13efa8ac69b7e21b4ee8402220f9f8ec4b3f`.

```terraform
# WAFBotSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_bot_signatures" "example" {
}

output "waf_bot_signatures_result" {
  value = data.xcsh_waf_bot_signatures.example
}
```

## Next pages

- [Examples](data-sources--waf_bot_signatures--examples.md)
- [xcsh_waf_bot_signatures](../data-sources/waf_bot_signatures.md)
