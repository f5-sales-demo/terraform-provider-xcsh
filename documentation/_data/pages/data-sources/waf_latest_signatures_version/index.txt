---
page_title: "xcsh_waf_latest_signatures_version"
subcategory: ""
description: "Resource retrieval operation."
xcsh_docs: {"aliases": ["waf latest signatures version"], "body_bytes": 1293, "body_sha256": "sha256:35513843912a5b5b847c6a187222277e3d18d9cbdadf8c45b1db2c7b5f6d0130", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:waf_latest_signatures_version:reference", "xcsh-docs:data-sources:waf_latest_signatures_version:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_latest_signatures_version:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_latest_signatures_version:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/waf_latest_signatures_version/index.md", "product": "distributed-cloud", "provider_name": "waf_latest_signatures_version", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3220021120200003-2131100022002331-1323212033111120-3012322121033220-3000100023133103-1211320023113021-0322300310221012-2100112203232030", "registry_path": "docs/data-sources/waf_latest_signatures_version.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_latest_signatures_version/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource retrieval operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
