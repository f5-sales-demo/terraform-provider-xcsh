---
page_title: "discovery_k8s"
subcategory: ""
description: "discovery_k8s for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1613, "body_sha256": "sha256:8d50a29fd16397247578151569632632947e2abc76dd7bde39b77e4cb813cc1f", "canonical_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:default_all", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:namespace_mapping", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:publish_info"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "parent_id": "xcsh-docs:data-sources:discovery:reference", "path": "docs/guides/data-sources--discovery--properties--discovery_k8s.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_k8s"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# discovery_k8s

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
- [Property reference](data-sources--discovery--reference.md)
- discovery_k8s

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for discovery k8s.

Upstream description:

Discovery configuration for K8s.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-namespace_mapping_choice": "[\"default_all\",\"namespace_mapping\"]"
}
```

## Direct properties

- [access_info](data-sources--discovery--properties--discovery_k8s--access_info.md): complete subsection reference.

- [default_all](data-sources--discovery--properties--discovery_k8s--default_all.md): complete subsection reference.

- [namespace_mapping](data-sources--discovery--properties--discovery_k8s--namespace_mapping.md): complete subsection reference.

- [publish_info](data-sources--discovery--properties--discovery_k8s--publish_info.md): complete subsection reference.

## Next pages

- [discovery_k8s.access_info](data-sources--discovery--properties--discovery_k8s--access_info.md)
- [discovery_k8s.default_all](data-sources--discovery--properties--discovery_k8s--default_all.md)
- [discovery_k8s.namespace_mapping](data-sources--discovery--properties--discovery_k8s--namespace_mapping.md)
- [discovery_k8s.publish_info](data-sources--discovery--properties--discovery_k8s--publish_info.md)
- [Property reference](data-sources--discovery--reference.md)
- [xcsh_discovery](../data-sources/discovery.md)
