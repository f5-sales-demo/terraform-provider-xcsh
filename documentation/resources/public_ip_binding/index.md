---
page_title: "xcsh_public_ip_binding"
subcategory: ""
description: "Manage the regional virtual-site binding of an already allocated public IP. Creation adopts only the binding; deletion restores its captured original binding. This resource never allocates, deallocates or deletes the public IP."
xcsh_docs: {"aliases": ["public ip binding"], "body_bytes": 1703, "body_sha256": "sha256:87e5dc2b1fb751edb20c173aab7965307a4ef4ed904d2183a01396182451748f", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:public_ip_binding:reference", "xcsh-docs:resources:public_ip_binding:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:public_ip_binding:collection", "completeness": "complete", "id": "xcsh-docs:resources:public_ip_binding:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/public_ip_binding/index.md", "product": "distributed-cloud", "provider_name": "public_ip_binding", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0203231102121100-0021222023322132-1033002220322133-0210023131031121-3331122113020013-3010312110213323-0102222031221012-2102213113003031", "registry_path": "docs/resources/public_ip_binding.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/public_ip_binding/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manage the regional virtual-site binding of an already allocated public IP. Creation adopts only the binding; deletion restores its captured original binding. This resource never allocates, deallocates or deletes the public IP.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/public_ip_binding/examples/)
