---
page_title: "discovery_k8s.access_info"
subcategory: ""
description: "K8s API server access."
xcsh_docs: {"aliases": ["discovery k8s access info"], "body_bytes": 2655, "body_sha256": "sha256:171e597f30c5c3b46277e882a49a6bf2b2f0fcc08faa0846c246205495e22785", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:isolated", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:reachable"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info", "parent_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s", "path": "documentation/data-sources/discovery/properties/discovery_k8s/access_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0111202022021232-1330011202221333-1103203232010232-3302310031331322-2013002133003130-2322212211311302-1013323221012111-1022113122013233", "registry_path": "docs/guides/data-sources--discovery--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "access_info"], "schema_version": 1, "sections": [{"aliases": ["connection info"], "anchor": "section", "description": "Configuration details to access discovery service REST API.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:connection_info", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "access_info", "connection_info"], "syntax": "attribute", "type": "object"}, {"aliases": ["isolated"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:isolated", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "access_info", "isolated"], "syntax": "attribute", "type": "object"}, {"aliases": ["kubeconfig url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": "single", "relationships": [], "schema_path": ["discovery_k8s", "access_info", "kubeconfig_url"], "syntax": "attribute", "type": "object"}, {"aliases": ["reachable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:discovery:properties:discovery_k8s:access_info:reachable", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "access_info", "reachable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/properties/discovery_k8s/access_info/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "K8s API server access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.access_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/)
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

- [connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/): complete subsection reference.

- [isolated](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/isolated/): complete subsection reference.

- [kubeconfig_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/): complete subsection reference.

- [reachable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/reachable/): complete subsection reference.

## Next pages

- [discovery_k8s.access_info.connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/connection_info/)
- [discovery_k8s.access_info.isolated](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/isolated/)
- [discovery_k8s.access_info.kubeconfig_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/)
- [discovery_k8s.access_info.reachable](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/access_info/reachable/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/discovery_k8s/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/)
