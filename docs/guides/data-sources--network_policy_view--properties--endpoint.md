---
page_title: "endpoint"
subcategory: ""
description: "endpoint for xcsh_network_policy_view."
xcsh_docs: {"aliases": [], "body_bytes": 1979, "body_sha256": "sha256:8a50393ddae1872c460ddea57468cbfbe85d6a521d95dc05fbac9be2e4be396a", "canonical_id": "xcsh-docs:data-sources:network_policy_view:properties:endpoint", "child_ids": ["xcsh-docs:data-sources:network_policy_view:properties:endpoint:any", "xcsh-docs:data-sources:network_policy_view:properties:endpoint:inside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:endpoint:label_selector", "xcsh-docs:data-sources:network_policy_view:properties:endpoint:outside_endpoints", "xcsh-docs:data-sources:network_policy_view:properties:endpoint:prefix_list"], "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_view:properties:endpoint", "parent_id": "xcsh-docs:data-sources:network_policy_view:reference", "path": "docs/guides/data-sources--network_policy_view--properties--endpoint.md", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["endpoint"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/properties/endpoint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "endpoint for xcsh_network_policy_view.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# endpoint

Breadcrumbs:

- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
- [Property reference](data-sources--network_policy_view--reference.md)
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

- [any](data-sources--network_policy_view--properties--endpoint--any.md): complete subsection reference.

- [inside_endpoints](data-sources--network_policy_view--properties--endpoint--inside_endpoints.md): complete subsection reference.

- [label_selector](data-sources--network_policy_view--properties--endpoint--label_selector.md): complete subsection reference.

- [outside_endpoints](data-sources--network_policy_view--properties--endpoint--outside_endpoints.md): complete subsection reference.

- [prefix_list](data-sources--network_policy_view--properties--endpoint--prefix_list.md): complete subsection reference.

## Next pages

- [endpoint.any](data-sources--network_policy_view--properties--endpoint--any.md)
- [endpoint.inside_endpoints](data-sources--network_policy_view--properties--endpoint--inside_endpoints.md)
- [endpoint.label_selector](data-sources--network_policy_view--properties--endpoint--label_selector.md)
- [endpoint.outside_endpoints](data-sources--network_policy_view--properties--endpoint--outside_endpoints.md)
- [endpoint.prefix_list](data-sources--network_policy_view--properties--endpoint--prefix_list.md)
- [Property reference](data-sources--network_policy_view--reference.md)
- [xcsh_network_policy_view](../data-sources/network_policy_view.md)
