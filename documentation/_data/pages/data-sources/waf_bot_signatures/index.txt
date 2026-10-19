---
page_title: "xcsh_waf_bot_signatures"
subcategory: ""
description: "Bot detection and defense configuration."
xcsh_docs: {"aliases": ["waf bot signatures"], "body_bytes": 1230, "body_sha256": "sha256:238f66b8f186a800a08e00ce16cedc74de347a7bffd6872c42fdde67b14aa4ec", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:waf_bot_signatures:reference", "xcsh-docs:data-sources:waf_bot_signatures:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:waf_bot_signatures:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:waf_bot_signatures:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/waf_bot_signatures/index.md", "product": "distributed-cloud", "provider_name": "waf_bot_signatures", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-1232013323203331-2030333000300300-3031302123320310-3302000013332010-1200120322302012-2021223130131323-2112213021213111-3103303123120210", "registry_path": "docs/data-sources/waf_bot_signatures.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/waf_bot_signatures/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Bot detection and defense configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": [], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_waf_bot_signatures

Breadcrumbs:

- xcsh_waf_bot_signatures

Bot detection and defense configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/waf_bot_signatures/examples/)
