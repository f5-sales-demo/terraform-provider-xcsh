---
page_title: "xcsh_site_signatures_update"
subcategory: ""
description: "xcsh_site_signatures_update for xcsh_site_signatures_update."
xcsh_docs: {"aliases": [], "body_bytes": 1198, "body_sha256": "sha256:806f75084e41120c7a0ea6ef134a0b57515dd8f24608ab06d5e8de6645dab886", "child_ids": ["xcsh-docs:actions:site_signatures_update:reference", "xcsh-docs:actions:site_signatures_update:examples", "xcsh-docs:actions:site_signatures_update:lifecycle"], "collection_id": "xcsh-docs:actions:site_signatures_update:collection", "completeness": "complete", "id": "xcsh-docs:actions:site_signatures_update:fundamentals", "parent_id": null, "path": "documentation/actions/site_signatures_update/index.md", "provider_name": "site_signatures_update", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_signatures_update/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_signatures_update for xcsh_site_signatures_update.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_site_signatures_update

Breadcrumbs:

- xcsh_site_signatures_update

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteSignaturesUpdate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_signatures_update" "example" {
  config {
    namespace = "example-value"
  }
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_signatures_update/lifecycle/)
