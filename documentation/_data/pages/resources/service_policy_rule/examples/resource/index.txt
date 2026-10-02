---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_service_policy_rule."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1390, "body_sha256": "sha256:c688bdc2cdafa13fc4f9149b8911007db927d9b0e682111c5eca292d0f6db348", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:service_policy_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ddddb543e61078456915a99cb29c3a42a961001d81b7eb50761d3e637258e884", "source_path": "examples/resources/xcsh_service_policy_rule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:service_policy_rule:example:resource", "parent_id": "xcsh-docs:resources:service_policy_rule:examples", "path": "documentation/resources/service_policy_rule/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "service_policy_rule", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0031012330311210-2200131123310002-3303003331230202-0221301312323212-1332133113013300-3031023113223131-2121002220121312-2020130021122221", "registry_path": "docs/guides/resources--service_policy_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy_rule/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_service_policy_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["service_policy_ruleCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy_rule/resource.tf`; digest `sha256:ddddb543e61078456915a99cb29c3a42a961001d81b7eb50761d3e637258e884`.

```terraform
# ServicePolicyRule Resource Example
# Manages service_policy_rule creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ServicePolicyRule configuration
resource "xcsh_service_policy_rule" "example" {
  name      = "example-service-policy-rule"
  namespace = "staging"

  action = "DENY"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/examples/)
- [xcsh_service_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy_rule/)
