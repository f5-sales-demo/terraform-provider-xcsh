---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_latest_signatures_version."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1083, "body_sha256": "sha256:110f9509bf8e2a9520b0dffddfa321501112f96f0ca7eef4584875835158ba63", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_latest_signatures_version:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6a87350aef0ac879fd92c4cae37ff6595f45e1a414503ebb0313bae760a73e29", "source_path": "examples/data-sources/xcsh_waf_latest_signatures_version/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_latest_signatures_version:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_latest_signatures_version:examples", "path": "documentation/data-sources/waf_latest_signatures_version/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "waf_latest_signatures_version", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3322302323021011-1001303221012110-1002132321030331-0110222012202212-3312233122102101-2222201132221212-1031100103330133-0330300331023101", "registry_path": "docs/guides/data-sources--waf_latest_signatures_version--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_latest_signatures_version/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_waf_latest_signatures_version.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
