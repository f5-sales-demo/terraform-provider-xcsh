---
page_title: "discovery_k8s.access_info"
subcategory: ""
description: "K8s API server access."
xcsh_docs: {"aliases": ["discovery k8s access info"], "body_bytes": 2079, "body_sha256": "sha256:46a11109c8c46f07ad6080f9db83b72cbcc451ec94222ddf6243de093caa7bb2", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:isolated", "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:reachable"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info", "parent_id": "xcsh-docs:resources:discovery:properties:discovery_k8s", "path": "documentation/resources/discovery/properties/discovery_k8s/access_info/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2112223233220021-0020313132030030-3112113333133011-3100320231332301-0311031233330113-3333031221000303-3100221211121222-2023332230113320", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:connection_info,kubeconfig_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:isolated,reachable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:isolated", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:connection_info,kubeconfig_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:isolated,reachable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:reachable", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s", "access_info"], "schema_version": 1, "sections": [{"aliases": ["discovery k8s access info connection info"], "anchor": "section", "description": "Configuration details to access discovery service REST API.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "schema-discovery_k8s--access_info--connection_info--api_server", "enforcement": "provider-schema", "group": "discovery_k8s.access_info.connection_info:RequiredObjectAttributes:api_server", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "type": "requires"}], "schema_path": ["discovery_k8s", "access_info", "connection_info"], "syntax": "block", "type": "object"}, {"aliases": ["discovery k8s access info isolated"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:isolated", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "access_info", "isolated"], "syntax": "attribute", "type": "object"}, {"aliases": ["discovery k8s access info kubeconfig url"], "anchor": "section", "description": "SecretType is used in an object to indicate a sensitive/confidential field.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info.kubeconfig_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url:blindfold_secret_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info.kubeconfig_url:ConflictingObjectAttributes:blindfold_secret_info,clear_secret_info", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url:clear_secret_info", "type": "conflicts"}], "schema_path": ["discovery_k8s", "access_info", "kubeconfig_url"], "syntax": "block", "type": "object"}, {"aliases": ["discovery k8s access info reachable"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:reachable", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "access_info", "reachable"], "syntax": "attribute", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/access_info/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "K8s API server access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["discoveryCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Additional upstream details:

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
