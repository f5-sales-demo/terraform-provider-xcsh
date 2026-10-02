---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_code_base_integration."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1326, "body_sha256": "sha256:5fb34bcc68cfa07739e8330379e2bf8004dc1cd3d09c31076261c858282cb2c5", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:code_base_integration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:96e41779f04da03e1e24b87e34c40ee5955713fd86df36344fb3cbec779ab18f", "source_path": "examples/resources/xcsh_code_base_integration/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:code_base_integration:example:resource", "parent_id": "xcsh-docs:resources:code_base_integration:examples", "path": "documentation/resources/code_base_integration/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0302131003030332-2210232300310230-2102122221132322-3031133031321130-3332332002121301-0322313030311313-3101101132020113-2300310303302110", "registry_path": "docs/guides/resources--code_base_integration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/code_base_integration/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_code_base_integration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_code_base_integration/resource.tf`; digest `sha256:96e41779f04da03e1e24b87e34c40ee5955713fd86df36344fb3cbec779ab18f`.

```terraform
# CodeBaseIntegration Resource Example
# Manages integration details in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CodeBaseIntegration configuration
resource "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/examples/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/code_base_integration/)
