---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dns_compliance_checks."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1139, "body_sha256": "sha256:ccf07f047d6ae958a4527e4204631bd950e977c81139a6c1d541e95a85cc81c8", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dns_compliance_checks:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:494f38d0fa830fb2beff3f30865ce3a3a079a872bc62acf125c470d1e3df90d4", "source_path": "examples/data-sources/xcsh_dns_compliance_checks/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_compliance_checks:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_compliance_checks:examples", "path": "documentation/data-sources/dns_compliance_checks/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "dns_compliance_checks", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3223301002123130-1102123320223113-1000323131213300-2020130331010232-3112022012022111-3203321303113121-1112002121312113-2331320022132211", "registry_path": "docs/guides/data-sources--dns_compliance_checks--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_compliance_checks/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_dns_compliance_checks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["dns_compliance_checksCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_compliance_checks/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_compliance_checks/data-source.tf`; digest `sha256:494f38d0fa830fb2beff3f30865ce3a3a079a872bc62acf125c470d1e3df90d4`.

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
