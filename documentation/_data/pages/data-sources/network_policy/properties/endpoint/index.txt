---
page_title: "endpoint"
subcategory: "Security"
description: "Shape of the endpoint choices for a view."
xcsh_docs: {"aliases": ["endpoint"], "body_bytes": 2607, "body_sha256": "sha256:48221a2250755c9809b90c0f8cb1276fb2a88559c63b59b9d5878c28e929f0ed", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_policy:properties:endpoint:any", "xcsh-docs:data-sources:network_policy:properties:endpoint:inside_endpoints", "xcsh-docs:data-sources:network_policy:properties:endpoint:label_selector", "xcsh-docs:data-sources:network_policy:properties:endpoint:outside_endpoints", "xcsh-docs:data-sources:network_policy:properties:endpoint:prefix_list"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:properties:endpoint", "parent_id": "xcsh-docs:data-sources:network_policy:reference", "path": "documentation/data-sources/network_policy/properties/endpoint/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2010012201232320-2010013322332122-0013213131201033-3002130013232130-3220032133312220-3320003223121202-3200123230202233-3230211302203332", "registry_path": "docs/guides/data-sources--network_policy--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint"], "schema_version": 1, "sections": [{"aliases": ["endpoint any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy:properties:endpoint:any", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint inside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy:properties:endpoint:inside_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "inside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:network_policy:properties:endpoint:label_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint", "label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint outside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy:properties:endpoint:outside_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "outside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["endpoint prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:network_policy:properties:endpoint:prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint", "prefix_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/properties/endpoint/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Shape of the endpoint choices for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["network_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/)
- endpoint

<a id="section"></a>

Type: `"single"`. Computed.

Shape of the endpoint choices for a view.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-endpoint_choice": "[\"any\",\"inside_endpoints\",\"label_selector\",\"outside_endpoints\",\"prefix_list\"]"
}
```

## Direct properties

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/endpoint/any/): complete subsection reference.

- [inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/endpoint/inside_endpoints/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/endpoint/label_selector/): complete subsection reference.

- [outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/endpoint/outside_endpoints/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/endpoint/prefix_list/): complete subsection reference.

## Next pages

- [endpoint.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/endpoint/any/)
- [endpoint.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/endpoint/inside_endpoints/)
- [endpoint.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/endpoint/label_selector/)
- [endpoint.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/endpoint/outside_endpoints/)
- [endpoint.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/endpoint/prefix_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/properties/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
