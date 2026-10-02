---
page_title: "discovery_k8s.access_info"
subcategory: ""
description: "K8s API server access."
xcsh_docs: {"aliases": ["discovery k8s access info"], "body_bytes": 2986, "body_sha256": "sha256:7d1ac0e2725e40e47dfa6ec66efb065b4c0987d27a4558aa329e15e8e6ce9d0e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:isolated", "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:reachable"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s", "path": "documentation/resources/discovery/properties/discovery_k8s/access_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:connection_info,kubeconfig_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:isolated,reachable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:isolated", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:connection_info,kubeconfig_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:isolated,reachable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:reachable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "access_info"], "schema_version": 1, "sections": [{"aliases": ["connection info"], "anchor": "section", "description": "Configuration details to access discovery service REST API.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-discovery_k8s--access_info--connection_info--api_server", "enforcement": "provider-schema", "group": "discovery_k8s.access_info.connection_info:RequiredObjectAttributes:api_server", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "type": "requires"}], "schema_path": ["discovery_k8s", "access_info", "connection_info"], "syntax": "block", "type": "object"}, {"aliases": ["isolated"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:isolated", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "access_info", "isolated"], "syntax": "attribute", "type": "object"}, {"aliases": ["kubeconfig url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info.kubeconfig_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info.kubeconfig_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url:clear_secret_info", "type": "conflicts"}], "schema_path": ["discovery_k8s", "access_info", "kubeconfig_url"], "syntax": "block", "type": "object"}, {"aliases": ["reachable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:reachable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "access_info", "reachable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/access_info/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "K8s API server access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s.access_info

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- discovery_k8s.access_info

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for access info.

Upstream description:

K8s API server access.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("connection_info",
    "kubeconfig_url"),
  validators.ConflictingObjectAttributes("isolated",
    "reachable")}
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
  "x-ves-oneof-field-config_type": "[\"connection_info\",\"kubeconfig_url\"]",
  "x-ves-oneof-field-k8s_pod_network_choice": "[\"isolated\",\"reachable\"]"
}
```

Terraform syntax:

```terraform
access_info {
  # Configure direct properties listed below.
}
```

## Direct properties

- [connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/): complete subsection reference.

- [isolated](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/isolated/): complete subsection reference.

- [kubeconfig_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/): complete subsection reference.

- [reachable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/reachable/): complete subsection reference.

## Next pages

- [discovery_k8s.access_info.connection_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/connection_info/)
- [discovery_k8s.access_info.isolated](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/isolated/)
- [discovery_k8s.access_info.kubeconfig_url](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/kubeconfig_url/)
- [discovery_k8s.access_info.reachable](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/reachable/)
- [discovery_k8s](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
