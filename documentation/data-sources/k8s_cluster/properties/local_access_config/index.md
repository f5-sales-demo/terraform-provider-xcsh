---
page_title: "local_access_config"
subcategory: ""
description: "Parameters required to enable local access."
xcsh_docs: {"aliases": ["local access config"], "body_bytes": 3731, "body_sha256": "sha256:1cce26157d822443bfbae3173bbc6712cbaa56ab363da29a96a59ca25cd884f7", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:local_access_config:default_port"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:local_access_config", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "documentation/data-sources/k8s_cluster/properties/local_access_config/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1232130303333302-2123302020123001-1012302120120100-2031110103100133-1333002211210233-0210013221233333-2013211033200130-2122000103121331", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_access_config"], "schema_version": 1, "sections": [{"aliases": ["local access config default port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:local_access_config:default_port", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_access_config", "default_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["local access config local domain"], "anchor": "schema-local_access_config--local_domain", "description": "Local K8s API server will be accessible at <site name>.<local domain>.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:local_access_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_access_config", "local_domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["local access config port"], "anchor": "schema-local_access_config--port", "description": "Exclusive with Use custom K8s port for API server. Available port range is less than 65000 except reserved ports.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:local_access_config", "enum_extraction_complete": false, "enum_validators": [], "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_access_config", "port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/local_access_config/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Parameters required to enable local access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_access_config

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- local_access_config

<a id="section"></a>

Type: `"single"`. Computed.

\[OneOf: local\_access\_config, no\_local\_access; Default: no\_local\_access\] Parameters required
to enable local access.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-ves-oneof-field-port_choice": "[\"default_port\",\"port\"]"
}
```

OneOf alternatives in this subsection:

- [local_access_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/local_access_config/#section)
- [no_local_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/no_local_access/#section)

Select alternatives according to the provider validators above.

## Direct properties

- [default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/local_access_config/default_port/): complete subsection reference.

<a id="schema-local_access_config--local_domain"></a>

### local_domain property

Type: `"string"`. Computed.

Local K8s API server will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

Receipt-pinned upstream constraints:

```json
{
  "maxLength": 192,
  "minLength": 1,
  "x-f5xc-constraints": {
    "category": "discovery",
    "constraintType": "string",
    "deterministic": true,
    "format": "hostname",
    "maxLength": 192,
    "metadata": {
      "confidence": 0.99,
      "source": "discovery",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minLength": 1
  },
  "x-f5xc-required-for": {
    "create": true,
    "minimum_config": true,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.message.required": "true",
    "ves.io.schema.rules.string.hostname": "true",
    "ves.io.schema.rules.string.max_len": "192",
    "ves.io.schema.rules.string.min_len": "1"
  }
}
```

<a id="schema-local_access_config--port"></a>

### port property

Type: `"number"`. Computed.

Exclusive with \[default\_port\] Use custom K8s port for API server. Available port range is less
than 65000 except reserved ports.

Receipt-pinned upstream constraints:

```json
{
  "x-f5xc-constraints": {
    "category": "networking",
    "constraintType": "number",
    "deterministic": true,
    "maximum": 65535,
    "metadata": {
      "confidence": 0.99,
      "source": "inferred",
      "validatedAt": "2026-10-06T12:36:10+00:00"
    },
    "minimum": 1,
    "multipleOf": 1
  },
  "x-f5xc-required-for": {
    "create": false,
    "minimum_config": false,
    "read": false,
    "update": false
  },
  "x-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  },
  "x-ves-validation-rules": {
    "ves.io.schema.rules.uint32.not_in_ranges": "0,6443,8005-8007,8443-8444,8505-8507,9005-9007,9090,9505-9507,9100,9115,9999,20914,23802,30805,30855,30905,30955,32222,18091-18095,65000-65334"
  }
}
```
