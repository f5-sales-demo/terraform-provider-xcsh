---
page_title: "xcsh_infraprotect_mitigation_ips"
subcategory: ""
description: "Reads Infraprotect Mitigation Ips information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["infraprotect mitigation ips"], "body_bytes": 1426, "body_sha256": "sha256:0f1880dc2d2e652be60929e287548d1a762cd5e0979c9a2c3a62a8aaf47bacef", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:infraprotect_mitigation_ips:reference", "xcsh-docs:data-sources:infraprotect_mitigation_ips:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/infraprotect_mitigation_ips/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1223233033223000-1032001012013022-1301320302130033-2311101100202302-3130231002213230-3320300113130330-0113131111131313-3222111203302102", "registry_path": "docs/data-sources/infraprotect_mitigation_ips.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads Infraprotect Mitigation Ips information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_infraprotect_mitigation_ips

Breadcrumbs:

- xcsh_infraprotect_mitigation_ips

Reads Infraprotect Mitigation Ips information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# InfraprotectMitigationIps DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_infraprotect_mitigation_ips" "example" {
  mitigation_id = "example-value"
  namespace     = "example-value"
}

output "infraprotect_mitigation_ips_result" {
  value = data.xcsh_infraprotect_mitigation_ips.example
}
```

## Root configuration

Required root properties: `mitigation_id`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/examples/)
