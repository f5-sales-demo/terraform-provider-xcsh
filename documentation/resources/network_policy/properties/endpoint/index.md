---
page_title: "endpoint"
subcategory: "Security"
description: "endpoint for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 3562, "body_sha256": "sha256:1b9220de45777fa137f6a308ff0afa9d7d712b87b6dbfbcd8228da152e30378f", "child_ids": ["xcsh-docs:resources:network_policy:properties:endpoint:any", "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list"], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:endpoint", "parent_id": "xcsh-docs:resources:network_policy:reference", "path": "documentation/resources/network_policy/properties/endpoint/index.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "properties", "schema_path": ["endpoint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/endpoint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- endpoint

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Shape of the endpoint choices for a view.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("any",
    "inside_endpoints"),
  validators.ConflictingObjectAttributes("any",
    "label_selector"),
  validators.ConflictingObjectAttributes("any",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("any",
    "prefix_list"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "label_selector"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("inside_endpoints",
    "prefix_list"),
  validators.ConflictingObjectAttributes("label_selector",
    "outside_endpoints"),
  validators.ConflictingObjectAttributes("label_selector",
    "prefix_list"),
  validators.ConflictingObjectAttributes("outside_endpoints",
    "prefix_list")}
```

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

Terraform syntax:

```terraform
endpoint {
  # Configure direct properties listed below.
}
```

## Direct properties

- [any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/any/): complete subsection reference.

- [inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/inside_endpoints/): complete subsection reference.

- [label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/label_selector/): complete subsection reference.

- [outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/outside_endpoints/): complete subsection reference.

- [prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/prefix_list/): complete subsection reference.

## Next pages

- [endpoint.any](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/any/)
- [endpoint.inside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/inside_endpoints/)
- [endpoint.label_selector](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/label_selector/)
- [endpoint.outside_endpoints](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/outside_endpoints/)
- [endpoint.prefix_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/endpoint/prefix_list/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/)
