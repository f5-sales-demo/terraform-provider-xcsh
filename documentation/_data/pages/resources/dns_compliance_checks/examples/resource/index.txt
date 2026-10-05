---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dns_compliance_checks."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1378, "body_sha256": "sha256:11957eae165393e63696e06d128e60b5b8aacdcc539868dd96799aeb1101849c", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_compliance_checks:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4b55554787f8c66271c8f849e8f82effc0413dcd79554cefe77c85a2e413eba1", "source_path": "examples/resources/xcsh_dns_compliance_checks/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_compliance_checks:example:resource", "parent_id": "xcsh-docs:resources:dns_compliance_checks:examples", "path": "documentation/resources/dns_compliance_checks/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dns_compliance_checks", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0310000323310100-1023120330100322-0333100003022130-1231033132331112-0002303000233003-1030200222032233-3320101210031232-2022232323030301", "registry_path": "docs/guides/resources--dns_compliance_checks--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_compliance_checks/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_dns_compliance_checks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["dns_compliance_checksCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_dns_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_compliance_checks/resource.tf`; digest `sha256:4b55554787f8c66271c8f849e8f82effc0413dcd79554cefe77c85a2e413eba1`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/examples/)
- [xcsh_dns_compliance_checks](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_compliance_checks/)
