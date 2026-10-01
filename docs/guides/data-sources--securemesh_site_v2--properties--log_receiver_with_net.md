---
page_title: "log_receiver_with_net"
subcategory: ""
description: "log_receiver_with_net for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 2137, "body_sha256": "sha256:9682cb6aaddc92e758c994c485ae59bb7c0a9ff8fda4df9ab763108fb36dbeaa", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net:log_receiver", "xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net:use_management_network", "xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net:use_slo_sli"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:properties:log_receiver_with_net", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:reference", "path": "docs/guides/data-sources--securemesh_site_v2--properties--log_receiver_with_net.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["log_receiver_with_net"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/properties/log_receiver_with_net/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "log_receiver_with_net for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# log_receiver_with_net

Breadcrumbs:

- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- log_receiver_with_net

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: log\_receiver\_with\_net, logs\_streaming\_disabled\] Select log receiver for logs
streaming with network option.

Upstream description:

Select log receiver for logs streaming with network option.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-network_choice": "[\"use_management_network\",\"use_slo_sli\"]"
}
```

OneOf alternatives in this subsection:

- [log_receiver_with_net](data-sources--securemesh_site_v2--properties--log_receiver_with_net.md#section)
- [logs_streaming_disabled](data-sources--securemesh_site_v2--properties--logs_streaming_disabled.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [log_receiver](data-sources--securemesh_site_v2--properties--log_receiver_with_net--log_receiver.md): complete subsection reference.

- [use_management_network](data-sources--securemesh_site_v2--properties--log_receiver_with_net--use_management_network.md): complete subsection reference.

- [use_slo_sli](data-sources--securemesh_site_v2--properties--log_receiver_with_net--use_slo_sli.md): complete subsection reference.

## Next pages

- [log_receiver_with_net.log_receiver](data-sources--securemesh_site_v2--properties--log_receiver_with_net--log_receiver.md)
- [log_receiver_with_net.use_management_network](data-sources--securemesh_site_v2--properties--log_receiver_with_net--use_management_network.md)
- [log_receiver_with_net.use_slo_sli](data-sources--securemesh_site_v2--properties--log_receiver_with_net--use_slo_sli.md)
- [Property reference](data-sources--securemesh_site_v2--reference.md)
- [xcsh_securemesh_site_v2](../data-sources/securemesh_site_v2.md)
