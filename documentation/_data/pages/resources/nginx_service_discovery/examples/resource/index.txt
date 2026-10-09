---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_nginx_service_discovery."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1204, "body_sha256": "sha256:148dc793c6a51575e56c73aa7e7220dc21eb8e38e95ea6b6a8778e646d2f8167", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:dd43ba5ab8c8a702c201e30638f18d6c0eafcf265b5180bf65bbcaa03289e5e6", "source_path": "examples/resources/xcsh_nginx_service_discovery/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:nginx_service_discovery:example:resource", "parent_id": "xcsh-docs:resources:nginx_service_discovery:examples", "path": "documentation/resources/nginx_service_discovery/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-0203300202021231-3011031103221111-0023202232111213-2230220131210232-1301122131321333-1303010122113202-2311331103011311-2122021221313311", "registry_path": "docs/guides/resources--nginx_service_discovery--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/examples/resource/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Resource for xcsh_nginx_service_discovery.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
