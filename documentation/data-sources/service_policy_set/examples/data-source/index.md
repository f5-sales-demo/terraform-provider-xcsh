---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_service_policy_set."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1355, "body_sha256": "sha256:1b71640fb84473862ba406ec1ef5496ec0d3c0fa15146e525f8526dbec6701f4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:service_policy_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3bd6ef68376941c1abe8ba51cff222b7a95e9a7e72218b34827623539153d92a", "source_path": "examples/data-sources/xcsh_service_policy_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:service_policy_set:example:data-source", "parent_id": "xcsh-docs:data-sources:service_policy_set:examples", "path": "documentation/data-sources/service_policy_set/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "service_policy_set", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "data-sources", "registry_anchor": "canonical-0330102310230122-1223203021122031-0203023200302031-0313123001221103-2231000301310011-2303023113023020-2021003211310330-1333032122312302", "registry_path": "docs/guides/data-sources--service_policy_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/service_policy_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_service_policy_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_service_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_service_policy_set/data-source.tf`; digest `sha256:3bd6ef68376941c1abe8ba51cff222b7a95e9a7e72218b34827623539153d92a`.

```terraform
# ServicePolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicySet by name
data "xcsh_service_policy_set" "example" {
  name      = "example-service-policy-set"
  namespace = "staging"
}

output "service_policy_set_id" {
  value = data.xcsh_service_policy_set.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/examples/)
- [xcsh_service_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/service_policy_set/)
