---
page_title: "discovery_consul.publish_info"
subcategory: ""
description: "discovery_consul.publish_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2067, "body_sha256": "sha256:b6a8c719bf2fd7bf8604573a49dadb93d0e8b8478bc41c9eb4860177a03da18c", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:publish_info:disable_spec", "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info:publish"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul", "path": "documentation/resources/discovery/properties/discovery_consul/publish_info/index.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["discovery_consul", "publish_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/publish_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_consul.publish_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# discovery_consul.publish_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/)
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

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/publish_info/disable_spec/): complete subsection reference.

- [publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/publish_info/publish/): complete subsection reference.

## Next pages

- [discovery_consul.publish_info.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/publish_info/disable_spec/)
- [discovery_consul.publish_info.publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/publish_info/publish/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_consul/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
