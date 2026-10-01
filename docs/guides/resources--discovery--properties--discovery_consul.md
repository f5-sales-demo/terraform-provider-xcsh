---
page_title: "discovery_consul"
subcategory: ""
description: "discovery_consul for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1675, "body_sha256": "sha256:f4508b01697bed4559966c397f7d3ee49e918f7efa180c353641251c836b3eea", "canonical_id": "xcsh-docs:resources:discovery:properties:discovery_consul", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:access_info", "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul", "parent_id": "xcsh-docs:resources:discovery:reference", "path": "docs/guides/resources--discovery--properties--discovery_consul.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_consul"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- discovery_consul

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

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

- [discovery_consul](resources--discovery--properties--discovery_consul.md#section)
- [discovery_k8s](resources--discovery--properties--discovery_k8s.md#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
discovery_consul {
  # Configure direct properties listed below.
}
```

## Direct properties

- [access_info](resources--discovery--properties--discovery_consul--access_info.md): complete subsection reference.

- [publish_info](resources--discovery--properties--discovery_consul--publish_info.md): complete subsection reference.

## Next pages

- [discovery_consul.access_info](resources--discovery--properties--discovery_consul--access_info.md)
- [discovery_consul.publish_info](resources--discovery--properties--discovery_consul--publish_info.md)
- [Property reference](resources--discovery--reference.md)
- [xcsh_discovery](../resources/discovery.md)
