---
page_title: "service.containers.default_flavor"
subcategory: "Container"
description: "This can be used for messages where no values are needed."
xcsh_docs: {"aliases": ["service containers default flavor"], "body_bytes": 1381, "body_sha256": "sha256:1f4abf748005747251a8a25a29f1b1568d2ada43d10394ad2f274ed869a85ce0", "capabilities": ["container"], "category": "container", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:containers:default_flavor", "parent_id": "xcsh-docs:resources:workload:properties:service:containers", "path": "documentation/resources/workload/properties/service/containers/default_flavor/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3001233032001100-3130331312023200-1301300310120122-1301123133311223-2123211012022020-1331112121220323-2202211100213203-0020000120133312", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "containers", "default_flavor"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/containers/default_flavor/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "This can be used for messages where no values are needed.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.containers.default_flavor

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/)
- service.containers.default_flavor

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for default flavor.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
default_flavor = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [service.containers](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/containers/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
