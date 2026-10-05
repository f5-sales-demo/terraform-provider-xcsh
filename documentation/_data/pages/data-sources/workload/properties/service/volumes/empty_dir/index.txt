---
page_title: "service.volumes.empty_dir"
subcategory: "Container"
description: "Volume containing a temporary directory whose lifetime is the same as a replica of a workload."
xcsh_docs: {"aliases": ["service volumes empty dir"], "body_bytes": 2198, "body_sha256": "sha256:cba66d75a7741e2277d6dfc25d88a39292efad68c2228c6bd0b1ccb767de4760", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:service:volumes:empty_dir:mount"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:service:volumes:empty_dir", "parent_id": "xcsh-docs:data-sources:workload:properties:service:volumes", "path": "documentation/data-sources/workload/properties/service/volumes/empty_dir/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0002103122000011-0113223020303230-2332232221203300-3113120213201230-0003320012012003-1033230200333223-0300323232122131-3222113003320322", "registry_path": "docs/guides/data-sources--workload--reference--group-016.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["service", "volumes", "empty_dir"], "schema_version": 1, "sections": [{"aliases": ["service volumes empty dir mount"], "anchor": "section", "description": "Volume mount describes how volume is mounted inside a workload.", "document_id": "xcsh-docs:data-sources:workload:properties:service:volumes:empty_dir:mount", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["service", "volumes", "empty_dir", "mount"], "syntax": "attribute", "type": "object"}, {"aliases": ["service volumes empty dir size limit"], "anchor": "schema-service--volumes--empty_dir--size_limit", "description": "Configuration parameter for size limit", "document_id": "xcsh-docs:data-sources:workload:properties:service:volumes:empty_dir", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["service", "volumes", "empty_dir", "size_limit"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/service/volumes/empty_dir/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Volume containing a temporary directory whose lifetime is the same as a replica of a workload.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# service.volumes.empty_dir

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/)
- service.volumes.empty_dir

<a id="section"></a>

Type: `"single"`. Computed.

Volume containing a temporary directory whose lifetime is the same as a replica of a workload.

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

## Direct properties

- [mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/empty_dir/mount/): complete subsection reference.

<a id="schema-service--volumes--empty_dir--size_limit"></a>

### size_limit property

Type: `"number"`. Computed.

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

- [service.volumes.empty_dir.mount](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/empty_dir/mount/)
- [service.volumes](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/service/volumes/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
