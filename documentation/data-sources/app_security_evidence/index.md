---
page_title: "xcsh_app_security_evidence"
subcategory: ""
description: "Reads App Security Evidence information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["app security evidence"], "body_bytes": 1317, "body_sha256": "sha256:2fa3eece648f3b7364a7c9fcb8c4fbde8ce751bd88df36003058859178e7a64d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:app_security_evidence:reference", "xcsh-docs:data-sources:app_security_evidence:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_security_evidence:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_security_evidence:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/app_security_evidence/index.md", "product": "distributed-cloud", "provider_name": "app_security_evidence", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3103302000231201-2113100232313221-0102211103022300-2213011330310113-3301220213112210-1312012211213332-0301203231301211-3001330131022123", "registry_path": "docs/data-sources/app_security_evidence.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_security_evidence/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reads App Security Evidence information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_app_security_evidence

Breadcrumbs:

- xcsh_app_security_evidence

Reads App Security Evidence information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/examples/)
