---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_usb_policy."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1209, "body_sha256": "sha256:35cda0b5a6a7fd79c315147accd167168b803929761f20c86476488fb1f0ff70", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:usb_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d81199b0fb106ffd6cc8d506207e603f1169d2cbf87d6b31f0ac720b8b44794e", "source_path": "examples/resources/xcsh_usb_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:usb_policy:example:resource", "parent_id": "xcsh-docs:resources:usb_policy:examples", "path": "documentation/resources/usb_policy/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "usb_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1021311230322121-1223221313333000-3302100222013302-0322101321332312-2111102333322331-3020210003100310-1200012222211011-1102301103010312", "registry_path": "docs/guides/resources--usb_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/usb_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_usb_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["usb_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_usb_policy/resource.tf`; digest `sha256:d81199b0fb106ffd6cc8d506207e603f1169d2cbf87d6b31f0ac720b8b44794e`.

```terraform
# UsbPolicy Resource Example
# Manages new USB policy object in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UsbPolicy configuration
resource "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/examples/)
- [xcsh_usb_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/usb_policy/)
