---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_waf_latest_signatures_version."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1362, "body_sha256": "sha256:1deb6a8f0ea30c15bd94fc4d6d99a497994a049e19f7575b767aeb559bddce71", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_latest_signatures_version:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6a87350aef0ac879fd92c4cae37ff6595f45e1a414503ebb0313bae760a73e29", "source_path": "examples/data-sources/xcsh_waf_latest_signatures_version/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:waf_latest_signatures_version:example:data-source", "parent_id": "xcsh-docs:data-sources:waf_latest_signatures_version:examples", "path": "documentation/data-sources/waf_latest_signatures_version/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "waf_latest_signatures_version", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-3322302323021011-1001303221012110-1002132321030331-0110222012202212-3312233122102101-2222201132221212-1031100103330133-0330300331023101", "registry_path": "docs/guides/data-sources--waf_latest_signatures_version--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_latest_signatures_version/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_waf_latest_signatures_version.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
