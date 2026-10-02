---
page_title: "xcsh_bigip_virtual_server"
subcategory: ""
description: "Manages a BIG-IP Virtual Server resource in F5 Distributed Cloud for big-ip virtual server specification. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["bigip virtual server"], "body_bytes": 1469, "body_sha256": "sha256:6fa171737da052034a593f3ddee1542bb7a3927b4a218ac08c1200b87ceef5a6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bigip_virtual_server:reference", "xcsh-docs:data-sources:bigip_virtual_server:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bigip_virtual_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bigip_virtual_server:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bigip_virtual_server/index.md", "product": "distributed-cloud", "provider_name": "bigip_virtual_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1131322003302203-3133232130331303-3111132332133231-2012331201133203-0022310012333001-0003330233211211-0333312201021000-3230332012332020", "registry_path": "docs/data-sources/bigip_virtual_server.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bigip_virtual_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a BIG-IP Virtual Server resource in F5 Distributed Cloud for big-ip virtual server specification. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bigip_virtual_server

Breadcrumbs:

- xcsh_bigip_virtual_server

Manages a BIG-IP Virtual Server resource in F5 Distributed Cloud for big-ip virtual server
specification. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BigIPVirtualServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BigIPVirtualServer by name
data "xcsh_bigip_virtual_server" "example" {
  name      = "example-bigip-virtual-server"
  namespace = "staging"
}

output "bigip_virtual_server_id" {
  value = data.xcsh_bigip_virtual_server.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bigip_virtual_server/examples/)
