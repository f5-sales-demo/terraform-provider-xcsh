---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nginx_server."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1279, "body_sha256": "sha256:f55629e478cedd4d0533642f94d55dc90973204ed168291dc38c3be4d57fe11f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:255e0e3089bba8c6998465d72659b209b4c83e33fec880e96e0bccd476d39e6e", "source_path": "examples/data-sources/xcsh_nginx_server/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nginx_server:example:data-source", "parent_id": "xcsh-docs:data-sources:nginx_server:examples", "path": "documentation/data-sources/nginx_server/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "nginx_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3222330301303300-0232330221032232-3300301221130110-1220023211310012-3300131212330201-2103330102213200-1301120232101131-2312031002211310", "registry_path": "docs/guides/data-sources--nginx_server--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_nginx_server.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_server/data-source.tf`; digest `sha256:255e0e3089bba8c6998465d72659b209b4c83e33fec880e96e0bccd476d39e6e`.

```terraform
# NginxServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxServer by name
data "xcsh_nginx_server" "example" {
  name      = "example-nginx-server"
  namespace = "staging"
}

output "nginx_server_id" {
  value = data.xcsh_nginx_server.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/examples/)
- [xcsh_nginx_server](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_server/)
