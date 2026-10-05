---
page_title: "xcsh_app_security_evidence"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["app security evidence"], "body_bytes": 1266, "body_sha256": "sha256:f84e6a5f4cbb2673f762e57ec41e7698398e8efc07def7baa17c4ff06b64eb0a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_security_evidence:reference", "xcsh-docs:data-sources:app_security_evidence:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_security_evidence:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_security_evidence:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/app_security_evidence/index.md", "product": "distributed-cloud", "provider_name": "app_security_evidence", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3103302000231201-2113100232313221-0102211103022300-2213011330310113-3301220213112210-1312012211213332-0301203231301211-3001330131022123", "registry_path": "docs/data-sources/app_security_evidence.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_security_evidence/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_app_security_evidence

Breadcrumbs:

- xcsh_app_security_evidence

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppSecurityEvidence DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_app_security_evidence" "example" {
  namespace = "example-value"
}

output "app_security_evidence_result" {
  value = data.xcsh_app_security_evidence.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/examples/)
