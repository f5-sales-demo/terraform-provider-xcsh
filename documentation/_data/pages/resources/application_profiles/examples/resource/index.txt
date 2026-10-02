---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_application_profiles."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1315, "body_sha256": "sha256:109bf4b6a69c5661072726ca84b9571783e558299a7530b03dec4949f46ae7d0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:application_profiles:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9b63d4ae5d7cafe5639b8e62df3ae35e92aa58fd14501f2e62af1a281ea40d1f", "source_path": "examples/resources/xcsh_application_profiles/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:application_profiles:example:resource", "parent_id": "xcsh-docs:resources:application_profiles:examples", "path": "documentation/resources/application_profiles/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "application_profiles", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-0320321111331220-0331023100121133-3000212131200102-2302232203333031-3221123230331223-2102011000000010-1100112021001032-2031132130301202", "registry_path": "docs/guides/resources--application_profiles--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/application_profiles/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_application_profiles.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["application_profilesCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_application_profiles/resource.tf`; digest `sha256:9b63d4ae5d7cafe5639b8e62df3ae35e92aa58fd14501f2e62af1a281ea40d1f`.

```terraform
# ApplicationProfiles Resource Example
# Manages Application Profiles in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ApplicationProfiles configuration
resource "xcsh_application_profiles" "example" {
  name      = "example-application-profiles"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/examples/)
- [xcsh_application_profiles](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/application_profiles/)
