---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_cloud_elastic_ip."
xcsh_docs: {"aliases": [], "body_bytes": 1042, "body_sha256": "sha256:9172ca13052c1d15e16ea4602babc5e9e8a6a4feb35cb5a6d085f2ef9819eed1", "canonical_id": "xcsh-docs:resources:cloud_elastic_ip:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:cloud_elastic_ip:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e8be04df02915d54850afcd099a22da22eee85dc886b468e9bf65ca321094b18", "source_path": "examples/resources/xcsh_cloud_elastic_ip/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:cloud_elastic_ip:example:resource", "parent_id": "xcsh-docs:resources:cloud_elastic_ip:examples", "path": "docs/guides/resources--cloud_elastic_ip--example--resource.md", "provider_name": "cloud_elastic_ip", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cloud_elastic_ip/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_cloud_elastic_ip.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_elastic_ipCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md)
- [Examples](resources--cloud_elastic_ip--examples.md)
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

- [Examples](resources--cloud_elastic_ip--examples.md)
- [xcsh_cloud_elastic_ip](../resources/cloud_elastic_ip.md)
