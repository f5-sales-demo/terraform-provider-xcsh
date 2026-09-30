---
page_title: "xcsh_third_party_application"
subcategory: ""
description: "xcsh_third_party_application for xcsh_third_party_application."
xcsh_docs: {"aliases": [], "body_bytes": 1404, "body_sha256": "sha256:66ed6c5cb0abedf0342ff38a7d058a19c793bea40ab388dd0a1f3f706874f3da", "child_ids": ["xcsh-docs:data-sources:third_party_application:reference", "xcsh-docs:data-sources:third_party_application:examples"], "collection_id": "xcsh-docs:data-sources:third_party_application:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:third_party_application:fundamentals", "parent_id": null, "path": "documentation/data-sources/third_party_application/index.md", "provider_name": "third_party_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/third_party_application/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_third_party_application for xcsh_third_party_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_third_party_application

Breadcrumbs:

- xcsh_third_party_application

Manages a Third Party Application resource in F5 Distributed Cloud for third party application
specification. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ThirdPartyApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ThirdPartyApplication by name
data "xcsh_third_party_application" "example" {
  name      = "example-third-party-application"
  namespace = "staging"
}

output "third_party_application_id" {
  value = data.xcsh_third_party_application.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/third_party_application/examples/)
