---
page_title: "discovery_k8s.publish_info"
subcategory: ""
description: "discovery_k8s.publish_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 3170, "body_sha256": "sha256:a6cbb23fb9a4a76a4644422c8389b11487c2907c837bee04be78762499c06bba", "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:disable_spec", "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns"], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s", "path": "documentation/resources/discovery/properties/discovery_k8s/publish_info/index.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "properties", "schema_path": ["discovery_k8s", "publish_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/publish_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s.publish_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# discovery_k8s.publish_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- discovery_k8s.publish_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for publish info.

Upstream description:

K8s Configuration to publish VIPs.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("disable_spec",
    "dns_delegation"),
  validators.ConflictingObjectAttributes("disable_spec",
    "publish"),
  validators.ConflictingObjectAttributes("disable_spec",
    "publish_fqdns"),
  validators.ConflictingObjectAttributes("dns_delegation",
    "publish"),
  validators.ConflictingObjectAttributes("dns_delegation",
    "publish_fqdns"),
  validators.ConflictingObjectAttributes("publish",
    "publish_fqdns")}
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
  "x-ves-oneof-field-publish_choice": "[\"disable\",\"dns_delegation\",\"publish\",\"publish_fqdns\"]"
}
```

Terraform syntax:

```terraform
publish_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/disable_spec/): complete subsection reference.

- [dns_delegation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/): complete subsection reference.

- [publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish/): complete subsection reference.

- [publish_fqdns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish_fqdns/): complete subsection reference.

## Next pages

- [discovery_k8s.publish_info.disable_spec](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/disable_spec/)
- [discovery_k8s.publish_info.dns_delegation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/dns_delegation/)
- [discovery_k8s.publish_info.publish](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish/)
- [discovery_k8s.publish_info.publish_fqdns](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/publish_fqdns/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
