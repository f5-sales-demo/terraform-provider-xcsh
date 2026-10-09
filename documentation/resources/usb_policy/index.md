---
page_title: "xcsh_usb_policy"
subcategory: ""
description: "Manages new USB policy object in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["usb policy"], "body_bytes": 1486, "body_sha256": "sha256:7cbd3767b58ae88ae657e3d5d7100c953201db3404ebf859c7169a940051f18d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:usb_policy:reference", "xcsh-docs:resources:usb_policy:examples", "xcsh-docs:resources:usb_policy:import", "xcsh-docs:resources:usb_policy:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:usb_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:usb_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/usb_policy/index.md", "product": "distributed-cloud", "provider_name": "usb_policy", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3313033212003300-1002332031103020-1323001313000200-3003230033322232-3111031201110211-0323302100321032-3212002010220013-3201110212312232", "registry_path": "docs/resources/usb_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/usb_policy/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Manages new USB policy object in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["usb_policyCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_usb_policy

Breadcrumbs:

- xcsh_usb_policy

Manages new USB policy object in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UsbPolicy Resource Example
# Manages new USB policy object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UsbPolicy configuration
resource "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/lifecycle/timeouts/)
