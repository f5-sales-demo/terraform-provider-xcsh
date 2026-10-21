---
page_title: "xcsh_dns_compliance_checks"
subcategory: ""
description: "Manages DNS Compliance Checks Specification in a given namespace. If one already exists it will give an error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["dns compliance checks"], "body_bytes": 1743, "body_sha256": "sha256:8cae2388cb1edc4b7c48dcabcf1a4cc32bc21b85f9cdecb2912211f2887482c1", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_compliance_checks:reference", "xcsh-docs:resources:dns_compliance_checks:examples", "xcsh-docs:resources:dns_compliance_checks:import", "xcsh-docs:resources:dns_compliance_checks:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_compliance_checks:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_compliance_checks:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/dns_compliance_checks/index.md", "product": "distributed-cloud", "provider_name": "dns_compliance_checks", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2320330031033230-1032132011323312-3322232320012321-0210130122222101-3131030333303122-1013130113312022-2213230000130100-1010120232210210", "registry_path": "docs/resources/dns_compliance_checks.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_compliance_checks/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Manages DNS Compliance Checks Specification in a given namespace. If one already exists it will give an error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["dns_compliance_checksCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_compliance_checks

Breadcrumbs:

- xcsh_dns_compliance_checks

Manages DNS Compliance Checks Specification in a given namespace. If one already exists it will give
an error in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSComplianceChecks Resource Example
# Manages DNS Compliance Checks Specification in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSComplianceChecks configuration
resource "xcsh_dns_compliance_checks" "example" {
  name      = "example-dns-compliance-checks"
  namespace = "staging"

  domain_denylist = ["example-value"]
}
```

## Root configuration

Required root properties: `domain_denylist`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/lifecycle/timeouts/)
