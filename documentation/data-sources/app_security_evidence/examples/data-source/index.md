---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_security_evidence."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1305, "body_sha256": "sha256:72426b7b64a2c4762ab085f7da1e84659079f3173384cf734048a8e2c8b413df", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:app_security_evidence:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a712a6cc48d2b8d3894590d0570f0c9aa69c910e418b124dca6ff57b24d59d2a", "source_path": "examples/data-sources/xcsh_app_security_evidence/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_security_evidence:example:data-source", "parent_id": "xcsh-docs:data-sources:app_security_evidence:examples", "path": "documentation/data-sources/app_security_evidence/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "app_security_evidence", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1021122003121201-1300022010132211-1212102231110310-2122003313212300-1012310133112113-1113332321002312-2112112331331110-2221133121001122", "registry_path": "docs/guides/data-sources--app_security_evidence--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_security_evidence/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_app_security_evidence.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_security_evidence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_security_evidence/data-source.tf`; digest `sha256:a712a6cc48d2b8d3894590d0570f0c9aa69c910e418b124dca6ff57b24d59d2a`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/examples/)
- [xcsh_app_security_evidence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_security_evidence/)
