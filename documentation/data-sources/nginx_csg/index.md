---
page_title: "xcsh_nginx_csg"
subcategory: ""
description: "Manages a Nginx Csg resource in F5 Distributed Cloud for get nginx csg configuration. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["nginx csg"], "body_bytes": 1341, "body_sha256": "sha256:52af73d2ea8ce980ba5432dace47aaad18015192c52cd9083ea6396d6556abb5", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_csg:reference", "xcsh-docs:data-sources:nginx_csg:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_csg:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_csg:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/nginx_csg/index.md", "product": "distributed-cloud", "provider_name": "nginx_csg", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3010010030112133-0023203113011300-1111020333302000-2330320322101020-2330023211231000-3213122030322222-1002103132233101-0213032020002111", "registry_path": "docs/data-sources/nginx_csg.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_csg/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Nginx Csg resource in F5 Distributed Cloud for get nginx csg configuration. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_nginx_csg

Breadcrumbs:

- xcsh_nginx_csg

Manages a Nginx Csg resource in F5 Distributed Cloud for get nginx csg configuration. configuration.
(read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxCsg Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxCsg by name
data "xcsh_nginx_csg" "example" {
  name      = "example-nginx-csg"
  namespace = "staging"
}

output "nginx_csg_id" {
  value = data.xcsh_nginx_csg.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/examples/)
