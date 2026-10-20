---
page_title: "xcsh_nginx_service_discovery"
subcategory: ""
description: "Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace. configuration."
xcsh_docs: {"aliases": ["nginx service discovery"], "body_bytes": 1847, "body_sha256": "sha256:c5235fbcbd4bb586153a176a051b84285b68492bc6d13ee81847c12cf31e3aa1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:nginx_service_discovery:reference", "xcsh-docs:resources:nginx_service_discovery:examples", "xcsh-docs:resources:nginx_service_discovery:import", "xcsh-docs:resources:nginx_service_discovery:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:nginx_service_discovery:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/nginx_service_discovery/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2133010001131111-1103202000012230-1200101233222202-0010112120031023-2011010122321223-1011233232030201-3203200000233102-2110012102122023", "registry_path": "docs/resources/nginx_service_discovery.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/nginx_service_discovery/lifecycle/timeouts/)
