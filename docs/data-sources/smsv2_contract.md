---
page_title: "xcsh_smsv2_contract"
subcategory: ""
description: "xcsh_smsv2_contract for xcsh_smsv2_contract."
xcsh_docs: {"aliases": [], "body_bytes": 2030, "body_sha256": "sha256:ec92e904974ccbc86c33f43d93eb437a180907102dbea2c1d277ee1edd4aa136", "canonical_id": "xcsh-docs:data-sources:smsv2_contract:fundamentals", "child_ids": ["xcsh-docs:data-sources:smsv2_contract:reference", "xcsh-docs:data-sources:smsv2_contract:examples"], "collection_id": "xcsh-docs:data-sources:smsv2_contract:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:smsv2_contract:fundamentals", "parent_id": null, "path": "docs/data-sources/smsv2_contract.md", "provider_name": "smsv2_contract", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/smsv2_contract/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_smsv2_contract for xcsh_smsv2_contract.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_smsv2_contract

Breadcrumbs:

- xcsh_smsv2_contract

Publishes the immutable clean-break SMSv2 AWS, Azure, and KVM capability contracts compiled into
this provider release.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Read the immutable clean-break SMSv2 contract compiled into the provider.
# Required capabilities are checked during planning before any F5 API request.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 6.0.0"
    }
  }
}

data "xcsh_smsv2_contract" "current" {
  required_capabilities = ["runtime_status"]
}

output "smsv2_contract" {
  value = {
    id                               = data.xcsh_smsv2_contract.current.contract_id
    version                          = data.xcsh_smsv2_contract.current.contract_version
    api_release                      = data.xcsh_smsv2_contract.current.api_release_tag
    telemetry_schema_id              = data.xcsh_smsv2_contract.current.telemetry_schema_id
    capabilities                     = data.xcsh_smsv2_contract.current.capabilities
    f5xc_authorities                 = data.xcsh_smsv2_contract.current.f5xc_authorities
    aws_authorities                  = data.xcsh_smsv2_contract.current.aws_authorities
    azure_route_server_ebgp_multihop = data.xcsh_smsv2_contract.current.azure_route_server_ebgp_multihop
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--smsv2_contract--reference.md)
- [Examples](../guides/data-sources--smsv2_contract--examples.md)
