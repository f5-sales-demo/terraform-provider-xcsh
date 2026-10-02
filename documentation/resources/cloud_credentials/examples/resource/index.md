---
page_title: "Resource"
subcategory: "Infrastructure"
description: "Resource for xcsh_cloud_credentials."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1336, "body_sha256": "sha256:f62b5cd695f6dbe929e309f4821c479a662d02a8820e047c05ce2a357445510b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cloud_credentials:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:59ca9345886180c4c5420989ec1dac0bd8bd93a36d37acaaca6a81918a8c91f1", "source_path": "examples/resources/xcsh_cloud_credentials/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cloud_credentials:example:resource", "parent_id": "xcsh-docs:resources:cloud_credentials:examples", "path": "documentation/resources/cloud_credentials/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cloud_credentials", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1311302123221221-1320030133321200-2312101213223331-1230210011313132-1110231031310201-0133030030003102-2212201012002100-0200233100032300", "registry_path": "docs/guides/resources--cloud_credentials--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_credentials/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cloud_credentials.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_credentialsCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_credentials/resource.tf`; digest `sha256:59ca9345886180c4c5420989ec1dac0bd8bd93a36d37acaaca6a81918a8c91f1`.

```terraform
# CloudCredentials Resource Example
# Manages a Cloud Credentials resource in F5 Distributed Cloud for api to create cloud_credentials object.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudCredentials configuration
resource "xcsh_cloud_credentials" "example" {
  name      = "example-cloud-credentials"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/examples/)
- [xcsh_cloud_credentials](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_credentials/)
