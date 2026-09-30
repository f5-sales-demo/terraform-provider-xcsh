---
page_title: "xcsh_network_policy_set"
subcategory: ""
description: "xcsh_network_policy_set for xcsh_network_policy_set."
xcsh_docs: {"aliases": [], "body_bytes": 1270, "body_sha256": "sha256:24f54a77847e6d331569511bfc2f203af84e9d5a6e3377ccb48f06482291fee5", "canonical_id": "xcsh-docs:data-sources:network_policy_set:fundamentals", "child_ids": ["xcsh-docs:data-sources:network_policy_set:reference", "xcsh-docs:data-sources:network_policy_set:examples"], "collection_id": "xcsh-docs:data-sources:network_policy_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_set:fundamentals", "parent_id": null, "path": "docs/data-sources/network_policy_set.md", "provider_name": "network_policy_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_policy_set for xcsh_network_policy_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_network_policy_set

Breadcrumbs:

- xcsh_network_policy_set

Manages a Network Policy Set resource in F5 Distributed Cloud for get network policy set in a given
namespace. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicySet by name
data "xcsh_network_policy_set" "example" {
  name      = "example-network-policy-set"
  namespace = "staging"
}

output "network_policy_set_id" {
  value = data.xcsh_network_policy_set.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--network_policy_set--reference.md)
- [Examples](../guides/data-sources--network_policy_set--examples.md)
