---
page_title: "discovery_k8s.access_info.connection_info.tls_info.key_url"
subcategory: ""
description: "SecretType is used in an object to indicate a sensitive/confidential field."
xcsh_docs: {"aliases": ["discovery k8s access info connection info tls info key url"], "body_bytes": 1976, "body_sha256": "sha256:920aca2cd891c38ec440eaddc3259ab7a79a32340d561729f3d9816b3a8fcde6", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info:key_url:blindfold_secret_info", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info:key_url:clear_secret_info"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info:key_url", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info", "path": "documentation/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0011320033010321-0001130232332102-1301132001122023-0022211233131130-0131202000103300-1032132130202331-1002110302222223-3313020213200132", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "access_info", "connection_info", "tls_info", "key_url"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s access info connection info tls info key url blindfold secret info"], "anchor": "section", "description": "BlindfoldSecretInfoType specifies information about the Secret managed by F5XC Secret Management.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info:key_url:blindfold_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "access_info", "connection_info", "tls_info", "key_url", "blindfold_secret_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery k8s access info connection info tls info key url clear secret info"], "anchor": "section", "description": "ClearSecretInfoType specifies information about the Secret that is not encrypted.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info:tls_info:key_url:clear_secret_info", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "access_info", "connection_info", "tls_info", "key_url", "clear_secret_info"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "SecretType is used in an object to indicate a sensitive/confidential field.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["discoveryCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.access_info.connection_info.tls_info.key_url

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/)
- [discovery_k8s.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/)
- [discovery_k8s.access_info.connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/)
- [discovery_k8s.access_info.connection_info.tls_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/)
- discovery_k8s.access_info.connection_info.tls_info.key_url

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

- [blindfold_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/blindfold_secret_info/): complete subsection reference.

- [clear_secret_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/tls_info/key_url/clear_secret_info/): complete subsection reference.
