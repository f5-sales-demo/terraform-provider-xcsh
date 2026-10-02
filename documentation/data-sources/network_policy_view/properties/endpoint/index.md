---
page_title: "endpoint"
subcategory: ""
description: "Shape of the endpoint choices for a view."
xcsh_docs: {"aliases": ["endpoint"], "body_bytes": 2687, "body_sha256": "sha256:ab01fa7e0e5497e6afbe021832f365909aa25bed892395009e6e411312caa31b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:endpoint:any", "xcsh-docs:data-sources:network_policy_view:properties:endpoint:inside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:endpoint:label_selector", "xcsh-docs:data-sources:network_policy_view:properties:endpoint:outside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:endpoint:prefix_list"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:endpoint", "parent_id": "xcsh-docs:data-sources:network_policy_view:reference", "path": "documentation/data-sources/network_policy_view/properties/endpoint/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2013033230003031-2133221231232102-0002022323032133-2321012131320311-0112333301203110-0113030102232111-3310122020311022-3221111322120120", "registry_path": "docs/guides/data-sources--network_policy_view--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["endpoint"], "schema_version": 1, "sections": [{"aliases": ["any"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:endpoint:any", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "any"], "syntax": "attribute", "type": "object"}, {"aliases": ["inside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:endpoint:inside_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "inside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["label selector"], "anchor": "section", "description": "This type can be used to establish a 'selector reference' from one object(called selector) to a set of other objects(called selectees) based on the value of expressions. A label selector is a label query over a set of resources. An empty label selector matches all objects. A null label selector matches no objects.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:endpoint:label_selector", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint", "label_selector"], "syntax": "attribute", "type": "object"}, {"aliases": ["outside endpoints"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:endpoint:outside_endpoints", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["endpoint", "outside_endpoints"], "syntax": "attribute", "type": "object"}, {"aliases": ["prefix list"], "anchor": "section", "description": "List of IPv4 prefixes that represent an endpoint.", "document_id": "xcsh-docs:data-sources:network_policy_view:properties:endpoint:prefix_list", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["endpoint", "prefix_list"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/endpoint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Shape of the endpoint choices for a view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/)
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

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/any/): complete subsection reference.

- [inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/inside_endpoints/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/label_selector/): complete subsection reference.

- [outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/outside_endpoints/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/prefix_list/): complete subsection reference.

## Next pages

- [endpoint.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/any/)
- [endpoint.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/inside_endpoints/)
- [endpoint.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/label_selector/)
- [endpoint.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/outside_endpoints/)
- [endpoint.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/endpoint/prefix_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/properties/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
