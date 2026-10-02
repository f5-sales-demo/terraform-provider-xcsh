---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_voltstack_site."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1305, "body_sha256": "sha256:6cd0d56cd15d24bc21ed2631969b0af90f8fc7e93fd2de461919eed2c967d2a3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:voltstack_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9d6c1f9af1750a76c0f8df78e6df342564c3af32fa86a487ff12441eece04681", "source_path": "examples/data-sources/xcsh_voltstack_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:voltstack_site:example:data-source", "parent_id": "xcsh-docs:data-sources:voltstack_site:examples", "path": "documentation/data-sources/voltstack_site/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "voltstack_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0102021313332110-1223033311121022-2101021213312010-1023032021203231-0322110220312000-0132332021232110-3020121012101232-1121212233201233", "registry_path": "docs/guides/data-sources--voltstack_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/voltstack_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_voltstack_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["voltstack_siteCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_voltstack_site/data-source.tf`; digest `sha256:9d6c1f9af1750a76c0f8df78e6df342564c3af32fa86a487ff12441eece04681`.

```terraform
# VoltstackSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VoltstackSite by name
data "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"
}

output "voltstack_site_id" {
  value = data.xcsh_voltstack_site.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/examples/)
- [xcsh_voltstack_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/voltstack_site/)
