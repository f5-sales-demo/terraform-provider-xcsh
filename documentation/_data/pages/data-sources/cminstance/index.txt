---
page_title: "xcsh_cminstance"
subcategory: ""
description: "Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["cminstance"], "body_bytes": 1332, "body_sha256": "sha256:ddb5860528c3d335029eee7efbf619c7ed43c0ae93c6bbca9988bf2878db4ad8", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:cminstance:reference", "xcsh-docs:data-sources:cminstance:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cminstance:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/cminstance/index.md", "product": "distributed-cloud", "provider_name": "cminstance", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0002331230202011-3212203300130123-1211123012101021-2203230131300312-2133023023220023-3111001023122322-0012003002331212-3011011010021020", "registry_path": "docs/data-sources/cminstance.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cminstance/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cminstanceCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cminstance

Breadcrumbs:

- xcsh_cminstance

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cminstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cminstance by name
data "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"
}

output "cminstance_id" {
  value = data.xcsh_cminstance.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/examples/)
