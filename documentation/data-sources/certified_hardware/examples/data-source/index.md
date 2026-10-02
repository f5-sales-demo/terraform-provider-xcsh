---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_certified_hardware."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1357, "body_sha256": "sha256:b91985fdaf4d778d416188c7c2813ce84e241404ecdb3761bfe056af2c33ba22", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:certified_hardware:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:44db2061b6da9d984930695f1d32baecee3d5ce0c74f32e0c377962e7ba246ba", "source_path": "examples/data-sources/xcsh_certified_hardware/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:certified_hardware:example:data-source", "parent_id": "xcsh-docs:data-sources:certified_hardware:examples", "path": "documentation/data-sources/certified_hardware/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "certified_hardware", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1001223122213112-2321321330020223-3001031021223130-0312000001213120-2202031212330020-1033003021022202-3021321100302202-2313303211212213", "registry_path": "docs/guides/data-sources--certified_hardware--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/certified_hardware/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_certified_hardware.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certified_hardware/data-source.tf`; digest `sha256:44db2061b6da9d984930695f1d32baecee3d5ce0c74f32e0c377962e7ba246ba`.

```terraform
# CertifiedHardware Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CertifiedHardware by name
data "xcsh_certified_hardware" "example" {
  name      = "example-certified-hardware"
  namespace = "staging"
}

output "certified_hardware_id" {
  value = data.xcsh_certified_hardware.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/examples/)
- [xcsh_certified_hardware](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/certified_hardware/)
