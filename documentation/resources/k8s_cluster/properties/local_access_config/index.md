---
page_title: "local_access_config"
subcategory: ""
description: "Parameters required to enable local access."
xcsh_docs: {"aliases": ["local access config"], "body_bytes": 4959, "body_sha256": "sha256:63b9d2a6fc59329a82fa94e3cd704feda888a0d4a2833e3c84b0978915e5b75f", "capabilities": ["container"], "category": "container", "child_ids": ["xcsh-docs:resources:k8s_cluster:properties:local_access_config:default_port"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:k8s_cluster:collection", "completeness": "complete", "id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config", "parent_id": "xcsh-docs:resources:k8s_cluster:reference", "path": "documentation/resources/k8s_cluster/properties/local_access_config/index.md", "product": "distributed-cloud", "provider_name": "k8s_cluster", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "resources", "registry_anchor": "canonical-2330113202310232-1313112231102321-0100021023331021-3030000310021310-0120111130300032-1301121303202001-0330002111131003-3333210313223011", "registry_path": "docs/guides/resources--k8s_cluster--reference--group-001.md", "relationships": [{"anchor": "schema-local_access_config--port", "enforcement": "provider-schema", "group": "local_access_config:ConflictingObjectAttributes:default_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config", "type": "conflicts"}, {"anchor": "section", "enforcement": "provider-schema", "group": "local_access_config:ConflictingObjectAttributes:default_port,port", "source": "ast-validator:ConflictingObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config:default_port", "type": "conflicts"}, {"anchor": "schema-local_access_config--local_domain", "enforcement": "provider-schema", "group": "local_access_config:RequiredObjectAttributes:local_domain", "source": "ast-validator:RequiredObjectAttributes", "target_id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config", "type": "requires"}], "retrieval_version": 1, "role": "properties", "schema_path": ["local_access_config"], "schema_version": 1, "sections": [{"aliases": ["default port"], "anchor": "section", "description": "This can be used for messages where no values are needed.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config:default_port", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_access_config", "default_port"], "syntax": "attribute", "type": "object"}, {"aliases": ["local domain"], "anchor": "schema-local_access_config--local_domain", "description": "Local K8s API server will be accessible at <site name>.<local domain>.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_access_config", "local_domain"], "syntax": "attribute", "type": "string"}, {"aliases": ["port"], "anchor": "schema-local_access_config--port", "description": "Exclusive with Use custom K8s port for API server. Available port range is less than 65000 except reserved ports.", "document_id": "xcsh-docs:resources:k8s_cluster:properties:local_access_config", "flags": ["optional"], "max_items": null, "min_items": null, "nesting": null, "relationships": [], "schema_path": ["local_access_config", "port"], "syntax": "attribute", "type": "number"}], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/k8s_cluster/properties/local_access_config/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Parameters required to enable local access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["k8s_clusterCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# local_access_config

Breadcrumbs:

- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- local_access_config

<a id="section"></a>

Type: `"object"`. single nested block, Optional.

\[OneOf: local\_access\_config, no\_local\_access; Default: no\_local\_access\] Parameters required
to enable local access.

Upstream description:

Parameters required to enable local access.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Object{validators.RequiredObjectAttributes("local_domain"),
  validators.ConflictingObjectAttributes("default_port",
    "port")}
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
  "x-ves-oneof-field-port_choice": "[\"default_port\",\"port\"]"
}
```

OneOf alternatives in this subsection:

- [local_access_config](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/local_access_config/#section)
- [no_local_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/no_local_access/#section)

Select alternatives according to the provider validators above.

Terraform syntax:

```terraform
local_access_config {
  # Configure direct properties listed below.
}
```

## Direct properties

- [default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/local_access_config/default_port/): complete subsection reference.

<a id="schema-local_access_config--local_domain"></a>

### local_domain property

Type: `"string"`. Optional.

Local K8s API server will be accessible at &lt;site name&gt;.&lt;local domain&gt;.

Provider validators and defaults (from schema source):

```go
Validators: []validator.String{
  stringvalidator.LengthBetween(1, 192),
}
```

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

Type: `"number"`. Optional.

Exclusive with \[default\_port\] Use custom K8s port for API server. Available port range is less
than 65000 except reserved ports.

Upstream description:

Exclusive with \[default\_port\] Use custom K8s port for API server. Available port range is less
than 65000 except reserved ports.

Provider validators and defaults (from schema source):

```go
Validators: []validator.Int64{
  int64validator.Between(1, 65535),
}
```

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

- [local_access_config.default_port](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/local_access_config/default_port/)
- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/properties/)
- [xcsh_k8s_cluster](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/k8s_cluster/)
