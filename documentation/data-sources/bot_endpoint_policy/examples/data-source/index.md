---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_endpoint_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1368, "body_sha256": "sha256:aa5a013f38e366cfa8da7b3d47640976adcf4235e54f98e2ea21c781044e9488", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_endpoint_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:742dea18fab3d10fddee899cf0137a64f4bc6d260ffbdc975152403793a768eb", "source_path": "examples/data-sources/xcsh_bot_endpoint_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_endpoint_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_endpoint_policy:examples", "path": "documentation/data-sources/bot_endpoint_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_endpoint_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2013332022003012-1123320013322303-2011313130110113-3110223331110213-0223300030010201-0100322130311131-1321211030302021-0232303101301132", "registry_path": "docs/guides/data-sources--bot_endpoint_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_endpoint_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_endpoint_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_endpoint_policy/data-source.tf`; digest `sha256:742dea18fab3d10fddee899cf0137a64f4bc6d260ffbdc975152403793a768eb`.

```terraform
# BotEndpointPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotEndpointPolicy by name
data "xcsh_bot_endpoint_policy" "example" {
  name      = "example-bot-endpoint-policy"
  namespace = "staging"
}

output "bot_endpoint_policy_id" {
  value = data.xcsh_bot_endpoint_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/examples/)
- [xcsh_bot_endpoint_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_endpoint_policy/)
