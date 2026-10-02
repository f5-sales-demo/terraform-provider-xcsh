---
page_title: "xcsh_protocol_policer"
subcategory: ""
description: "Manages protocol_policer object, protocol_policer object contains list of L4 protocol match condition and corresponding traffic rate limits in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["protocol policer"], "body_bytes": 1439, "body_sha256": "sha256:52bdbe5dfb940072694197ba3741c49c05ea52cd585e101dbd5c57185a4e29e9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:protocol_policer:reference", "xcsh-docs:data-sources:protocol_policer:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:protocol_policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protocol_policer:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/protocol_policer/index.md", "product": "distributed-cloud", "provider_name": "protocol_policer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0333320001113103-2301232212333300-2111200131000031-2123222220130110-2130232032101120-0000213212312330-3312310223000320-3003231321101232", "registry_path": "docs/data-sources/protocol_policer.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protocol_policer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages protocol_policer object, protocol_policer object contains list of L4 protocol match condition and corresponding traffic rate limits in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protocol_policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_protocol_policer

Breadcrumbs:

- xcsh_protocol_policer

Manages protocol\_policer object, protocol\_policer object contains list of L4 protocol match
condition and corresponding traffic rate limits in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtocolPolicer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolPolicer by name
data "xcsh_protocol_policer" "example" {
  name      = "example-protocol-policer"
  namespace = "system"
}

output "protocol_policer_id" {
  value = data.xcsh_protocol_policer.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protocol_policer/examples/)
