---
page_title: "xcsh_public_ip_binding"
subcategory: ""
description: "Manage the regional virtual-site binding of an already allocated public IP. Creation adopts only the binding; deletion restores its captured original binding. This resource never allocates, deallocates or deletes the public IP."
xcsh_docs: {"aliases": ["public ip binding"], "body_bytes": 1690, "body_sha256": "sha256:4891a3f4105add408169a28ba7a109046ecafbf3409adb3f33538af08fbd9d85", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:public_ip_binding:reference", "xcsh-docs:resources:public_ip_binding:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:public_ip_binding:collection", "completeness": "complete", "id": "xcsh-docs:resources:public_ip_binding:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/public_ip_binding/index.md", "product": "distributed-cloud", "provider_name": "public_ip_binding", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0203231102121100-0021222023322132-1033002220322133-0210023131031121-3331122113020013-3010312110213323-0102222031221012-2102213113003031", "registry_path": "docs/resources/public_ip_binding.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/public_ip_binding/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Manage the regional virtual-site binding of an already allocated public IP. Creation adopts only the binding; deletion restores its captured original binding. This resource never allocates, deallocates or deletes the public IP.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": [], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_public_ip_binding

Breadcrumbs:

- xcsh_public_ip_binding

Manage the regional virtual-site binding of an already allocated public IP. Creation adopts only the
binding; deletion restores its captured original binding. This resource never allocates, deallocates
or deletes the public IP.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# PublicIPBinding Resource Example
# Manage the regional virtual-site binding of an already allocated public IP.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PublicIPBinding configuration
resource "xcsh_public_ip_binding" "example" {
  name      = "example-public-ip-binding"
  namespace = "staging"

  expected_ip            = "example-value"
  virtual_site           = "example-value"
  virtual_site_namespace = "example-value"
}
```

## Root configuration

Required root properties: `expected_ip`, `name`, `namespace`, `virtual_site`, `virtual_site_namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/examples/)
