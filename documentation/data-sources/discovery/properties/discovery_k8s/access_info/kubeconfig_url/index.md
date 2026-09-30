---
page_title: "discovery_k8s.access_info.kubeconfig_url"
subcategory: ""
description: "discovery_k8s.access_info.kubeconfig_url for xcsh_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 2127, "body_sha256": "sha256:4767a00ff9fa37a8489b3b2fa27c00785a6e77c698c7aa1f9820122b8d01ff2a", "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:kubeconfig_url:blindfold_secret_info", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:kubeconfig_url:clear_secret_info"], "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info", "path": "documentation/data-sources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/index.md", "provider_name": "discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "properties", "schema_path": ["discovery_k8s", "access_info", "kubeconfig_url"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "discovery_k8s.access_info.kubeconfig_url for xcsh_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# discovery_k8s.access_info.kubeconfig_url

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/)
- [discovery_k8s.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/)
- discovery_k8s.access_info.kubeconfig_url

<a id="section"></a>

Type: `"single"`. Computed.

SecretType is used in an object to indicate a sensitive/confidential field.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-secret_info_oneof": "[\"blindfold_secret_info\",\"clear_secret_info\"]"
}
```

## Direct properties

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/clear_secret_info/): complete subsection reference.

## Next pages

- [discovery_k8s.access_info.kubeconfig_url.blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/blindfold_secret_info/)
- [discovery_k8s.access_info.kubeconfig_url.clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/clear_secret_info/)
- [discovery_k8s.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
