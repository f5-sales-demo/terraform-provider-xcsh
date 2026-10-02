---
page_title: "xcsh_ip_prefix_set"
subcategory: ""
description: "Manages ip_prefix_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["ip prefix set"], "body_bytes": 1370, "body_sha256": "sha256:b2e81121c30c4e595d7ad3d8b51d60bc6bd45ce9df68bef695525649320775ac", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:ip_prefix_set:reference", "xcsh-docs:data-sources:ip_prefix_set:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ip_prefix_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ip_prefix_set:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/ip_prefix_set/index.md", "product": "distributed-cloud", "provider_name": "ip_prefix_set", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2223223220333133-0031221032103123-3133100020012123-1131313331133102-3100111213100301-1210002313103313-0210003100132003-2132031023323130", "registry_path": "docs/data-sources/ip_prefix_set.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ip_prefix_set/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages ip_prefix_set creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["ip_prefix_setCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ip_prefix_set

Breadcrumbs:

- xcsh_ip_prefix_set

Manages ip\_prefix\_set creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# IPPrefixSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing IPPrefixSet by name
data "xcsh_ip_prefix_set" "example" {
  name      = "example-ip-prefix-set"
  namespace = "staging"
}

output "ip_prefix_set_id" {
  value = data.xcsh_ip_prefix_set.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ip_prefix_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ip_prefix_set/examples/)
