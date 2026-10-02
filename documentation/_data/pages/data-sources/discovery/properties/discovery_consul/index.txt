---
page_title: "discovery_consul"
subcategory: ""
description: "Discovery configuration for Hashicorp Consul."
xcsh_docs: {"aliases": ["discovery consul"], "body_bytes": 2090, "body_sha256": "sha256:b4b9af66a313d05f7f2da4384b1d263953d7b47182aa3324f78cfde05b9ca365", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info", "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul", "parent_id": "xcsh-docs:data-sources:discovery:reference", "path": "documentation/data-sources/discovery/properties/discovery_consul/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1003210333302012-3322023032310130-1131033123212201-2323023110233111-1012002213022222-2003130303011310-0223012102032111-0001301010022110", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul"], "schema_version": 1, "sections": [{"aliases": ["access info"], "anchor": "section", "description": "Hashicorp Consul API server information.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_consul", "access_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["publish info"], "anchor": "section", "description": "Consul Configuration to publish VIPs.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:publish_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_consul", "publish_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Discovery configuration for Hashicorp Consul.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- discovery_consul

<a id="section"></a>

Type: `"single"`. Computed.

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

- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/#section)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/): complete subsection reference.

- [publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/publish_info/): complete subsection reference.

## Next pages

- [discovery_consul.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/)
- [discovery_consul.publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/publish_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
