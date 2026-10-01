---
page_title: "discovery_k8s.publish_info"
subcategory: ""
description: "discovery_k8s.publish_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1980, "body_sha256": "sha256:6e7660abd756106cb56ce22f26ad980ff1c204550fc596638e53d7865a224c32", "canonical_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:disable_spec", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info:publish_fqdns"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "path": "docs/guides/data-sources--discovery--properties--discovery_k8s--publish_info.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_k8s", "publish_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/publish_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s.publish_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.publish_info

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
- [Property reference](data-sources--discovery--reference.md)
- [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md)
- discovery_k8s.publish_info

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for publish info.

Upstream description:

K8s Configuration to publish VIPs.

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

## Direct properties

- [disable_spec](data-sources--discovery--properties--discovery_k8s--publish_info--disable_spec.md): complete subsection reference.

- [dns_delegation](data-sources--discovery--properties--discovery_k8s--publish_info--dns_delegation.md): complete subsection reference.

- [publish](data-sources--discovery--properties--discovery_k8s--publish_info--publish.md): complete subsection reference.

- [publish_fqdns](data-sources--discovery--properties--discovery_k8s--publish_info--publish_fqdns.md): complete subsection reference.

## Next pages

- [discovery_k8s.publish_info.disable_spec](data-sources--discovery--properties--discovery_k8s--publish_info--disable_spec.md)
- [discovery_k8s.publish_info.dns_delegation](data-sources--discovery--properties--discovery_k8s--publish_info--dns_delegation.md)
- [discovery_k8s.publish_info.publish](data-sources--discovery--properties--discovery_k8s--publish_info--publish.md)
- [discovery_k8s.publish_info.publish_fqdns](data-sources--discovery--properties--discovery_k8s--publish_info--publish_fqdns.md)
- [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md)
- [xcsh_discovery](../data-sources/discovery.md)
