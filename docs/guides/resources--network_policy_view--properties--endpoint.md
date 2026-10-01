---
page_title: "endpoint"
subcategory: ""
description: "endpoint for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 2934, "body_sha256": "sha256:5b8a1afea76871724600242203f47ee39a780e776419cbeab1e493c98feb45de", "canonical_id": "xcsh-docs:resources:network_policy_view:properties:endpoint", "child_ids": ["xcsh-docs:resources:network_policy_view:properties:endpoint:any", "xcsh-docs:resources:network_policy_view:properties:endpoint:inside_endpoints", "xcsh-docs:resources:network_policy_view:properties:endpoint:label_selector", "xcsh-docs:resources:network_policy_view:properties:endpoint:outside_endpoints", "xcsh-docs:resources:network_policy_view:properties:endpoint:prefix_list"], "collection_id": "xcsh-docs:resources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_view:properties:endpoint", "parent_id": "xcsh-docs:resources:network_policy_view:reference", "path": "docs/guides/resources--network_policy_view--properties--endpoint.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_view/properties/endpoint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint

Breadcrumbs:

- [xcsh_network_policy_view](../resources/network_policy_view.md)
- [Property reference](resources--network_policy_view--reference.md)
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

- [any](resources--network_policy_view--properties--endpoint--any.md): complete subsection reference.

- [inside_endpoints](resources--network_policy_view--properties--endpoint--inside_endpoints.md): complete subsection reference.

- [label_selector](resources--network_policy_view--properties--endpoint--label_selector.md): complete subsection reference.

- [outside_endpoints](resources--network_policy_view--properties--endpoint--outside_endpoints.md): complete subsection reference.

- [prefix_list](resources--network_policy_view--properties--endpoint--prefix_list.md): complete subsection reference.

## Next pages

- [endpoint.any](resources--network_policy_view--properties--endpoint--any.md)
- [endpoint.inside_endpoints](resources--network_policy_view--properties--endpoint--inside_endpoints.md)
- [endpoint.label_selector](resources--network_policy_view--properties--endpoint--label_selector.md)
- [endpoint.outside_endpoints](resources--network_policy_view--properties--endpoint--outside_endpoints.md)
- [endpoint.prefix_list](resources--network_policy_view--properties--endpoint--prefix_list.md)
- [Property reference](resources--network_policy_view--reference.md)
- [xcsh_network_policy_view](../resources/network_policy_view.md)
