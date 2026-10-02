---
page_title: "xcsh_usb_policy"
subcategory: ""
description: "Manages new USB policy object in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["usb policy"], "body_bytes": 1281, "body_sha256": "sha256:a51bad337b2251a02a829d7f71abd2c0b7df374982bc1d33be83d999e8191d3e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:usb_policy:reference", "xcsh-docs:data-sources:usb_policy:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:usb_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:usb_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/usb_policy/index.md", "product": "distributed-cloud", "provider_name": "usb_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3133321203222121-0333232330020202-2311221201233121-1211132233302312-0010122132101103-2031031312101012-3000023221320113-0202220320111000", "registry_path": "docs/data-sources/usb_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/usb_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages new USB policy object in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["usb_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_usb_policy

Breadcrumbs:

- xcsh_usb_policy

Manages new USB policy object in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/usb_policy/examples/)
