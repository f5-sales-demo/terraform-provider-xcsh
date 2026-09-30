---
page_title: "xcsh_crl"
subcategory: ""
description: "xcsh_crl for xcsh_crl."
xcsh_docs: {"aliases": [], "body_bytes": 1536, "body_sha256": "sha256:42bf5de7e7ea53b5fed6d490574d0275c7a5520f8a5c354e95f5efc0894e533f", "child_ids": ["xcsh-docs:resources:crl:reference", "xcsh-docs:resources:crl:examples", "xcsh-docs:resources:crl:import", "xcsh-docs:resources:crl:timeouts"], "collection_id": "xcsh-docs:resources:crl:collection", "completeness": "complete", "id": "xcsh-docs:resources:crl:fundamentals", "parent_id": null, "path": "documentation/resources/crl/index.md", "provider_name": "crl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/crl/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_crl for xcsh_crl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["crlCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_crl

Breadcrumbs:

- xcsh_crl

Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CRL Resource Example
# Manages a CRL resource in F5 Distributed Cloud for api to create crl object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CRL configuration
resource "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"

  refresh_interval = 6
  server_address   = "example-value"
  server_port      = 1
  timeout          = 1
}
```

## Root configuration

Required root properties: `name`, `namespace`, `refresh_interval`, `server_address`, `server_port`, `timeout`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/crl/lifecycle/timeouts/)
