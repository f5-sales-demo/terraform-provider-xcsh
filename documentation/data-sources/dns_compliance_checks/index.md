---
page_title: "xcsh_dns_compliance_checks"
subcategory: ""
description: "Reads DNS Compliance Checks information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["dns compliance checks"], "body_bytes": 1414, "body_sha256": "sha256:14db4875c7b2ce80db7041c169e55ebec54cae2e4570273dc3b89647aa061b63", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_compliance_checks:reference", "xcsh-docs:data-sources:dns_compliance_checks:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_compliance_checks:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_compliance_checks:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/dns_compliance_checks/index.md", "product": "distributed-cloud", "provider_name": "dns_compliance_checks", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0031332033312303-1303103011111022-1022102332331120-2133012122303100-3003222020332210-1130031122302313-3111210203332032-1102020121010322", "registry_path": "docs/data-sources/dns_compliance_checks.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_compliance_checks/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Reads DNS Compliance Checks information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["dns_compliance_checksCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_compliance_checks

Breadcrumbs:

- xcsh_dns_compliance_checks

Reads DNS Compliance Checks information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSComplianceChecks Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSComplianceChecks by name
data "xcsh_dns_compliance_checks" "example" {
  name      = "example-dns-compliance-checks"
  namespace = "staging"
}

output "dns_compliance_checks_id" {
  value = data.xcsh_dns_compliance_checks.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/examples/)
