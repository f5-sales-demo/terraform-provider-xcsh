---
page_title: "discovery_k8s.publish_info.publish_fqdns"
subcategory: ""
description: "discovery_k8s.publish_info.publish_fqdns for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1139, "body_sha256": "sha256:5968f72b886d023967c603cbab468a7b5c9bdc2be1a41069cf5903d23ea4418b", "canonical_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "child_ids": [], "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info", "path": "docs/guides/resources--discovery--properties--discovery_k8s--publish_info--publish_fqdns.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_k8s", "publish_info", "publish_fqdns"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/publish_info/publish_fqdns/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s.publish_info.publish_fqdns for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.publish_info.publish_fqdns

Breadcrumbs:

- [xcsh_discovery](../resources/discovery.md)
- [Property reference](resources--discovery--reference.md)
- [discovery_k8s](resources--discovery--properties--discovery_k8s.md)
- [discovery_k8s.publish_info](resources--discovery--properties--discovery_k8s--publish_info.md)
- discovery_k8s.publish_info.publish_fqdns

<a id="section"></a>

Type: `["object", {}]`. Optional.

Configuration parameter for publish fqdns.

Upstream description:

This can be used for messages where no values are needed.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  }
}
```

Terraform syntax:

```terraform
publish_fqdns = {}
```

## Direct properties

This is an empty object or choice marker. It has no direct properties.

## Next pages

- [discovery_k8s.publish_info](resources--discovery--properties--discovery_k8s--publish_info.md)
- [xcsh_discovery](../resources/discovery.md)
