---
page_title: "endpoint"
subcategory: "Security"
description: "endpoint for xcsh_network_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1899, "body_sha256": "sha256:4438a9164256200b5e27161b068b70747bbcf9046d0fd6fe377600078fd9cdc2", "canonical_id": "xcsh-docs:data-sources:network_policy:properties:endpoint", "child_ids": ["xcsh-docs:data-sources:network_policy:properties:endpoint:any", "xcsh-docs:data-sources:network_policy:properties:endpoint:inside_endpoints", "xcsh-docs:data-sources:network_policy:properties:endpoint:label_selector", "xcsh-docs:data-sources:network_policy:properties:endpoint:outside_endpoints", "xcsh-docs:data-sources:network_policy:properties:endpoint:prefix_list"], "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy:properties:endpoint", "parent_id": "xcsh-docs:data-sources:network_policy:reference", "path": "docs/guides/data-sources--network_policy--properties--endpoint.md", "provider_name": "network_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/properties/endpoint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint for xcsh_network_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint

Breadcrumbs:

- [xcsh_network_policy](../data-sources/network_policy.md)
- [Property reference](data-sources--network_policy--reference.md)
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

- [any](data-sources--network_policy--properties--endpoint--any.md): complete subsection reference.

- [inside_endpoints](data-sources--network_policy--properties--endpoint--inside_endpoints.md): complete subsection reference.

- [label_selector](data-sources--network_policy--properties--endpoint--label_selector.md): complete subsection reference.

- [outside_endpoints](data-sources--network_policy--properties--endpoint--outside_endpoints.md): complete subsection reference.

- [prefix_list](data-sources--network_policy--properties--endpoint--prefix_list.md): complete subsection reference.

## Next pages

- [endpoint.any](data-sources--network_policy--properties--endpoint--any.md)
- [endpoint.inside_endpoints](data-sources--network_policy--properties--endpoint--inside_endpoints.md)
- [endpoint.label_selector](data-sources--network_policy--properties--endpoint--label_selector.md)
- [endpoint.outside_endpoints](data-sources--network_policy--properties--endpoint--outside_endpoints.md)
- [endpoint.prefix_list](data-sources--network_policy--properties--endpoint--prefix_list.md)
- [Property reference](data-sources--network_policy--reference.md)
- [xcsh_network_policy](../data-sources/network_policy.md)
