---
page_title: "xcsh_service_policy_set"
subcategory: ""
description: "xcsh_service_policy_set for xcsh_service_policy_set."
xcsh_docs: {"aliases": [], "body_bytes": 1401, "body_sha256": "sha256:742ef80029e6a9e954ac9c69f87a015575a8cb72d87b7e3873184b638c971b75", "child_ids": ["xcsh-docs:data-sources:service_policy_set:reference", "xcsh-docs:data-sources:service_policy_set:examples"], "collection_id": "xcsh-docs:data-sources:service_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:service_policy_set:fundamentals", "parent_id": null, "path": "documentation/data-sources/service_policy_set/index.md", "provider_name": "service_policy_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_service_policy_set for xcsh_service_policy_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_service_policy_set

Breadcrumbs:

- xcsh_service_policy_set

Manages a Service Policy Set resource in F5 Distributed Cloud for get service\_policy\_set reads a
given object from storage backend for metadata.namespace. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicySet by name
data "xcsh_service_policy_set" "example" {
  name      = "example-service-policy-set"
  namespace = "staging"
}

output "service_policy_set_id" {
  value = data.xcsh_service_policy_set.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/examples/)
