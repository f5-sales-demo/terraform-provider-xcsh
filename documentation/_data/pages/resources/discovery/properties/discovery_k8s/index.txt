---
page_title: "discovery_k8s"
subcategory: ""
description: "Discovery configuration for K8s."
xcsh_docs: {"aliases": ["discovery k8s"], "body_bytes": 2584, "body_sha256": "sha256:fad8238a17aa4ec0cfd1307cec1005df5f9b8d32d0f7b5be8c988deee6389a9b", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:discovery:properties:discovery_k8s:access_info", "xcsh-docs:resources:discovery:properties:discovery_k8s:default_all", "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping", "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:discovery:properties:discovery_k8s", "parent_id": "xcsh-docs:resources:discovery:reference", "path": "documentation/resources/discovery/properties/discovery_k8s/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3022033122122312-0120120003232212-3132230030223031-0130311200323223-2303301232130211-1221100220233100-0330012221020323-1001233220233000", "registry_path": "docs/guides/resources--discovery--reference--group-001.md", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s:ConflictingObjectAttributes:default_all,namespace_mapping", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:default_all", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s:ConflictingObjectAttributes:default_all,namespace_mapping", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping", "type": "conflicts"}], "retrieval_version": 1, "role": "properties", "schema_path": ["discovery_k8s"], "schema_version": 1, "sections": [{"aliases": ["access info"], "anchor": "section", "description": "K8s API server access.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:connection_info,kubeconfig_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:connection_info", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:isolated,reachable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:isolated", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:connection_info,kubeconfig_url", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:kubeconfig_url", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.access_info:ConflictingObjectAttributes:isolated,reachable", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:access_info:reachable", "type": "conflicts"}], "schema_path": ["discovery_k8s", "access_info"], "syntax": "block", "type": "object"}, {"aliases": ["default all"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:default_all", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["discovery_k8s", "default_all"], "syntax": "attribute", "type": "object"}, {"aliases": ["namespace mapping"], "anchor": "section", "description": "Select the mapping between K8s namespaces from which services will be discovered and App Namespace to which the discovered services will be shared.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.namespace_mapping:RequiredObjectAttributes:items", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:namespace_mapping:items", "type": "requires"}], "schema_path": ["discovery_k8s", "namespace_mapping"], "syntax": "block", "type": "object"}, {"aliases": ["publish info"], "anchor": "section", "description": "K8s Configuration to publish VIPs.", "document_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info", "flags": [], "max_items": null, "min_items": null, "nesting": "single", "relationships": [{"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,dns_delegation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,publish", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:disable_spec", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,dns_delegation", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:dns_delegation,publish", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:dns_delegation,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:dns_delegation", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,publish", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:dns_delegation,publish", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:publish,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:disable_spec,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:dns_delegation,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "discovery_k8s.publish_info:ConflictingObjectAttributes:publish,publish_fqdns", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:discovery:properties:discovery_k8s:publish_info:publish_fqdns", "type": "conflicts"}], "schema_path": ["discovery_k8s", "publish_info"], "syntax": "block", "type": "object"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/discovery/properties/discovery_k8s/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Discovery configuration for K8s.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["discoveryCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# discovery_k8s

Breadcrumbs:

- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- discovery_k8s

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

Configuration parameter for discovery k8s.

Upstream description:

Discovery configuration for K8s.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.ConflictingObjectAttributes("default_all",
    "namespace_mapping")}
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
  "x-ves-oneof-field-namespace_mapping_choice": "[\"default_all\",\"namespace_mapping\"]"
}
```

Terraform syntax:

```terraform
discovery_k8s {
  # Configure direct properties listed below.
}
```

## Direct properties

- [access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/): complete subsection reference.

- [default_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/default_all/): complete subsection reference.

- [namespace_mapping](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/namespace_mapping/): complete subsection reference.

- [publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/): complete subsection reference.

## Next pages

- [discovery_k8s.access_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/access_info/)
- [discovery_k8s.default_all](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/default_all/)
- [discovery_k8s.namespace_mapping](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/namespace_mapping/)
- [discovery_k8s.publish_info](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/discovery_k8s/publish_info/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/properties/)
- [xcsh_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/discovery/)
