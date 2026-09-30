---
page_title: "endpoint"
subcategory: "Security"
description: "endpoint for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 2755, "body_sha256": "sha256:1d6dbe12486a3f5210edfe3ac4aab134e934f2725f154f05964a20bf57ad3731", "canonical_id": "xcsh-docs:resources:network_policy:properties:endpoint", "child_ids": ["xcsh-docs:resources:network_policy:properties:endpoint:any", "xcsh-docs:resources:network_policy:properties:endpoint:inside_endpoints", "xcsh-docs:resources:network_policy:properties:endpoint:label_selector", "xcsh-docs:resources:network_policy:properties:endpoint:outside_endpoints", "xcsh-docs:resources:network_policy:properties:endpoint:prefix_list"], "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:properties:endpoint", "parent_id": "xcsh-docs:resources:network_policy:reference", "path": "docs/guides/resources--network_policy--properties--endpoint.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/properties/endpoint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# endpoint

Breadcrumbs:

- [xcsh_network_policy](../resources/network_policy.md)
- [Property reference](resources--network_policy--reference.md)
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

- [any](resources--network_policy--properties--endpoint--any.md): complete subsection reference.

- [inside_endpoints](resources--network_policy--properties--endpoint--inside_endpoints.md): complete subsection reference.

- [label_selector](resources--network_policy--properties--endpoint--label_selector.md): complete subsection reference.

- [outside_endpoints](resources--network_policy--properties--endpoint--outside_endpoints.md): complete subsection reference.

- [prefix_list](resources--network_policy--properties--endpoint--prefix_list.md): complete subsection reference.

## Next pages

- [endpoint.any](resources--network_policy--properties--endpoint--any.md)
- [endpoint.inside_endpoints](resources--network_policy--properties--endpoint--inside_endpoints.md)
- [endpoint.label_selector](resources--network_policy--properties--endpoint--label_selector.md)
- [endpoint.outside_endpoints](resources--network_policy--properties--endpoint--outside_endpoints.md)
- [endpoint.prefix_list](resources--network_policy--properties--endpoint--prefix_list.md)
- [Property reference](resources--network_policy--reference.md)
- [xcsh_network_policy](../resources/network_policy.md)
