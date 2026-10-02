---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dns_compliance_checks."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1378, "body_sha256": "sha256:11957eae165393e63696e06d128e60b5b8aacdcc539868dd96799aeb1101849c", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_compliance_checks:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4b55554787f8c66271c8f849e8f82effc0413dcd79554cefe77c85a2e413eba1", "source_path": "examples/resources/xcsh_dns_compliance_checks/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_compliance_checks:example:resource", "parent_id": "xcsh-docs:resources:dns_compliance_checks:examples", "path": "documentation/resources/dns_compliance_checks/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dns_compliance_checks", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0310000323310100-1023120330100322-0333100003022130-1231033132331112-0002303000233003-1030200222032233-3320101210031232-2022232323030301", "registry_path": "docs/guides/resources--dns_compliance_checks--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_compliance_checks/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_dns_compliance_checks.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_compliance_checksCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
