---
page_title: "discovery_consul.access_info"
subcategory: ""
description: "Hashicorp Consul API server information."
xcsh_docs: {"aliases": ["discovery consul access info"], "body_bytes": 1829, "body_sha256": "sha256:826c81136e95f23fee1bbab4f952655b64738d75896115e2f041a1777572e54a", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:connection_info", "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul", "path": "documentation/data-sources/discovery/properties/discovery_consul/access_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3123311310310103-2223023120311311-2332022202220111-3233313230103113-2123330001002133-1333033002132332-1211332221002031-3301200311111201", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_consul", "access_info"], "schema_version": 1, "sections": [{"aliases": ["connection info"], "anchor": "section", "description": "Configuration details to access discovery service REST API.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:connection_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_consul", "access_info", "connection_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["authentication", "credential setup", "credentials", "http basic auth info"], "anchor": "section", "description": "Authentication parameters to access Hashicorp Consul.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_consul:access_info:http_basic_auth_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_consul", "access_info", "http_basic_auth_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_consul/access_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Hashicorp Consul API server information.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_consul.access_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/)
- discovery_consul.access_info

<a id="section"></a>

Type: `"single"`. Computed.

Hashicorp Consul API server information.

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

## Direct properties

- [connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/connection_info/): complete subsection reference.

- [http_basic_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/): complete subsection reference.

## Next pages

- [discovery_consul.access_info.connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/connection_info/)
- [discovery_consul.access_info.http_basic_auth_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/access_info/http_basic_auth_info/)
- [discovery_consul](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_consul/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
