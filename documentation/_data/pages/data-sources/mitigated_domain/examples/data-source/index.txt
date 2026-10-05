---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_mitigated_domain."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1331, "body_sha256": "sha256:4c88c91fc8940393673f8efd7a091c0e422d7c7b770b4ee3774eee328e94b8b4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:mitigated_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:bc0c9edeff4c50185f9f5d99bbe4c09ceec7f010ca7ea7432a8786d884607a16", "source_path": "examples/data-sources/xcsh_mitigated_domain/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:mitigated_domain:example:data-source", "parent_id": "xcsh-docs:data-sources:mitigated_domain:examples", "path": "documentation/data-sources/mitigated_domain/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "mitigated_domain", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-3313030103121120-3301201332001313-0103102033021113-3202232130333030-1211001200203032-1130030301302313-3303232310100111-0003021032202212", "registry_path": "docs/guides/data-sources--mitigated_domain--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/mitigated_domain/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_mitigated_domain.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["mitigated_domainCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_mitigated_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_mitigated_domain/data-source.tf`; digest `sha256:bc0c9edeff4c50185f9f5d99bbe4c09ceec7f010ca7ea7432a8786d884607a16`.

```terraform
# MitigatedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MitigatedDomain by name
data "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"
}

output "mitigated_domain_id" {
  value = data.xcsh_mitigated_domain.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/examples/)
- [xcsh_mitigated_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/mitigated_domain/)
