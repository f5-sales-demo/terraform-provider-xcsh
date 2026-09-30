---
page_title: "xcsh_addon_service"
subcategory: ""
description: "xcsh_addon_service for xcsh_addon_service."
xcsh_docs: {"aliases": [], "body_bytes": 1674, "body_sha256": "sha256:ff52e252fbed0533ad33f9fb8eb3297d14c58cf46b5bf0577771a06b50927e84", "child_ids": ["xcsh-docs:data-sources:addon_service:reference", "xcsh-docs:data-sources:addon_service:examples"], "collection_id": "xcsh-docs:data-sources:addon_service:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:addon_service:fundamentals", "parent_id": null, "path": "documentation/data-sources/addon_service/index.md", "provider_name": "addon_service", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/addon_service/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_addon_service for xcsh_addon_service.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_addon_service

Breadcrumbs:

- xcsh_addon_service

Retrieves information about an F5 Distributed Cloud Addon Service.

Addon services are system-managed resources that provide additional functionality such as Bot
Defense, Client Side Defense, and other security features. This data source allows you to query
addon service details including tier requirements and activation type.

~&gt; \*\*Note:\*\* Addon services cannot be created or modified via Terraform. To activate or
subscribe to an addon service, please use the F5 Distributed Cloud Console or contact your account
team.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AddonService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AddonService by name
data "xcsh_addon_service" "example" {
  name      = "example-addon-service"
  namespace = "staging"
}

output "addon_service_id" {
  value = data.xcsh_addon_service.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/addon_service/examples/)
