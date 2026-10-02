---
page_title: "stateful_service.advertise_options"
subcategory: "Container"
description: "Advertise OPTIONS are used to configure how and where to advertise the workload using load balancers."
xcsh_docs: {"aliases": ["stateful service advertise options"], "body_bytes": 2871, "body_sha256": "sha256:444d5806d9e2ed775d87786733efe9dd1afa5f177c63d42d3a4e2c8e1e4f38d3", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public", "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:do_not_advertise"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:workload:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options", "parent_id": "xcsh-docs:data-sources:workload:properties:stateful_service", "path": "documentation/data-sources/workload/properties/stateful_service/advertise_options/index.md", "product": "distributed-cloud", "provider_name": "workload", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2031021020201310-1033011221123311-0101031100313233-2003312312021132-1232222222022302-3032300332013210-2302212320011133-1111233212001220", "registry_path": "docs/guides/data-sources--workload--reference--group-017.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["stateful_service", "advertise_options"], "schema_version": 1, "sections": [{"aliases": ["advertise custom"], "anchor": "section", "description": "Advertise this workload via loadbalancer on specific sites.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_custom", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_custom"], "syntax": "attribute", "type": "object"}, {"aliases": ["advertise in cluster"], "anchor": "section", "description": "Advertise the workload locally in-cluster.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_in_cluster", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_in_cluster"], "syntax": "attribute", "type": "object"}, {"aliases": ["advertise on public"], "anchor": "section", "description": "Advertise this workload via loadbalancer on Internet with default VIP.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:advertise_on_public", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["stateful_service", "advertise_options", "advertise_on_public"], "syntax": "attribute", "type": "object"}, {"aliases": ["do not advertise"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:workload:properties:stateful_service:advertise_options:do_not_advertise", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["stateful_service", "advertise_options", "do_not_advertise"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload/properties/stateful_service/advertise_options/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Advertise OPTIONS are used to configure how and where to advertise the workload using load balancers.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workloadCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# stateful_service.advertise_options

Breadcrumbs:

- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- stateful_service.advertise_options

<a id="section"></a>

Type: `"single"`. Computed.

Advertise OPTIONS are used to configure how and where to advertise the workload using load
balancers.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-advertise_choice": "[\"advertise_custom\",\"advertise_in_cluster\",\"advertise_on_public\",\"do_not_advertise\"]"
}
```

## Direct properties

- [advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/): complete subsection reference.

- [advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/): complete subsection reference.

- [advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/): complete subsection reference.

- [do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/do_not_advertise/): complete subsection reference.

## Next pages

- [stateful_service.advertise_options.advertise_custom](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_custom/)
- [stateful_service.advertise_options.advertise_in_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_in_cluster/)
- [stateful_service.advertise_options.advertise_on_public](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/advertise_on_public/)
- [stateful_service.advertise_options.do_not_advertise](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/advertise_options/do_not_advertise/)
- [stateful_service](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/properties/stateful_service/)
- [xcsh_workload](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/workload/)
