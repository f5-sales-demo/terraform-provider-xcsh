---
page_title: "local_access_config"
subcategory: ""
description: "Parameters required to enable local access."
xcsh_docs: {"aliases": ["local access config"], "body_bytes": 4356, "body_sha256": "sha256:36c5c0bc7b4af1a5c555c65db186bc4f597213cc5e3b6cd86f0d5f0a20b8dafd", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:data-sources:k8s_cluster:properties:local_access_config:default_port"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:k8s_cluster:properties:local_access_config", "parent_id": "xcsh-docs:data-sources:k8s_cluster:reference", "path": "documentation/data-sources/k8s_cluster/properties/local_access_config/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1232130303333302-2123302020123001-1012302120120100-2031110103100133-1333002211210233-0210013221233333-2013211033200130-2122000103121331", "registry_path": "docs/guides/data-sources--k8s_cluster--reference--group-001.md", "relationships": [], "retrieval_version": 1, "role": "properties", "schema_path": ["local_access_config"], "schema_version": 1, "sections": [{"aliases": ["default port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:local_access_config:default_port", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_access_config", "default_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["local domain"], "anchor": "schema-local_access_config--local_domain", "description": "Local K8s API server will be accessible at <site name>.<local domain>.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:local_access_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_access_config", "local_domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "schema-local_access_config--port", "description": "Exclusive with Use custom K8s port for API server. Available port range is less than 65000 except reserved ports.", "document_id": "xcsh-docs:data-sources:k8s_cluster:properties:local_access_config", "flags": ["computed"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_access_config", "port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/k8s_cluster/properties/local_access_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters required to enable local access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

Upstream description:

Parameters required to enable local access.

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

Upstream description:

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
      "validatedAt": "2026-09-29T03:20:54+00:00"
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

## Next pages

- [local_access_config.default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/local_access_config/default_port/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/properties/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/k8s_cluster/)
