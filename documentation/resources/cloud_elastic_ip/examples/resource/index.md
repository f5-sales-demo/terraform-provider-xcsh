---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cloud_elastic_ip."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1347, "body_sha256": "sha256:a73c2e3dda4f60269b5fda65ed5508b148c1450003e7e8ca6284eca4563163a4", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cloud_elastic_ip:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e8be04df02915d54850afcd099a22da22eee85dc886b468e9bf65ca321094b18", "source_path": "examples/resources/xcsh_cloud_elastic_ip/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cloud_elastic_ip:example:resource", "parent_id": "xcsh-docs:resources:cloud_elastic_ip:examples", "path": "documentation/resources/cloud_elastic_ip/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "cloud_elastic_ip", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2303013103321313-0312200010132320-2031231330230232-2013133233303102-0010021011001023-2022030310033122-2213032031103011-1301301200210022", "registry_path": "docs/guides/resources--cloud_elastic_ip--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_elastic_ip/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cloud_elastic_ip.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_elastic_ipCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_cloud_elastic_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_elastic_ip/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_elastic_ip/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_cloud_elastic_ip/resource.tf`; digest `sha256:e8be04df02915d54850afcd099a22da22eee85dc886b468e9bf65ca321094b18`.

```terraform
# CloudElasticIP Resource Example
# Manages Cloud Elastic IP creates Cloud Elastic IP object Object is attached to a site in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CloudElasticIP configuration
resource "xcsh_cloud_elastic_ip" "example" {
  name      = "example-cloud-elastic-ip"
  namespace = "staging"

  item_count = 1
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_elastic_ip/examples/)
- [xcsh_cloud_elastic_ip](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cloud_elastic_ip/)
