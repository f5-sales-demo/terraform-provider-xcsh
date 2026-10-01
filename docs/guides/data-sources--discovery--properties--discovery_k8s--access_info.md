---
page_title: "discovery_k8s.access_info"
subcategory: ""
description: "discovery_k8s.access_info for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2006, "body_sha256": "sha256:2fb4190d245a1a2ab9597188896c69b85a86ccb44fcf78515f44d2187f4f3035", "canonical_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:isolated", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:reachable"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "path": "docs/guides/data-sources--discovery--properties--discovery_k8s--access_info.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "properties", "schema_path": ["discovery_k8s", "access_info"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/access_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s.access_info for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.access_info

Breadcrumbs:

- [xcsh_discovery](../data-sources/discovery.md)
- [Property reference](data-sources--discovery--reference.md)
- [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md)
- discovery_k8s.access_info

<a id="section"></a>

Type: `"single"`. Computed.

Configuration parameter for access info.

Upstream description:

K8s API server access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-config_type": "[\"connection_info\",\"kubeconfig_url\"]",
  "x-ves-oneof-field-k8s_pod_network_choice": "[\"isolated\",\"reachable\"]"
}
```

## Direct properties

- [connection_info](data-sources--discovery--properties--discovery_k8s--access_info--connection_info.md): complete subsection reference.

- [isolated](data-sources--discovery--properties--discovery_k8s--access_info--isolated.md): complete subsection reference.

- [kubeconfig_url](data-sources--discovery--properties--discovery_k8s--access_info--kubeconfig_url.md): complete subsection reference.

- [reachable](data-sources--discovery--properties--discovery_k8s--access_info--reachable.md): complete subsection reference.

## Next pages

- [discovery_k8s.access_info.connection_info](data-sources--discovery--properties--discovery_k8s--access_info--connection_info.md)
- [discovery_k8s.access_info.isolated](data-sources--discovery--properties--discovery_k8s--access_info--isolated.md)
- [discovery_k8s.access_info.kubeconfig_url](data-sources--discovery--properties--discovery_k8s--access_info--kubeconfig_url.md)
- [discovery_k8s.access_info.reachable](data-sources--discovery--properties--discovery_k8s--access_info--reachable.md)
- [discovery_k8s](data-sources--discovery--properties--discovery_k8s.md)
- [xcsh_discovery](../data-sources/discovery.md)
