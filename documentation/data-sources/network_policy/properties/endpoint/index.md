---
page_title: "endpoint"
subcategory: "Security"
description: "endpoint for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2607, "body_sha256": "sha256:48221a2250755c9809b90c0f8cb1276fb2a88559c63b59b9d5878c28e929f0ed", "child_ids": ["xcsh-docs:data-sources:network_policy:properties:endpoint:any", "xcsh-docs:data-sources:network_policy:properties:endpoint:inside_endpoints", "xcsh-docs:data-sources:network_policy:properties:endpoint:label_selector", "xcsh-docs:data-sources:network_policy:properties:endpoint:outside_endpoints", "xcsh-docs:data-sources:network_policy:properties:endpoint:prefix_list"], "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:properties:endpoint", "parent_id": "xcsh-docs:data-sources:network_policy:reference", "path": "documentation/data-sources/network_policy/properties/endpoint/index.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["endpoint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/properties/endpoint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
