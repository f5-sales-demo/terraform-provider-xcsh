---
page_title: "xcsh_nginx_service_discovery"
subcategory: ""
description: "Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace. configuration."
xcsh_docs: {"aliases": ["nginx service discovery"], "body_bytes": 1533, "body_sha256": "sha256:71a22b6dbdc1ead7fc5e3c5e5914414e36e6edc9521cc203a99fb2d96d3869aa", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:nginx_service_discovery:reference", "xcsh-docs:data-sources:nginx_service_discovery:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_service_discovery:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/nginx_service_discovery/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0200221313232203-1031330113023320-2021213031131323-2131001232200331-1100111223302322-2103332312010110-1212231333021023-0201033030332311", "registry_path": "docs/data-sources/nginx_service_discovery.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_service_discovery/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_nginx_service_discovery

Breadcrumbs:

- xcsh_nginx_service_discovery

Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service
discovery object for a site or virtual site in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxServiceDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxServiceDiscovery by name
data "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}

output "nginx_service_discovery_id" {
  value = data.xcsh_nginx_service_discovery.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/examples/)
