---
page_title: "discovery_consul"
subcategory: ""
description: "discovery_consul for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1481, "body_sha256": "sha256:76b1ce9ef526d99aec87e3564e6bd7ee8a0c668215a3a19c7c3774f1700536d5", "canonical_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info", "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul", "parent_id": "xcsh-docs:data-sources:discovery:reference", "path": "docs/guides/data-sources--discovery--properties--discovery_consul.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_consul"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# discovery_consul

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
- [Property reference](data-sources--discovery--reference.md)
- discovery_consul

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: discovery\_consul, discovery\_k8s\] Discovery configuration for Hashicorp Consul.

Upstream description:

Discovery configuration for Hashicorp Consul.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-namespace_mapping_choice": "[]"
}
```

OneOf alternatives in this subsection:

- [discovery_consul](data-sources--discovery--properties--discovery_consul.md#section)
- [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md#section)

Select alternatives according to the provider validators above.

## Direct properties

- [access_info](data-sources--discovery--properties--discovery_consul--access_info.md): complete subsection reference.

- [publish_info](data-sources--discovery--properties--discovery_consul--publish_info.md): complete subsection reference.

## Next pages

- [discovery_consul.access_info](data-sources--discovery--properties--discovery_consul--access_info.md)
- [discovery_consul.publish_info](data-sources--discovery--properties--discovery_consul--publish_info.md)
- [Property reference](data-sources--discovery--reference.md)
- [xcsh_discovery](../data-sources/discovery.md)
