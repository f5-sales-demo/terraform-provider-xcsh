---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_usb_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1253, "body_sha256": "sha256:e1021f32143078567a2bdb9822b38d685cdebcda05d496e67c737a1fe87937fb", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:usb_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:40aac30ea686e00cce92cecb2aee01cb69e14aef947ae725b686bf0b6e9af7b0", "source_path": "examples/data-sources/xcsh_usb_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:usb_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:usb_policy:examples", "path": "documentation/data-sources/usb_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "usb_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0233311010100122-3103202203302102-2223112121120010-2331221331101103-3002112021222331-3023103233320103-3032001201110233-1233133230012303", "registry_path": "docs/guides/data-sources--usb_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/usb_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_usb_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["usb_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_usb_policy/data-source.tf`; digest `sha256:40aac30ea686e00cce92cecb2aee01cb69e14aef947ae725b686bf0b6e9af7b0`.

```terraform
# UsbPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UsbPolicy by name
data "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}

output "usb_policy_id" {
  value = data.xcsh_usb_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/examples/)
- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/)
