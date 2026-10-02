---
page_title: "discovery_consul.publish_info"
subcategory: ""
description: "Consul Configuration to publish VIPs."
xcsh_docs: {"aliases": ["discovery consul publish info"], "body_bytes": 2166, "body_sha256": "sha256:f02494b06235eb1c518a70201082c1d0e3d1a6cf88c5bcf21490ed3c47781fa2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_consul:publish_info:disable_spec", "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info:publish"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_consul", "path": "documentation/resources/discovery/properties/discovery_consul/publish_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3113220101333013-1213033110231132-1230122212322132-3100032133301212-2102011203013113-2321020301320012-3012220232023130-3302132230103012", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_consul.publish_info:ConflictingObjectAttributes:disable_spec,publish", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_consul.publish_info:ConflictingObjectAttributes:disable_spec,publish", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info:publish", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "publish_info"], "schema_version": 1, "sections": [{"aliases": ["disable spec"], "anchor": "section", "description": "Enable this option", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info:disable_spec", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "publish_info", "disable_spec"], "syntax": "attribute", "type": "object"}, {"aliases": ["publish"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_consul:publish_info:publish", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_consul", "publish_info", "publish"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_consul/publish_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Consul Configuration to publish VIPs.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
