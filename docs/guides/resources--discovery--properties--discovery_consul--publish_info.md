---
page_title: "discovery_consul.publish_info"
subcategory: ""
description: "discovery_consul.publish_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1713, "body_sha256": "sha256:5f0c587a648cfad48d1ddf4f3eac77c691aacf87a7cb56b4d4915f6385cdb4d8", "canonical_id": "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:publish_info:disable_spec", "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info:publish"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul", "path": "docs/guides/resources--discovery--properties--discovery_consul--publish_info.md", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_consul", "publish_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/publish_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.publish_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.publish_info

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- [discovery_consul](resources--discovery--properties--discovery_consul.md)
- discovery_consul.publish_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for publish info.

Upstream description:

Consul Configuration to publish VIPs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "publish")}
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
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"publish\"]"
}
```

Terraform syntax:

```terraform
publish_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](resources--discovery--properties--discovery_consul--publish_info--disable_spec.md): complete subsection reference.

- [publish](resources--discovery--properties--discovery_consul--publish_info--publish.md): complete subsection reference.

## Next pages

- [discovery_consul.publish_info.disable_spec](resources--discovery--properties--discovery_consul--publish_info--disable_spec.md)
- [discovery_consul.publish_info.publish](resources--discovery--properties--discovery_consul--publish_info--publish.md)
- [discovery_consul](resources--discovery--properties--discovery_consul.md)
- [xcsh_discovery](../resources/discovery.md)
