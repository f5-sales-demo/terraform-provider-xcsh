---
page_title: "service.volumes.empty_dir"
subcategory: "Container"
description: "Volume containing a temporary directory whose lifetime is the same as a replica of a workload."
xcsh_docs: {"aliases": ["service volumes empty dir"], "body_bytes": 2441, "body_sha256": "sha256:f7ac8e5b6834388c46afaba4143b03e52221553d19c1090f32b18372c10ff638", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:workload:properties:service:volumes:empty_dir:mount"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:workload:collection", "completeness": "complete", "id": "xcsh-docs:resources:workload:properties:service:volumes:empty_dir", "parent_id": "xcsh-docs:resources:workload:properties:service:volumes", "path": "documentation/resources/workload/properties/service/volumes/empty_dir/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1221223030032300-2321120333220021-3120212121011003-2323110223332233-3233022113111010-1300231132321110-2323033323222232-2220132032230000", "registry_path": "docs/guides/resources--workload--reference--group-016.md", "relationships": [{"anchor": "schema-service--volumes--empty_dir--size_limit", "enforcement": "provider-schema", "group": "service.volumes.empty_dir:RequiredObjectAttributes:size_limit", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:volumes:empty_dir", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "volumes", "empty_dir"], "schema_version": 1, "sections": [{"aliases": ["service volumes empty dir mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:resources:workload:properties:service:volumes:empty_dir:mount", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-service--volumes--empty_dir--mount--mount_path", "enforcement": "provider-schema", "group": "service.volumes.empty_dir.mount:RequiredObjectAttributes:mount_path", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:workload:properties:service:volumes:empty_dir:mount", "type": "requires"}], "schema_path": ["service", "volumes", "empty_dir", "mount"], "syntax": "block", "type": "object"}, {"aliases": ["service volumes empty dir size limit"], "anchor": "schema-service--volumes--empty_dir--size_limit", "description": "Configuration parameter for size limit", "document_id": "xcsh-docs:resources:workload:properties:service:volumes:empty_dir", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "volumes", "empty_dir", "size_limit"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/workload/properties/service/volumes/empty_dir/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Volume containing a temporary directory whose lifetime is the same as a replica of a workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["workloadCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes.empty_dir

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/)
- service.volumes.empty_dir

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("size_limit")}
```

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
empty_dir {
  # Configure direct properties listed below.
}
```

## Direct properties

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/empty_dir/mount/): complete subsection reference.

<a id="schema-service--volumes--empty_dir--size_limit"></a>

### size_limit property

Type: `"number"`. Optional.

Size Limit (in GiB). Configuration parameter for size limit

Upstream description:

Configuration parameter for size limit

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.double.lte": "10",
    "ves.io.schema.rules.message.required": "true"
  }
}
```

## Next pages

- [service.volumes.empty_dir.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/empty_dir/mount/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/properties/service/volumes/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/workload/)
