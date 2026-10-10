---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_bot_signatures."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 996, "body_sha256": "sha256:cd3457540494b5a625319c4b0ba17d7bc090760a6deafcde9375548018366ddf", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_bot_signatures:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0cea795e7e9cddaecc9f49e84c9c13efa8ac69b7e21b4ee8402220f9f8ec4b3f", "source_path": "examples/data-sources/xcsh_waf_bot_signatures/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_bot_signatures:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_bot_signatures:examples", "path": "documentation/data-sources/waf_bot_signatures/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "waf_bot_signatures", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1203312111202322-1332133333021203-1220133213132101-3020210131322101-2112213132033113-3211312311133100-2013320301000210-0222020121330322", "registry_path": "docs/guides/data-sources--waf_bot_signatures--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_bot_signatures/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_waf_bot_signatures.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_waf_bot_signatures](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/examples/)
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
