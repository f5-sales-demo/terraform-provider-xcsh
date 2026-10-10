---
page_title: "xcsh_waf_latest_signatures_version"
subcategory: ""
description: "Reads latest WAF signatures version information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["waf latest signatures version"], "body_bytes": 1351, "body_sha256": "sha256:6a8fdb4da2825dfd6fef362d6aed9ae5e6ba7f96cabecbb497666499e67bc615", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:waf_latest_signatures_version:reference", "xcsh-docs:data-sources:waf_latest_signatures_version:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_latest_signatures_version:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_latest_signatures_version:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/waf_latest_signatures_version/index.md", "product": "distributed-cloud", "provider_name": "waf_latest_signatures_version", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3220021120200003-2131100022002331-1323212033111120-3012322121033220-3000100023133103-1211320023113021-0322300310221012-2100112203232030", "registry_path": "docs/data-sources/waf_latest_signatures_version.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_latest_signatures_version/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reads latest WAF signatures version information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_waf_latest_signatures_version

Breadcrumbs:

- xcsh_waf_latest_signatures_version

Reads latest WAF signatures version information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_latest_signatures_version/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_latest_signatures_version/examples/)
