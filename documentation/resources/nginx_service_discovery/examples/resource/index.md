---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_nginx_service_discovery."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1459, "body_sha256": "sha256:12099de8ca5b7e2bb1e8889325420747e503fe14d94bc8806d07289063d7c5dd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:dd43ba5ab8c8a702c201e30638f18d6c0eafcf265b5180bf65bbcaa03289e5e6", "source_path": "examples/resources/xcsh_nginx_service_discovery/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:nginx_service_discovery:example:resource", "parent_id": "xcsh-docs:resources:nginx_service_discovery:examples", "path": "documentation/resources/nginx_service_discovery/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0203300202021231-3011031103221111-0023202232111213-2230220131210232-1301122131321333-1303010122113202-2311331103011311-2122021221313311", "registry_path": "docs/guides/resources--nginx_service_discovery--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_nginx_service_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_nginx_service_discovery/resource.tf`; digest `sha256:dd43ba5ab8c8a702c201e30638f18d6c0eafcf265b5180bf65bbcaa03289e5e6`.

```terraform
# NginxServiceDiscovery Resource Example
# Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NginxServiceDiscovery configuration
resource "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/examples/)
- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/)
